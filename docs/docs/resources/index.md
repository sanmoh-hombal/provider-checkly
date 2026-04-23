# Resources

23 Checkly resources are available, organized across five API groups. Each resource has both cluster-scoped and namespace-scoped variants.

## Checks (`checks.checkly.crossplane.io/v1alpha1`)

| Kind | Terraform Resource | Description |
|------|-------------------|-------------|
| [Check](check.md) | `checkly_check` | API or browser check |
| [CheckGroup](checkgroup.md) | `checkly_check_group` | Logical grouping of checks |
| [CheckGroupV2](checkgroupv2.md) | `checkly_check_group_v2` | V2 check group with enhanced features |
| [Heartbeat](heartbeat.md) | `checkly_heartbeat` | Heartbeat (cron/ping) check |
| [HeartbeatMonitor](heartbeatmonitor.md) | `checkly_heartbeat_monitor` | Heartbeat monitor |
| [TCPCheck](tcpcheck.md) | `checkly_tcp_check` | TCP connectivity check |
| [TCPMonitor](tcpmonitor.md) | `checkly_tcp_monitor` | TCP port monitor |
| [URLMonitor](urlmonitor.md) | `checkly_url_monitor` | URL availability monitor |
| [DNSMonitor](dnsmonitor.md) | `checkly_dns_monitor` | DNS resolution monitor |
| [ICMPMonitor](icmpmonitor.md) | `checkly_icmp_monitor` | ICMP (ping) monitor |
| [PlaywrightCheckSuite](playwrightchecksuite.md) | `checkly_playwright_check_suite` | Playwright test suite |
| [PlaywrightCodeBundle](playwrightcodebundle.md) | `checkly_playwright_code_bundle` | Playwright code bundle |

## Alerts (`alerts.checkly.crossplane.io/v1alpha1`)

| Kind | Terraform Resource | Description |
|------|-------------------|-------------|
| [AlertChannel](alertchannel.md) | `checkly_alert_channel` | Notification channel (email, Slack, webhook, etc.) |
| [MaintenanceWindow](maintenancewindow.md) | `checkly_maintenance_windows` | Scheduled maintenance window |

## Infrastructure (`infra.checkly.crossplane.io/v1alpha1`)

| Kind | Terraform Resource | Description |
|------|-------------------|-------------|
| [Snippet](snippet.md) | `checkly_snippet` | Reusable code snippet |
| [EnvironmentVariable](environmentvariable.md) | `checkly_environment_variable` | Global environment variable |
| [PrivateLocation](privatelocation.md) | `checkly_private_location` | Private check runner location |
| [ClientCertificate](clientcertificate.md) | `checkly_client_certificate` | Client TLS certificate |

## Status Pages (`statuspage.checkly.crossplane.io/v1alpha1`)

| Kind | Terraform Resource | Description |
|------|-------------------|-------------|
| [Dashboard](dashboard.md) | `checkly_dashboard` | Public monitoring dashboard |
| [StatusPage](statuspage.md) | `checkly_status_page` | Public status page |
| [StatusPageService](statuspageservice.md) | `checkly_status_page_service` | Service entry on a status page |

## Triggers (`triggers.checkly.crossplane.io/v1alpha1`)

| Kind | Terraform Resource | Description |
|------|-------------------|-------------|
| [TriggerCheck](triggercheck.md) | `checkly_trigger_check` | Webhook trigger for a single check |
| [TriggerGroup](triggergroup.md) | `checkly_trigger_group` | Webhook trigger for a check group |
