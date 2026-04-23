# URLMonitor

URL availability monitor with response assertions.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `URLMonitor` |
| **Terraform Resource** | [`checkly_url_monitor`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/url_monitor) |

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
| `privateLocationRefs` | `array<object>` | References to PrivateLocation in infra to populate privateLocations. |
| `privateLocations` | `array<string>` | (Set of String) An array of one or more private locations slugs.
An array of one or more private locations slugs. |
| `privateLocationsSelector` | `object` | Selector for a list of PrivateLocation in infra to populate privateLocations. |
| `request` | `array<object>` | (Block Set, Min: 1, Max: 1) The parameters of the HTTP request. (see below for nested schema)
The parameters of the H... |
| `retryStrategy` | `array<object>` | (Block List, Max: 1) A strategy for retrying failed check/monitor runs. (see below for nested schema)
A strategy for ... |
| `runParallel` | `boolean` | robin. (Default false).
Determines whether the monitor should run on all selected locations in parallel or round-robi... |
| `shouldFail` | `boolean` | (Boolean) Allows to invert the behaviour of when the monitor is considered to fail. (Default false).
Allows to invert... |
| `tags` | `array<string>` | (Set of String) A list of tags for organizing and filtering checks and monitors.
A list of tags for organizing and fi... |
| `triggerIncident` | `array<object>` | (Block Set, Max: 1) Create and resolve an incident based on the alert configuration. Useful for status page automatio... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this monitor. (D... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: URLMonitor
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/urlmonitor
  labels:
    testing.upbound.io/example-name: example-url-monitor
  name: example-url-monitor
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    frequency: 2
    locations:
    - eu-west-1
    name: Example URL monitor
    request:
    - assertion:
      - comparison: EQUALS
        source: STATUS_CODE
        target: "200"
      url: https://welcome.checklyhq.com
    useGlobalAlertSettings: true
```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/url_monitor) for detailed field descriptions.
