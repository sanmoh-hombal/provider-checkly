# Dashboard

Public monitoring dashboard showing check results.

## Details

| | |
|---|---|
| **API Group** | `statuspage.checkly.crossplane.io/v1alpha1` |
| **Kind** | `Dashboard` |
| **Terraform Resource** | [`checkly_dashboard`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/dashboard) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
| `checksPerPage` | `number` | (Number) Determines how many checks to show per page. Possible values are between 1 and 20. (Default 15).
Determines ... |
| `customCss` | `string` | (String) Custom CSS to be applied to the dashboard.
Custom CSS to be applied to the dashboard. |
| `customDomain` | `string` | (String) A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
A custom u... |
| `customUrl` | `string` | (String) A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
A subdomain name under 'checkly... |
| `description` | `string` | (String) HTML  description for the dashboard.
HTML <meta> description for the dashboard. |
| `enableIncidents` | `boolean` | (Boolean) Enable incident support for the dashboard. (Default false).
Enable incident support for the dashboard. (Def... |
| `expandChecks` | `boolean` | (Boolean) Expand or collapse checks on the dashboard. (Default false).
Expand or collapse checks on the dashboard. (D... |
| `favicon` | `string` | (String) A URL pointing to an image file to use as browser favicon.
A URL pointing to an image file to use as browser... |
| `header` | `string` | (String) A piece of text displayed at the top of your dashboard.
A piece of text displayed at the top of your dashboard. |
| `hideTags` | `boolean` | (Boolean) Show or hide the tags on the dashboard. (Default false).
Show or hide the tags on the dashboard. (Default `... |
| `isPrivate` | `boolean` | (Boolean) Set your dashboard as private and generate key.
Set your dashboard as private and generate key. |
| `link` | `string` | (String) A link to for the dashboard logo.
A link to for the dashboard logo. |
| `logo` | `string` | (String) A URL pointing to an image file to use for the dashboard logo.
A URL pointing to an image file to use for th... |
| `paginate` | `boolean` | (Boolean) Determines if pagination is on or off. (Default true).
Determines if pagination is on or off. (Default `tru... |
| `paginationRate` | `number` | (Number) How often to trigger pagination in seconds. Possible values 30, 60 and 300. (Default 60).
How often to trigg... |
| `refreshRate` | `number` | (Number) How often to refresh the dashboard in seconds. Possible values 60, '300' and 600. (Default 60).
How often to... |
| `showCheckRunLinks` | `boolean` | (Boolean) Show or hide check run links on the dashboard. (Default false).
Show or hide check run links on the dashboa... |
| `showHeader` | `boolean` | (Boolean) Show or hide header and description on the dashboard. (Default true).
Show or hide header and description o... |
| `showP95` | `boolean` | (Boolean) Show or hide the P95 stats on the dashboard. (Default true).
Show or hide the P95 stats on the dashboard. (... |
| `showP99` | `boolean` | (Boolean) Show or hide the P99 stats on the dashboard. (Default true).
Show or hide the P99 stats on the dashboard. (... |
| `tags` | `array<string>` | (Set of String) A list of one or more tags that filter which checks to display on the dashboard.
A list of one or mor... |
| `useTagsAndOperator` | `boolean` | (Boolean) Set when to use AND operator for fetching dashboard tags. (Default false).
Set when to use AND operator for... |
| `width` | `string` | (String) Determines whether to use the full screen or focus in the center. Possible values are FULL and 960PX. (Defau... |

## Example

```yaml
apiVersion: statuspage.checkly.m.crossplane.io/v1alpha1
kind: Dashboard
metadata:
  annotations:
    meta.upbound.io/example-id: statuspage/v1alpha1/dashboard
  labels:
    testing.upbound.io/example-name: dashboard_1
  name: dashboard-1
  namespace: upbound-system
spec:
  forProvider:
    customDomain: status.example.com
    customUrl: checkly
    header: Public dashboard
    hideTags: false
    logo: https://www.checklyhq.com/logo.png
    paginate: false
    paginationRate: 30
    refreshRate: 60
    tags:
    - production
    width: FULL
```

## Notes

- Both **namespace-scoped** (`statuspage.checkly.crossplane.io/v1alpha1`) and **cluster-scoped** (`statuspage.checkly.m.crossplane.io/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/dashboard) for detailed field descriptions.
