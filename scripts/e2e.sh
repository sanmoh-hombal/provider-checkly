#!/usr/bin/env bash
# scripts/e2e.sh — E2E smoke test against a kind cluster + Checkly sandbox.
# Creates a kind cluster, installs CRDs, runs the provider out-of-cluster,
# applies every namespaced example, waits for Ready, then tears down.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "${ROOT}"

# ── Configuration ────────────────────────────────────────────────────────────
CLUSTER_NAME="${CLUSTER_NAME:-provider-checkly-e2e}"
TIMEOUT="${E2E_TIMEOUT:-5m}"
PROVIDER_LOG="${ROOT}/_output/provider-e2e.log"
PROVIDER_PID=""
PASS=0
FAIL=0
SKIP=0
declare -a APPLIED_FILES=()

# ── Source credentials from .env ─────────────────────────────────────────────
if [[ -f "${ROOT}/.env" ]]; then
  # shellcheck disable=SC1091
  source "${ROOT}/.env"
fi

: "${CHECKLY_API_KEY_SANDBOX:?CHECKLY_API_KEY_SANDBOX must be set}"
: "${CHECKLY_ACCOUNT_ID_SANDBOX:?CHECKLY_ACCOUNT_ID_SANDBOX must be set}"

# ── Ensure terraform 1.5.7 is on PATH ───────────────────────────────────────
export PATH="${HOME}/.local/bin:${ROOT}/_output/bin/tools:${PATH}"
TERRAFORM_VERSION="${TERRAFORM_VERSION:-1.5.7}"
if ! command -v terraform &>/dev/null; then
  echo "==> Installing Terraform ${TERRAFORM_VERSION}"
  TERRAFORM_DIR="${ROOT}/_output/bin/tools"
  mkdir -p "${TERRAFORM_DIR}"
  ARCH="$(uname -m)"; [[ "${ARCH}" == "x86_64" ]] && ARCH="amd64"; [[ "${ARCH}" == "aarch64" || "${ARCH}" == "arm64" ]] && ARCH="arm64"
  OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
  curl -fsSL "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${OS}_${ARCH}.zip" \
    -o /tmp/terraform.zip
  unzip -o /tmp/terraform.zip -d "${TERRAFORM_DIR}"
  rm -f /tmp/terraform.zip
  export PATH="${TERRAFORM_DIR}:${PATH}"
fi
echo "    terraform: $(terraform version | head -1)"

