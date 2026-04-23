# Heartbeat

Heartbeat (cron/ping) check that expects periodic pings.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `Heartbeat` |
| **Terraform Resource** | [`checkly_heartbeat`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/heartbeat) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines if the check is running or not. Possible values true, and false.
Determines if the check is runn... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the monitor. (see below for nested schema)
Determines... |
| `description` | `string` | (String) A description of the monitor.
A description of the monitor. |
| `heartbeat` | `array<object>` | (Block Set, Min: 1, Max: 1) (see below for nested schema) |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when a check fails/degrades/recovers.
Determines if any no... |
| `name` | `string` | (String) The name of the check.
The name of the check. |
| `tags` | `array<string>` | (Set of String) A list of tags for organizing and filtering checks.
A list of tags for organizing and filtering checks. |
| `triggerIncident` | `array<object>` | (Block Set, Max: 1) Create and resolve an incident based on the alert configuration. Useful for status page automatio... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this check.
When... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: Heartbeat
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/heartbeat
  labels:
    testing.upbound.io/example-name: example-heartbeat
  name: example-heartbeat
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    heartbeat:
    - grace: 1
      graceUnit: days
      period: 7
      periodUnit: days
    name: Example heartbeat
    useGlobalAlertSettings: true
```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/heartbeat) for detailed field descriptions.
