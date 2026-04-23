# Conformance

provider-checkly passes the Crossplane v2.0 conformance suite with a 100% pass rate.

## v0.1.0

| Metric | Value |
|--------|-------|
| Crossplane | v2.1.3 |
| Conformance suite | `crossplane/conformance:v2.0.0-cf.1` |
| Tests executed | 56 |
| Passed | **56** |
| Failed | 0 |
| **Pass rate** | **100%** |

All 17 test classes pass with zero failures.

For the full conformance report with reproduction steps, see [`docs/conformance/v0.1.0.md`](https://github.com/sanmoh-hombal/provider-checkly/blob/main/docs/conformance/v0.1.0.md).

## How to run conformance tests

```bash
# 1. Create kind cluster + install Crossplane with Operations enabled
make controlplane.up CROSSPLANE_ARGS="--enable-operations"

# 2. Build and deploy the provider
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=https://github.com/checkly/terraform-provider-checkly/releases/download/v1.22.0
make local-deploy

# 3. Pre-load images into kind
docker save sonobuoy/sonobuoy:v0.57.4 crossplane/conformance:v2.0.0-cf.1 \
  | docker exec -i local-dev-control-plane ctr --namespace=k8s.io images import -

# 4. Run conformance
sonobuoy run --wait \
  --plugin https://raw.githubusercontent.com/crossplane/conformance/release-2.0/plugin-crossplane.yaml \
  --sonobuoy-image sonobuoy/sonobuoy:v0.57.4

# 5. Retrieve results
sonobuoy results "$(sonobuoy retrieve)" --mode detailed
```
