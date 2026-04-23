# StatusPageService

Service entry linked to a status page.

## Details

| | |
|---|---|
| **API Group** | `statuspage.checkly.crossplane.io/v1alpha1` |
| **Kind** | `StatusPageService` |
| **Terraform Resource** | [`checkly_status_page_service`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/status_page_service) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | `string` | (String) The name of the service.
The name of the service. |
| `statusPageId` | `string` | (String) The ID of this resource. |
| `statusPageIdRef` | `object` | Reference to a StatusPage in statuspage to populate statusPageId. |
| `statusPageIdSelector` | `object` | Selector for a StatusPage in statuspage to populate statusPageId. |

## Example

```yaml
apiVersion: statuspage.checkly.m.crossplane.io/v1alpha1
kind: StatusPageService
metadata:
  annotations:
    meta.upbound.io/example-id: statuspage/v1alpha1/statuspageservice
  labels:
    testing.upbound.io/example-name: backend
  name: backend
  namespace: upbound-system
spec:
  forProvider:
    name: Backend
```

## Notes

- Both **namespace-scoped** (`statuspage.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`statuspage.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/status_page_service) for detailed field descriptions.
