# Quickstart

Get from zero to a managed Checkly API check in four steps.

## Prerequisites

- A running Kubernetes cluster (or use [kind](https://kind.sigs.k8s.io/) to create one)
- [Crossplane](https://docs.crossplane.io/latest/software/install/) **v2.x** installed
- A [Checkly](https://checklyhq.com) account with an API key and account ID

## 1. Create a cluster

Skip this if you already have a Kubernetes cluster with Crossplane installed.

```bash
kind create cluster --name crossplane-demo
```

## 2. Install Crossplane

```bash
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane crossplane-stable/crossplane \
  --namespace crossplane-system \
  --create-namespace
```

## 3. Install provider-checkly

```bash
crossplane xpkg install provider xpkg.crossplane.io/crossplane-contrib/provider-checkly:v0.1.0
```

Wait for the provider to become healthy:

```bash
kubectl get providers -w
```

## 4. Configure credentials and create a Check

Create the Checkly credentials secret:

```yaml
# secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: checkly-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<your-checkly-api-key>",
      "account_id": "<your-checkly-account-id>",
      "api_url": "https://api.checklyhq.com"
    }
```

Apply the `ProviderConfig`:

```yaml
# providerconfig.yaml
apiVersion: checkly.crossplane.io/v1beta1
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
```

Create your first API check:

```yaml
# check.yaml
apiVersion: checks.checkly.crossplane.io/v1alpha1
kind: Check
metadata:
  name: homepage-api
  namespace: default
spec:
  forProvider:
    name: homepage-api
    activated: true
    frequency: 5
    type: API
    locations:
      - eu-west-1
      - us-east-1
    tags:
      - managed-by-crossplane
    request:
      - method: GET
        url: https://example.com
        assertion:
          - source: STATUS_CODE
            property: ""
            comparison: EQUALS
            target: "200"
  providerConfigRef:
    name: default
```

Apply everything and verify:

```bash
kubectl apply -f secret.yaml -f providerconfig.yaml -f check.yaml

# Watch the check become ready
kubectl get checks -w
```

## Next steps

- Browse all [Resources](resources/index.md) to see what you can manage.
- See [Configuration](configuration.md) for credential options and provider tuning.
- Check out [example scenarios](https://github.com/crossplane-contrib/provider-checkly/tree/main/examples/scenarios) for real-world patterns.
