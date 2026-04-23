# PrivateLocation

Private check runner location inside your network.

## Details

| | |
|---|---|
| **API Group** | `infra.checkly.crossplane.io/v1alpha1` |
| **Kind** | `PrivateLocation` |
| **Terraform Resource** | [`checkly_private_location`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/private_location) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `icon` | `string` | (String) Icon assigned to the private location.
Icon assigned to the private location. |
| `name` | `string` | (String) The private location name.
The private location name. |
| `slugName` | `string` | (String) Valid slug name.
Valid slug name. |

## Example

```yaml
apiVersion: infra.checkly.m.crossplane.io/v1alpha1
kind: PrivateLocation
metadata:
  annotations:
    meta.upbound.io/example-id: infra/v1alpha1/privatelocation
  labels:
    testing.upbound.io/example-name: location
  name: location
  namespace: upbound-system
spec:
  forProvider:
    name: New Private Location
    slugName: new-private-location
```

## Notes

- Both **namespace-scoped** (`infra.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`infra.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/private_location) for detailed field descriptions.
