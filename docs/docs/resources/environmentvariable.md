# EnvironmentVariable

Global environment variable available to all checks.

## Details

| | |
|---|---|
| **API Group** | `infra.checkly.crossplane.io/v1alpha1` |
| **Kind** | `EnvironmentVariable` |
| **Terraform Resource** | [`checkly_environment_variable`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/environment_variable) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `locked` | `boolean` | (Boolean) |
| `secret` | `boolean` | (Boolean) |
| `valueSecretRef` | `object` | (String) |

## Example

```yaml
apiVersion: infra.checkly.m.crossplane.io/v1alpha1
kind: EnvironmentVariable
metadata:
  annotations:
    meta.upbound.io/example-id: infra/v1alpha1/environmentvariable
  labels:
    testing.upbound.io/example-name: variable_1
  name: variable-1
  namespace: upbound-system
spec:
  forProvider:
    locked: true
    valueSecretRef:
      key: example-key
      name: example-secret
```

## Notes

- Both **namespace-scoped** (`infra.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`infra.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/environment_variable) for detailed field descriptions.
