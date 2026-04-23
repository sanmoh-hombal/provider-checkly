# MaintenanceWindow

Scheduled maintenance window that suppresses alerts.

## Details

| | |
|---|---|
| **API Group** | `alerts.checkly.crossplane.io/v1alpha1` |
| **Kind** | `MaintenanceWindow` |
| **Terraform Resource** | [`checkly_maintenance_windows`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/maintenance_windows) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `endsAt` | `string` | (String) The end date of the maintenance window.
The end date of the maintenance window. |
| `name` | `string` | (String) The maintenance window name.
The maintenance window name. |
| `repeatEndsAt` | `string` | (String) The date on which the maintenance window should stop repeating.
The date on which the maintenance window sho... |
| `repeatInterval` | `number` | (Number) The repeat interval of the maintenance window from the first occurrence.
The repeat interval of the maintena... |
| `repeatUnit` | `string` | (String) The repeat cadence for the maintenance window. Possible values DAY, WEEK and MONTH.
The repeat cadence for t... |
| `startsAt` | `string` | (String) The start date of the maintenance window.
The start date of the maintenance window. |
| `tags` | `array<string>` | (Set of String) The names of the checks and groups maintenance window should apply to.
The names of the checks and gr... |

## Example

```yaml
apiVersion: alerts.checkly.m.crossplane.io/v1alpha1
kind: MaintenanceWindow
metadata:
  annotations:
    meta.upbound.io/example-id: alerts/v1alpha1/maintenancewindow
  labels:
    testing.upbound.io/example-name: maintenance-1
  name: maintenance-1
  namespace: upbound-system
spec:
  forProvider:
    endsAt: "2014-08-25T00:00:00.000Z"
    name: Maintenance Windows
    repeatEndsAt: "2014-08-24T00:00:00.000Z"
    repeatInterval: 1
    repeatUnit: MONTH
    startsAt: "2014-08-24T00:00:00.000Z"
    tags:
    - production
```

## Notes

- Both **namespace-scoped** (`alerts.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`alerts.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/maintenance_windows) for detailed field descriptions.
