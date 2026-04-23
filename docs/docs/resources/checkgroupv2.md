# CheckGroupV2

V2 check group with enhanced configuration options.

## Details

| | |
|---|---|
| **API Group** | `checks.checkly.crossplane.io/v1alpha1` |
| **Kind** | `CheckGroupV2` |
| **Terraform Resource** | [`checkly_check_group_v2`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/check_group_v2) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `activated` | `boolean` | (Boolean) Determines whether activated checks in the group should run or not. Deactivating the group will prevent all... |
| `apiCheckDefaults` | `array<object>` | (Block Set, Max: 1) (see below for nested schema) |
| `concurrency` | `number` | (Number) Determines the number of checks to run concurrently when triggering the check group via CI/CD or the API. (D... |
| `defaultRuntime` | `array<object>` | (Block List, Max: 1) Sets a default runtime for the group. Used as a fallback when a check belonging to the group has... |
| `enforceAlertSettings` | `array<object>` | (Block List, Max: 1) Enforces alert settings for the whole group. Overrides check configuration. (see below for neste... |
| `enforceLocations` | `array<object>` | (Block List, Max: 1) Enforces public and private locations for the whole group. Overrides check configuration. (see b... |
| `enforceRetryStrategy` | `array<object>` | (Block List, Max: 1) Enforces a retry strategy for the whole group. Overrides check configuration. (see below for nes... |
| `enforceSchedulingStrategy` | `array<object>` | (Block List, Max: 1) Enforces a scheduling strategy for the whole group. Overrides check configuration. (see below fo... |
| `environmentVariable` | `array<object>` | (Block List) Insert environment variables into the runtime environment of checks in the group. Only relevant for chec... |
| `muted` | `boolean` | (Boolean) Determines if any notifications will be sent out when a check in this group fails and/or recovers. Muting t... |
| `name` | `string` | (String) The name of the check group.
The name of the check group. |
| `setupScript` | `array<object>` | (Block List, Max: 1) A script to run in the setup phase of an API check. Runs in addition to the check's own setup sc... |
| `tags` | `array<string>` | (Set of String) Additional tags to append to all checks in the group.
Additional tags to append to all checks in the ... |
| `teardownScript` | `array<object>` | (Block List, Max: 1) A script to run in the teardown phase of an API check. Runs in addition to the check's own teard... |

## Example

```yaml
apiVersion: checks.checkly.m.crossplane.io/v1alpha1
kind: CheckGroupV2
metadata:
  annotations:
    meta.upbound.io/example-id: checks/v1alpha1/checkgroupv2
  labels:
    testing.upbound.io/example-name: just-a-folder
  name: just-a-folder
  namespace: upbound-system
spec:
  forProvider:
    name: Just a Folder

```

## Notes

- Both **namespace-scoped** (`checks.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`checks.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/check_group_v2) for detailed field descriptions.
