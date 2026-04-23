# PlaywrightCheckSuite

Playwright test suite that runs browser-based checks.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `PlaywrightCheckSuite` |
| **Terraform Resource** | [`checkly_playwright_check_suite`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/playwright_check_suite) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines whether the check will run periodically or not after being deployed.
Determines whether the chec... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the check. (see below for nested schema)
Determines t... |
| `bundle` | `array<object>` | (Block List, Min: 1, Max: 1) Attaches a code bundle to the check. (see below for nested schema)
Attaches a code bundl... |
| `description` | `string` | (String) A description of the check.
A description of the check. |
| `environmentVariable` | `array<object>` | (Block List) Insert environment variables into the execution environment. (see below for nested schema)
Insert enviro... |
| `frequency` | `number` | (Number) Controls how often the check should run. Defined in minutes. The allowed values are 1 (1 minute), 2 (2 minut... |
| `groupId` | `number` | (Number) The ID of the check group that this check is part of.
The ID of the check group that this check is part of. |
| `groupOrder` | `number` | (Number) The position of the check in the check group. It determines in what order checks and monitors are run when a... |
| `locations` | `array<string>` | (Set of String) An array of one or more data center locations where to run the this check.
An array of one or more da... |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when the check fails and/or recovers. (Default false).
Det... |
| `name` | `string` | (String) The name of the check.
The name of the check. |
| `privateLocationRefs` | `array<object>` | References to PrivateLocation in infra to populate privateLocations. |
| `privateLocations` | `array<string>` | (Set of String) An array of one or more private locations slugs.
An array of one or more private locations slugs. |
| `privateLocationsSelector` | `object` | Selector for a list of PrivateLocation in infra to populate privateLocations. |
| `runParallel` | `boolean` | robin. (Default false).
Determines whether the check should run on all selected locations in parallel or round-robin.... |
| `runtime` | `array<object>` | (Block List, Max: 1) Configure the runtime environment of the Playwright check. (see below for nested schema)
Configu... |
| `tags` | `array<string>` | (Set of String) A list of tags for organizing and filtering checks and monitors.
A list of tags for organizing and fi... |
| `triggerIncident` | `array<object>` | (Block Set, Max: 1) Create and resolve an incident based on the alert configuration. Useful for status page automatio... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this check. (Def... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: PlaywrightCheckSuite
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/playwrightchecksuite
  labels:
    testing.upbound.io/example-name: example-playwright-check
  name: example-playwright-check
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    bundle:
    - idSelector:
        matchLabels:
          testing.upbound.io/example-name: playwright-bundle
      metadata: ${checkly_playwright_code_bundle.playwright-bundle.metadata}
    frequency: 2
    locations:
    - eu-west-1
    name: Example Playwright check
    useGlobalAlertSettings: true

```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/playwright_check_suite) for detailed field descriptions.