# ── Cleanup on exit ──────────────────────────────────────────────────────────
cleanup() {
  local exit_code=$?
  echo ""
  echo "==> Cleaning up..."

  # Delete managed resources so Terraform destroys remote objects
  echo "    Deleting managed resources..."
  for f in "${APPLIED_FILES[@]+"${APPLIED_FILES[@]}"}"; do
    kubectl delete --ignore-not-found=true --wait=false -f "$f" 2>/dev/null || true
  done

  if [[ ${#APPLIED_FILES[@]} -gt 0 ]]; then
    echo "    Waiting 20s for provider to process deletions..."
    sleep 20
  fi

  if [[ -n "${PROVIDER_PID}" ]]; then
    echo "    Stopping provider (PID ${PROVIDER_PID})..."
    kill "${PROVIDER_PID}" 2>/dev/null || true
    wait "${PROVIDER_PID}" 2>/dev/null || true
  fi

  echo "    Deleting kind cluster..."
  kind delete cluster --name "${CLUSTER_NAME}" 2>/dev/null || true

  echo ""
  echo "══════════════════════════════════════════════════════════════"
  echo "  E2E RESULTS: ${PASS} passed, ${FAIL} failed, ${SKIP} skipped"
  echo "══════════════════════════════════════════════════════════════"
  if [[ -f "${PROVIDER_LOG}" ]]; then
    echo "  Provider log: ${PROVIDER_LOG}"
  fi
  exit ${exit_code}
}
trap cleanup EXIT

# ── Helper: apply a resource, return 0 on success ───────────────────────────
apply_resource() {
  local f="$1"
  echo -n "    apply ${f} ... "
  if kubectl apply -f "${f}" 2>&1; then
    APPLIED_FILES+=("${f}")
    return 0
  else
    echo "    ✗ APPLY FAILED"
    FAIL=$((FAIL + 1))
    return 1
  fi
}

# ── Helper: wait for a single resource to become Ready ───────────────────────
wait_ready_file() {
  local f="$1"
  local name
  name=$(grep -m1 '^  name:' "${f}" | awk '{print $2}')
  echo -n "    wait  ${name} (${f##*/}) ... "
  if kubectl wait -f "${f}" --for=condition=Ready --timeout="${TIMEOUT}" 2>&1; then
    PASS=$((PASS + 1))
    return 0
  else
    echo "    ✗ TIMEOUT / NOT READY"
    echo "    ── Status conditions ──"
    kubectl get -f "${f}" -o jsonpath='{range .status.conditions[*]}    {.type}: {.status} ({.reason}) {.message}{"\n"}{end}' 2>/dev/null || true
    FAIL=$((FAIL + 1))
    return 1
  fi
}

# ══════════════════════════════════════════════════════════════════════════════
# 1. Kind cluster
# ══════════════════════════════════════════════════════════════════════════════
echo "==> Creating kind cluster: ${CLUSTER_NAME}"
kind delete cluster --name "${CLUSTER_NAME}" 2>/dev/null || true
kind create cluster --name "${CLUSTER_NAME}" --wait 120s

# ══════════════════════════════════════════════════════════════════════════════
# 2. Install CRDs
# ══════════════════════════════════════════════════════════════════════════════
echo "==> Installing Crossplane CRDs from package/crds/"
kubectl apply -f package/crds/ --server-side=true 2>&1 | wc -l | xargs -I{} echo "    {} CRDs applied"

# ══════════════════════════════════════════════════════════════════════════════
# 3. Build the provider
# ══════════════════════════════════════════════════════════════════════════════
echo "==> Building provider binary"
GOOS="$(go env GOOS)"
GOARCH="$(go env GOARCH)"
PROVIDER_BIN="${ROOT}/_output/bin/${GOOS}_${GOARCH}/provider"
mkdir -p "$(dirname "${PROVIDER_BIN}")"
CGO_ENABLED=0 go build -trimpath -o "${PROVIDER_BIN}" \
  -ldflags '-s -w' \
  "${ROOT}/cmd/provider" 2>&1 | tail -3

if [[ ! -x "${PROVIDER_BIN}" ]]; then
  echo "ERROR: provider binary not found at ${PROVIDER_BIN}"
  exit 1
fi
echo "    Binary: ${PROVIDER_BIN}"

# ══════════════════════════════════════════════════════════════════════════════
# 4. Credentials
# ══════════════════════════════════════════════════════════════════════════════
echo "==> Setting up credentials"
kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -

CREDS_JSON=$(cat <<CEOF
{
  "api_key": "${CHECKLY_API_KEY_SANDBOX}",
  "account_id": "${CHECKLY_ACCOUNT_ID_SANDBOX}",
  "api_url": "https://api.checklyhq.com"
}
CEOF
)

for ns in crossplane-system default; do
  kubectl -n "${ns}" create secret generic checkly-creds \
    --from-literal=credentials="${CREDS_JSON}" \
    --dry-run=client -o yaml | kubectl apply -f -
done

# Cluster-scoped ProviderConfig
kubectl apply -f - <<'EOF'
apiVersion: checkly.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: checkly-creds
      namespace: crossplane-system
      key: credentials
EOF

# Namespaced ProviderConfig (in default ns)
kubectl apply -f - <<'EOF'
apiVersion: checkly.m.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: checkly-creds
      namespace: default
      key: credentials
EOF

# Prerequisite K8s secrets referenced by examples
echo "==> Creating prerequisite K8s secrets for examples"
kubectl -n default create secret generic checkgroup-env \
  --from-literal=base-url=https://api.example.com \
  --dry-run=client -o yaml | kubectl apply -f -
kubectl -n default create secret generic checkgroupv2-env \
  --from-literal=base-url=https://api.example.com \
  --dry-run=client -o yaml | kubectl apply -f -
kubectl -n default create secret generic envvar-api-base-url \
  --from-literal=value=https://api.example.com \
  --dry-run=client -o yaml | kubectl apply -f -

openssl req -x509 -newkey rsa:2048 -keyout /tmp/e2e-tls.key -out /tmp/e2e-tls.crt \
  -days 1 -nodes -subj '/CN=api.example.com' 2>/dev/null
kubectl -n default create secret generic tls-cert \
  --from-file=tls.crt=/tmp/e2e-tls.crt \
  --from-file=tls.key=/tmp/e2e-tls.key \
  --dry-run=client -o yaml | kubectl apply -f -
rm -f /tmp/e2e-tls.key /tmp/e2e-tls.crt

# ══════════════════════════════════════════════════════════════════════════════
# 5. Start the provider out-of-cluster
# ══════════════════════════════════════════════════════════════════════════════
echo "==> Setting up Terraform provider mirror"
export TERRAFORM_VERSION="${TERRAFORM_VERSION:-1.5.7}"
export TERRAFORM_PROVIDER_SOURCE="${TERRAFORM_PROVIDER_SOURCE:-checkly/checkly}"
export TERRAFORM_PROVIDER_VERSION="${TERRAFORM_PROVIDER_VERSION:-1.22.0}"
TERRAFORM_NATIVE_PROVIDER_BINARY="terraform-provider-checkly_v${TERRAFORM_PROVIDER_VERSION}"

MIRROR_DIR="${ROOT}/_output/terraform-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${GOOS}_${GOARCH}"
mkdir -p "${MIRROR_DIR}"

NATIVE_SRC="${ROOT}/.work/terraform/.terraform/providers/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${GOOS}_${GOARCH}/${TERRAFORM_NATIVE_PROVIDER_BINARY}"
if [[ -f "${NATIVE_SRC}" ]]; then
  cp -f "${NATIVE_SRC}" "${MIRROR_DIR}/"
else
  echo "    Native provider not cached locally, downloading from GitHub releases..."
  DOWNLOAD_URL="https://github.com/checkly/terraform-provider-checkly/releases/download/v${TERRAFORM_PROVIDER_VERSION}/terraform-provider-checkly_${TERRAFORM_PROVIDER_VERSION}_${GOOS}_${GOARCH}.zip"
  curl -fsSL "${DOWNLOAD_URL}" -o /tmp/tf-provider.zip
  unzip -o /tmp/tf-provider.zip -d "${MIRROR_DIR}"
  rm -f /tmp/tf-provider.zip
  chmod +x "${MIRROR_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}"
fi

export TF_CLI_CONFIG_FILE="${ROOT}/_output/terraform-mirror/.terraformrc"
cat > "${TF_CLI_CONFIG_FILE}" <<TFRC
provider_installation {
  filesystem_mirror {
    path    = "${ROOT}/_output/terraform-mirror"
    include = ["${TERRAFORM_PROVIDER_SOURCE}"]
  }
  direct {
    exclude = ["${TERRAFORM_PROVIDER_SOURCE}"]
  }
}
TFRC

export TERRAFORM_NATIVE_PROVIDER_PATH="${MIRROR_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}"
export TF_FORK=0

echo "==> Starting provider out-of-cluster (log: ${PROVIDER_LOG})"
mkdir -p "$(dirname "${PROVIDER_LOG}")"
"${PROVIDER_BIN}" --debug > "${PROVIDER_LOG}" 2>&1 &
PROVIDER_PID=$!
echo "    Provider PID: ${PROVIDER_PID}"
echo "    Waiting 10s for controllers to start..."
sleep 10

if ! kill -0 "${PROVIDER_PID}" 2>/dev/null; then
  echo "ERROR: provider process died. Last 20 lines:"
  tail -20 "${PROVIDER_LOG}"
  exit 1
fi

# ══════════════════════════════════════════════════════════════════════════════
# 6. Apply examples in phases
# ══════════════════════════════════════════════════════════════════════════════

# Phase 1: Independent resources
PHASE1=(
  examples/namespaced/alerts/alertchannel.yaml
  examples/namespaced/alerts/maintenancewindow.yaml
  examples/namespaced/checks/check.yaml
  examples/namespaced/checks/checkgroup.yaml
  examples/namespaced/checks/checkgroupv2.yaml
  examples/namespaced/checks/dnsmonitor.yaml
  examples/namespaced/checks/heartbeat.yaml
  examples/namespaced/checks/heartbeatmonitor.yaml
  examples/namespaced/checks/icmpmonitor.yaml
  examples/namespaced/checks/tcpcheck.yaml
  examples/namespaced/checks/tcpmonitor.yaml
  examples/namespaced/checks/urlmonitor.yaml
  examples/namespaced/infra/envvar.yaml
  examples/namespaced/infra/privatelocation.yaml
  examples/namespaced/infra/snippet.yaml
  examples/namespaced/statuspage/dashboard.yaml
  examples/namespaced/statuspage/statuspage.yaml
)

# Phase 2: Resources that depend on Phase 1
PHASE2=(
)

# Skipped: need real tar.gz, real IDs, or sandbox plan limitations
SKIPPED=(
  examples/namespaced/checks/playwrightcodebundle.yaml   # needs real prebuilt tar.gz archive
  examples/namespaced/checks/playwrightchecksuite.yaml   # depends on PlaywrightCodeBundle
  examples/namespaced/triggers/triggercheck.yaml          # needs real checkId from API
  examples/namespaced/triggers/triggergroup.yaml           # needs real groupId from API
  examples/namespaced/infra/clientcertificate.yaml         # 403: sandbox plan limitation
  examples/namespaced/statuspage/service.yaml              # TF schema mismatch: status_page_id removed
)

echo ""
echo "==> Phase 1: Applying ${#PHASE1[@]} independent resources"
for f in "${PHASE1[@]}"; do
  apply_resource "${f}" || true
done

echo ""
echo "==> Phase 1: Waiting for Ready (timeout=${TIMEOUT})"
for f in "${PHASE1[@]}"; do
  wait_ready_file "${f}" || true
done

if [[ ${#PHASE2[@]} -gt 0 ]]; then
echo ""
echo "==> Phase 2: Applying ${#PHASE2[@]} dependent resources"
for f in "${PHASE2[@]}"; do
  apply_resource "${f}" || true
done

echo ""
echo "==> Phase 2: Waiting for Ready (timeout=${TIMEOUT})"
for f in "${PHASE2[@]}"; do
  wait_ready_file "${f}" || true
done
fi

echo ""
echo "==> Skipped ${#SKIPPED[@]} resources (missing prerequisites):"
for f in "${SKIPPED[@]}"; do
  echo "    SKIP: ${f}"
  SKIP=$((SKIP + 1))
done

echo ""
echo "==> Final resource state:"
kubectl get managed -A --no-headers 2>/dev/null || echo "    (no managed resources found via 'kubectl get managed')"

# Exit non-zero if any resource failed
if [[ ${FAIL} -gt 0 ]]; then
  exit 1
fi
