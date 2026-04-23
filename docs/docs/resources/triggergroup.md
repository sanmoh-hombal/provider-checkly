# TriggerGroup

Webhook trigger that runs a check group on demand.

## Details

| | |
|---|---|
| **API Group** | `triggers.checkly.crossplane.io/v1alpha1` |
| **Kind** | `TriggerGroup` |
| **Terraform Resource** | [`checkly_trigger_group`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/trigger_group) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `groupId` | `number` | (Number) The id of the group that you want to attach the trigger to.
The id of the group that you want to attach the ... |
| `groupIdRef` | `object` | Reference to a CheckGroup in checks to populate groupId. |
| `groupIdSelector` | `object` | Selector for a CheckGroup in checks to populate groupId. |
| `token` | `string` | (String) The token value created to trigger the group
The token value created to trigger the group |
| `url` | `string` | (String) The request URL to trigger the group run.
The request URL to trigger the group run. |

## Example

```yaml
apiVersion: triggers.checkly.m.crossplane.io/v1alpha1
kind: TriggerGroup
metadata:
  annotations:
    meta.upbound.io/example-id: triggers/v1alpha1/triggergroup
  labels:
    testing.upbound.io/example-name: test_trigger_group
  name: test-trigger-group
  namespace: upbound-system
spec:
  forProvider:
    groupIdSelector:
      matchLabels:
        testing.upbound.io/example-name: example
```

## Notes

- Both **namespace-scoped** (`triggers.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`triggers.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/trigger_group) for detailed field descriptions.
