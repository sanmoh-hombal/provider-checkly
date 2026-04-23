# Check

API or browser check that monitors an endpoint.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `Check` |
| **Terraform Resource** | [`checkly_check`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/check) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines if the check is running or not. Possible values true, and false.
Determines if the check is runn... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the check. (see below for nested schema)
Determines t... |
| `degradedResponseTime` | `number` | (Number) The response time in milliseconds starting from which a check should be considered degraded. Possible values... |
| `description` | `string` | (String) A description of the check.
A description of the check. |
| `doubleCheck` | `boolean` | (Boolean, Deprecated) Setting this to true will trigger a retry when a check fails from the failing region and anothe... |
| `environmentVariable` | `array<object>` | (Block List) Insert environment variables into the runtime environment. Only relevant for browser checks. Use global ... |
| `environmentVariables` | `object` | (Map of String, Deprecated) Key/value pairs of environment variables to insert into the runtime environment.
Key/valu... |
| `frequency` | `number` | use frequency_offset to define the actual frequency), 1 (1 minute), 2 (2 minutes), 5 (5 minutes), 10 (10 minutes), 15... |
| `frequencyOffset` | `number` | use frequency to define the actual frequency), 10 (10 seconds), 20 (20 seconds) and 30 (30 seconds).
Only relevant wh... |
| `groupId` | `number` | (Number) The id of the check group this check is part of.
The id of the check group this check is part of. |
| `groupIdRef` | `object` | Reference to a CheckGroup in checks to populate groupId. |
| `groupIdSelector` | `object` | Selector for a CheckGroup in checks to populate groupId. |
| `groupOrder` | `number` | (Number) The position of this check in a check group. It determines in what order checks are run when a group is trig... |
| `localSetupScript` | `string` | (String) A valid piece of Node.js code to run in the setup phase.
A valid piece of Node.js code to run in the setup p... |
| `localTeardownScript` | `string` | (String) A valid piece of Node.js code to run in the teardown phase.
A valid piece of Node.js code to run in the tear... |
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
| `request` | `array<object>` | (Block Set, Max: 1) An API check might have one request config. (see below for nested schema)
An API check might have... |
| `retryStrategy` | `array<object>` | (Block List, Max: 1) A strategy for retrying failed check/monitor runs. (see below for nested schema)
A strategy for ... |
| `runParallel` | `boolean` | robin.
Determines if the check should run in all selected locations in parallel or round-robin. |
| `runtimeId` | `string` | (String) The id of the runtime to use for this check.
The id of the runtime to use for this check. |
| `script` | `string` | (String) A valid piece of Node.js JavaScript code describing a browser interaction with the Puppeteer/Playwright fram... |
| `setupSnippetId` | `number` | (Number) An ID reference to a snippet to use in the setup phase of an API check.
An ID reference to a snippet to use ... |
| `shouldFail` | `boolean` | (Boolean) Allows to invert the behaviour of when a check is considered to fail. Allows for validating error status li... |
| `sslCheck` | `boolean` | (Boolean, Deprecated) Determines if the SSL certificate should be validated for expiry.
Determines if the SSL certifi... |
| `sslCheckDomain` | `string` | (String) A valid fully qualified domain name (FQDN) to check its SSL certificate.
A valid fully qualified domain name... |
| `tags` | `array<string>` | (Set of String) A list of tags for organizing and filtering checks.
A list of tags for organizing and filtering checks. |
| `teardownSnippetId` | `number` | (Number) An ID reference to a snippet to use in the teardown phase of an API check.
An ID reference to a snippet to u... |
| `triggerIncident` | `array<object>` | (Block Set, Max: 1) Create and resolve an incident based on the alert configuration. Useful for status page automatio... |
| `type` | `string` | (String) The type of the check. Possible values are API, BROWSER, and MULTI_STEP.
The type of the check. Possible val... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this check.
When... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: Check
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/check
  labels:
    testing.upbound.io/example-name: example_check
  name: example-check
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    frequency: 1
    locations:
    - us-west-1
    name: Example check
    request:
    - assertion:
      - comparison: EQUALS
        source: STATUS_CODE
        target: "200"
      followRedirects: true
      skipSsl: false
      url: https://api.example.com/
    shouldFail: false
    type: API
    useGlobalAlertSettings: true

```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/check) for detailed field descriptions.
