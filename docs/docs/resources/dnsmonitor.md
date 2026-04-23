# DNSMonitor

DNS resolution monitor for domain records.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `DNSMonitor` |
| **Terraform Resource** | [`checkly_dns_monitor`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/dns_monitor) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines whether the monitor will run periodically or not after being deployed.
Determines whether the mo... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the monitor. (see below for nested schema)
Determines... |
| `degradedResponseTime` | `number` | (Number) The response time in milliseconds where the monitor should be considered degraded. Possible values are betwe... |
| `description` | `string` | (String) A description of the monitor.
A description of the monitor. |
| `frequency` | `number` | use frequency_offset to define the actual frequency), 1 (1 minute), 2 (2 minutes), 5 (5 minutes), 10 (10 minutes), 15... |
| `frequencyOffset` | `number` | use frequency to define the actual frequency), 10 (10 seconds), 20 (20 seconds) and 30 (30 seconds).
When `frequency`... |
| `groupId` | `number` | (Number) The ID of the check group that this monitor is part of.
The ID of the check group that this monitor is part of. |
| `groupIdRef` | `object` | Reference to a CheckGroup in checks to populate groupId. |
| `groupIdSelector` | `object` | Selector for a CheckGroup in checks to populate groupId. |
| `groupOrder` | `number` | (Number) The position of the monitor in the check group. It determines in what order checks and monitors are run when... |
| `locations` | `array<string>` | (Set of String) An array of one or more data center locations where to run the this monitor.
An array of one or more ... |
| `maxResponseTime` | `number` | (Number) The response time in milliseconds where the monitor should be considered failing. Possible values are betwee... |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when the monitor fails and/or recovers. (Default false).
D... |
| `name` | `string` | (String) The name of the monitor.
The name of the monitor. |
| `request` | `array<object>` | (Block List, Min: 1, Max: 1) The parameters of the HTTP request. (see below for nested schema)
The parameters of the ... |
| `retryStrategy` | `array<object>` | (Block List, Max: 1) A strategy for retrying failed check/monitor runs. (see below for nested schema)
A strategy for ... |
| `runParallel` | `boolean` | robin. (Default false).
Determines whether the monitor should run on all selected locations in parallel or round-robi... |
| `tags` | `array<string>` | (Set of String) A list of tags for organizing and filtering checks and monitors.
A list of tags for organizing and fi... |
| `triggerIncident` | `array<object>` | (Block Set, Max: 1) Create and resolve an incident based on the alert configuration. Useful for status page automatio... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this monitor. (D... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: DNSMonitor
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/dnsmonitor
  labels:
    testing.upbound.io/example-name: example-dns-monitor
  name: example-dns-monitor
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    frequency: 2
    locations:
    - eu-west-1
    name: Example DNS monitor
    request:
    - assertion:
      - comparison: EQUALS
        source: RESPONSE_CODE
        target: NOERROR
      nameServer:
      - host: 1.1.1.1
        port: 53
      query: welcome.checklyhq.com
      recordType: A
    useGlobalAlertSettings: true
```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/dns_monitor) for detailed field descriptions.
