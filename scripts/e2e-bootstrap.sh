#!/usr/bin/env bash
set -euo pipefail

CLUSTER_NAME="${CLUSTER_NAME:-provider-checkly-e2e}"

echo "==> Creating kind cluster: ${CLUSTER_NAME}"
kind create cluster --name "${CLUSTER_NAME}" --wait 60s

echo "==> Installing Crossplane CRDs"
kubectl apply -f package/crds/

echo "==> Creating crossplane-system namespace"
kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -

echo "==> Creating ProviderConfig secret (placeholder)"
kubectl apply -f - <<'EOF'
apiVersion: v1
kind: Secret
metadata:
  name: checkly-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "cu_test_placeholder",
      "account_id": "test-account",
      "api_url": "https://api.checklyhq.com"
    }
EOF

echo "==> Applying ProviderConfig"
kubectl apply -f - <<'EOF'
apiVersion: template.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: checkly-creds
      key: credentials
EOF

echo "==> E2E bootstrap complete"
