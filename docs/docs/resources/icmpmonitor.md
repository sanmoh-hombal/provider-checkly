# ICMPMonitor

ICMP (ping) monitor for host reachability.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `ICMPMonitor` |
| **Terraform Resource** | [`checkly_icmp_monitor`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/icmp_monitor) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines whether the monitor will run periodically or not after being deployed.
Determines whether the mo... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the monitor. (see below for nested schema)
Determines... |
| `degradedPacketLossThreshold` | `number` | (Number) The packet loss percentage where the monitor should be considered degraded. Possible values are between 0 an... |
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
| `locations` | `array<string>` | (Set of String) An array of one or more data center locations where to run this monitor.
An array of one or more data... |
| `maxPacketLossThreshold` | `number` | (Number) The packet loss percentage where the monitor should be considered failing. Possible values are between 0 and... |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when the monitor fails and/or recovers. (Default false).
D... |
| `name` | `string` | (String) The name of the monitor.
The name of the monitor. |
| `request` | `array<object>` | (Block List, Min: 1, Max: 1) The parameters of the ICMP request. (see below for nested schema)
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
kind: ICMPMonitor
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/icmpmonitor
  labels:
    testing.upbound.io/example-name: example-icmp-monitor
  name: example-icmp-monitor
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    frequency: 10
    locations:
    - eu-west-1
    name: Example ICMP monitor
    request:
    - assertion:
      - comparison: LESS_THAN
        property: avg
        source: LATENCY
        target: "200"
      hostname: example.com
      ipFamily: IPv4
      pingCount: 10
    useGlobalAlertSettings: true
```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/icmp_monitor) for detailed field descriptions.
