# AlertChannel

Notification channel (email, Slack, SMS, webhook, PagerDuty, etc.).

## Details

| | |
|---|---|
| **API Group** | `alerts.checkly.crossplane.io/v1alpha1` |
| **Kind** | `AlertChannel` |
| **Terraform Resource** | [`checkly_alert_channel`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/alert_channel) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `call` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `email` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `opsgenie` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `pagerduty` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `sendDegraded` | `boolean` | (Boolean) (Default false)
(Default `false`) |
| `sendFailure` | `boolean` | (Boolean) (Default true)
(Default `true`) |
| `sendRecovery` | `boolean` | (Boolean) (Default true)
(Default `true`) |
| `slack` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `sms` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `sslExpiry` | `boolean` | (Boolean) (Default false)
(Default `false`) |
| `sslExpiryThreshold` | `number` | (Number) Value must be between 1 and 30 (Default 30)
Value must be between 1 and 30 (Default `30`) |
| `webhook` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |

## Example

```yaml
apiVersion: alerts.checkly.m.crossplane.io/v1alpha1
kind: AlertChannel
metadata:
  annotations:
    meta.upbound.io/example-id: alerts/v1alpha1/alertchannel
  labels:
    testing.upbound.io/example-name: email_ac
  name: email-ac
  namespace: upbound-system
spec:
  forProvider:
    email:
    - address: john@example.com
    sendDegraded: true
    sendFailure: false
    sendRecovery: true
    sslExpiry: true
    sslExpiryThreshold: 22

```

## Notes

- Both **namespace-scoped** (`alerts.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`alerts.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/alert_channel) for detailed field descriptions.
