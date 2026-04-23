# TriggerCheck

Webhook trigger that runs a single check on demand.

## Details

| | |
|---|---|
| **API Group** | `triggers.checkly.crossplane.io/v1alpha1` |
| **Kind** | `TriggerCheck` |
| **Terraform Resource** | [`checkly_trigger_check`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/trigger_check) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `checkId` | `string` | (String) The id of the check that you want to attach the trigger to.
The id of the check that you want to attach the ... |
| `checkIdRef` | `object` | Reference to a Check in checks to populate checkId. |
| `checkIdSelector` | `object` | Selector for a Check in checks to populate checkId. |
| `token` | `string` | (String) The token value created to trigger the check
The token value created to trigger the check |
| `url` | `string` | (String) The request URL to trigger the check run.
The request URL to trigger the check run. |

## Example

```yaml
apiVersion: triggers.checkly.m.crossplane.io/v1alpha1
kind: TriggerCheck
metadata:
  annotations:
    meta.upbound.io/example-id: triggers/v1alpha1/triggercheck
  labels:
    testing.upbound.io/example-name: test_trigger_check
  name: test-trigger-check
  namespace: upbound-system
spec:
  forProvider:
    checkIdSelector:
      matchLabels:
        testing.upbound.io/example-name: example
```

## Notes

- Both **namespace-scoped** (`triggers.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`triggers.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/trigger_check) for detailed field descriptions.
