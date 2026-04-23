# CheckGroup

Logical grouping of checks with shared settings.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `CheckGroup` |
| **Terraform Resource** | [`checkly_check_group`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/check_group) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines if the checks in the group are running or not.
Determines if the checks in the group are running... |
| `alertChannelSubscription` | `array<object>` | (Block Set) An array of channel IDs and whether they're activated or not. If you don't set at least one alert channel... |
| `alertSettings` | `array<object>` | (Block List, Max: 1) Determines the alert escalation policy for the check. (see below for nested schema)
Determines t... |
| `apiCheckDefaults` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `concurrency` | `number` | (Number) Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.... |
| `doubleCheck` | `boolean` | (Boolean, Deprecated) Setting this to true will trigger a retry when a check fails from the failing region and anothe... |
| `environmentVariable` | `array<object>` | (Block List) Insert environment variables into the runtime environment. Only relevant for browser checks. Use global ... |
| `environmentVariables` | `object` | (Map of String, Deprecated) Key/value pairs of environment variables to insert into the runtime environment.
Key/valu... |
| `localSetupScript` | `string` | (String) A valid piece of Node.js code to run in the setup phase of an API check in this group.
A valid piece of Node... |
| `localTeardownScript` | `string` | (String) A valid piece of Node.js code to run in the teardown phase of an API check in this group.
A valid piece of N... |
| `locations` | `array<string>` | (Set of String) An array of one or more data center locations where to run the checks.
An array of one or more data c... |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when a check in this group fails and/or recovers.
Determin... |
| `name` | `string` | (String) The name of the check group.
The name of the check group. |
| `privateLocationRefs` | `array<object>` | References to PrivateLocation in infra to populate privateLocations. |
| `privateLocations` | `array<string>` | (Set of String) An array of one or more private locations slugs.
An array of one or more private locations slugs. |
| `privateLocationsSelector` | `object` | Selector for a list of PrivateLocation in infra to populate privateLocations. |
| `retryStrategy` | `array<object>` | (Block List, Max: 1) A strategy for retrying failed check/monitor runs. (see below for nested schema)
A strategy for ... |
| `runParallel` | `boolean` | robin.
Determines if the checks in the group should run in all selected locations in parallel or round-robin. |
| `runtimeId` | `string` | (String) The id of the runtime to use for this group.
The id of the runtime to use for this group. |
| `setupSnippetId` | `number` | (Number) An ID reference to a snippet to use in the setup phase of an API check.
An ID reference to a snippet to use ... |
| `tags` | `array<string>` | (Set of String) Tags for organizing and filtering checks.
Tags for organizing and filtering checks. |
| `teardownSnippetId` | `number` | (Number) An ID reference to a snippet to use in the teardown phase of an API check.
An ID reference to a snippet to u... |
| `useGlobalAlertSettings` | `boolean` | (Boolean) When true, the account level alert settings will be used, not the alert setting defined on this check group... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: CheckGroup
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/checkgroup
  labels:
    testing.upbound.io/example-name: test_group1
  name: test-group1
  namespace: upbound-system
spec:
  forProvider:
    activated: true
    alertSettings:
    - escalationType: RUN_BASED
      reminders:
      - amount: 2
        interval: 5
      runBasedEscalation:
      - failedRunThreshold: 1
    apiCheckDefaults:
    - assertion:
      - comparison: EQUALS
        property: ""
        source: STATUS_CODE
        target: "200"
      - comparison: CONTAINS
        property: ""
        source: TEXT_BODY
        target: welcome
      basicAuth:
      - password: pass
        username: user
      headers:
        X-Test: foo
      queryParameters:
        query: foo
      url: http://example.com/
    concurrency: 3
    environmentVariable:
    - key: TEST_ENV_VAR
      locked: false
      valueSecretRef:
        key: example-key
        name: example-secret
    - key: ADDITIONAL_ENV_VAR
      locked: true
      valueSecretRef:
        key: example-key
        name: example-secret
    localSetupScript: setup-test
    localTeardownScript: teardown-test
    locations:
    - eu-west-1
    muted: false
    name: My test group 1
    tags:
    - auto
    useGlobalAlertSettings: false

```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/check_group) for detailed field descriptions.
