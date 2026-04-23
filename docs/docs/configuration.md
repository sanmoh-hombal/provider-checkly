# Configuration

## ProviderConfig

Every managed resource references a `ProviderConfig` that supplies Checkly API credentials. The default config is named `default`.

```yaml
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

### Credentials format

The referenced Secret must contain a JSON object with the following fields:

| Field | Required | Description |
|-------|----------|-------------|
| `api_key` | Yes | Your Checkly API key |
| `account_id` | Yes | Your Checkly account ID |
| `api_url` | No | API endpoint (defaults to `https://api.checklyhq.com`) |

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: checkly-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "cu_xxx",
      "account_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
      "api_url": "https://api.checklyhq.com"
    }
```

## Cluster vs Namespace scope

The provider generates both cluster-scoped and namespace-scoped variants of every resource:

| Scope | API group pattern | Use case |
|-------|-------------------|----------|
| Namespace-scoped | `<group>.checkly.crossplane.io/v1alpha1` | Multi-tenant isolation per namespace (recommended) |
| Cluster-scoped | `<group>.checkly.m.crossplane.io/v1alpha1` | Single-tenant / backward compatibility |

All examples default to namespace-scoped resources.

## Multiple ProviderConfigs

You can create multiple `ProviderConfig` resources to manage checks across different Checkly accounts:

```yaml
apiVersion: checkly.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: staging
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: checkly-creds-staging
      key: credentials
```

Then reference it from a managed resource:

```yaml
spec:
  providerConfigRef:
    name: staging
```

## Terraform provider fallback

For Checkly resources not yet wrapped by this provider, use [`provider-terraform`](https://github.com/upbound/provider-terraform) as a fallback. This lets you write raw HCL blocks as Crossplane resources until native support is added.
