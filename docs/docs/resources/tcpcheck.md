# TCPCheck

TCP connectivity check against a host and port.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `TCPCheck` |
| **Terraform Resource** | [`checkly_tcp_check`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/tcp_check) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines if the check is running or not. Possible values true, and false.
Determines if the check is runn... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the monitor. (see below for nested schema)
Determines... |
| `degradedResponseTime` | `number` | (Number) The response time in milliseconds starting from which a check should be considered degraded. Possible values... |
| `description` | `string` | (String) A description of the monitor.
A description of the monitor. |
| `frequency` | `number` | use frequency_offset to define the actual frequency), 1 (1 minute), 2 (2 minutes), 5 (5 minutes), 10 (10 minutes), 15... |
| `frequencyOffset` | `number` | use frequency to define the actual frequency), 10 (10 seconds), 20 (20 seconds) and 30 (30 seconds).
When `frequency`... |
| `groupId` | `number` | (Number) The id of the check group this check is part of.
The id of the check group this check is part of. |
| `groupIdRef` | `object` | Reference to a CheckGroup in checks to populate groupId. |
| `groupIdSelector` | `object` | Selector for a CheckGroup in checks to populate groupId. |
| `groupOrder` | `number` | (Number) The position of this check in a check group. It determines in what order checks are run when a group is trig... |
| `locations` | `array<string>` | east-1"])
An array of one or more data center locations where to run the this check. (Default ["us-east-1"]) |
| `maxResponseTime` | `number` | (Number) The response time in milliseconds starting from which a check should be considered failing. Possible values ... |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when a check fails/degrades/recovers.
Determines if any no... |
| `name` | `string` | (String) The name of the check.
The name of the check. |
| `privateLocationRefs` | `array<object>` | References to PrivateLocation in infra to populate privateLocations. |
| `privateLocations` | `array<string>` | (Set of String) An array of one or more private locations slugs.
An array of one or more private locations slugs. |
| `privateLocationsSelector` | `object` | Selector for a list of PrivateLocation in infra to populate privateLocations. |
| `request` | `array<object>` | (Block Set, Min: 1, Max: 1) The parameters for the TCP connection. (see below for nested schema)
The parameters for t... |
| `retryStrategy` | `array<object>` | (Block List, Max: 1) A strategy for retrying failed check/monitor runs. (see below for nested schema)
A strategy for ... |
| `runParallel` | `boolean` | robin.
Determines if the check should run in all selected locations in parallel or round-robin. |
| `runtimeId` | `string` | (String) The ID of the runtime to use for this check.
The ID of the runtime to use for this check. |
| `shouldFail` | `boolean` | (Boolean) Allows to invert the behaviour of when a check is considered to fail.
Allows to invert the behaviour of whe... |
| `tags` | `array<string>` | (Set of String) A list of tags for organizing and filtering checks.
A list of tags for organizing and filtering checks. |
| `triggerIncident` | `array<object>` | (Block Set, Max: 1) Create and resolve an incident based on the alert configuration. Useful for status page automatio... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this check.
When... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: TCPCheck
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/tcpcheck
  labels:
    testing.upbound.io/example-name: example-tcp-check
  name: example-tcp-check
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    frequency: 1
    locations:
    - us-west-1
    name: Example TCP check
    request:
    - hostname: api.checklyhq.com
      port: 80
    shouldFail: false
    useGlobalAlertSettings: true
```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/tcp_check) for detailed field descriptions.
