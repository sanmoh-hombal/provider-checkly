# PlaywrightCodeBundle

Uploadable Playwright code bundle for browser checks.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `PlaywrightCodeBundle` |
| **Terraform Resource** | [`checkly_playwright_code_bundle`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/playwright_code_bundle) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `prebuiltArchive` | `array<object>` | (Block List, Min: 1, Max: 1) A prebuilt archive containing the code bundle. (see below for nested schema)
A prebuilt ... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: PlaywrightCodeBundle
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/playwrightcodebundle
  labels:
    testing.upbound.io/example-name: example-1
  name: example-1
  namespace: upbound-system
spec:
  forProvider:
    prebuiltArchive:
    - file: ${data.archive_file.playwright-bundle.output_path}
```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/playwright_code_bundle) for detailed field descriptions.
