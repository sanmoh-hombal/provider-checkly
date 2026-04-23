# Snippet

Reusable code snippet for setup/teardown scripts.

## Details

| | |
|---|---|
| **API Group** | `infra.checkly.crossplane.io/v1alpha1` |
| **Kind** | `Snippet` |
| **Terraform Resource** | [`checkly_snippet`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/snippet) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | `string` | (String) The name of the snippet
The name of the snippet |
| `script` | `string` | (String) Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.... |

## Example

```yaml
apiVersion: infra.checkly.m.crossplane.io/v1alpha1
kind: Snippet
metadata:
  annotations:
    meta.upbound.io/example-id: infra/v1alpha1/snippet
  labels:
    testing.upbound.io/example-name: example_1
  name: example-1
  namespace: upbound-system
spec:
  forProvider:
    name: Example 1
    script: console.log('test');
```

## Notes

- Both **namespace-scoped** (`infra.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`infra.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/snippet) for detailed field descriptions.
