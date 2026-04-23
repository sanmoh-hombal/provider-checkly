# StatusPage

Public status page with service health indicators.

## Details

| | |
|---|---|
| **API Group** | `statuspage.checkly.crossplane.io/v1alpha1` |
| **Kind** | `StatusPage` |
| **Terraform Resource** | [`checkly_status_page`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/status_page) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `card` | `array<object>` | (Block List, Min: 1) A list of cards to include on the status page. (see below for nested schema)
A list of cards to ... |
| `customDomain` | `string` | (String) A custom user domain, e.g. "status.example.com". See the docs on updating your DNS and SSL usage.
A custom u... |
| `defaultTheme` | `string` | (String) Possible values are AUTO, DARK, and LIGHT. (Default AUTO).
Possible values are `AUTO`, `DARK`, and `LIGHT`. ... |
| `favicon` | `string` | (String) A URL to an image file to use as the favicon of the status page.
A URL to an image file to use as the favico... |
| `logo` | `string` | (String) A URL to an image file to use as the logo for the status page.
A URL to an image file to use as the logo for... |
| `name` | `string` | (String) The name of the status page.
The name of the status page. |
| `redirectTo` | `string` | (String) The URL the user should be redirected to when clicking the logo.
The URL the user should be redirected to wh... |
| `url` | `string` | (String) The URL of the status page.
The URL of the status page. |

## Example

```yaml
apiVersion: statuspage.checkly.m.crossplane.io/v1alpha1
kind: StatusPage
metadata:
  annotations:
    meta.upbound.io/example-id: statuspage/v1alpha1/statuspage
  labels:
    testing.upbound.io/example-name: example
  name: example
  namespace: upbound-system
spec:
  forProvider:
    card:
    - name: Services
      serviceAttachment:
      - serviceId: ${checkly_status_page_service.api.id}
      - serviceId: ${checkly_status_page_service.database.id}
    defaultTheme: DARK
    name: Example Application
    url: my-example-status-page

```

## Notes

- Both **namespace-scoped** (`statuspage.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`statuspage.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/status_page) for detailed field descriptions.
