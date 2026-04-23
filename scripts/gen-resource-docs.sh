#!/usr/bin/env python3
"""Generate resource documentation pages from CRD OpenAPI schemas and examples."""

import os
import sys
import yaml

REPO_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
CRD_DIR = os.path.join(REPO_ROOT, "package", "crds")
EXAMPLES_DIR = os.path.join(REPO_ROOT, "examples-generated", "namespaced")
OUT_DIR = os.path.join(REPO_ROOT, "docs", "docs", "resources")

os.makedirs(OUT_DIR, exist_ok=True)

# Each entry: (crd_basename, kind, group, tf_resource, description, example_group, example_file)
RESOURCES = [
    ("checks.checkly.crossplane.io_checks", "Check", "checks.checkly.crossplane.io", "checkly_check", "API or browser check that monitors an endpoint.", "checks", "check"),
    ("checks.checkly.crossplane.io_checkgroups", "CheckGroup", "checks.checkly.crossplane.io", "checkly_check_group", "Logical grouping of checks with shared settings.", "checks", "checkgroup"),
    ("checks.checkly.crossplane.io_checkgroupv2s", "CheckGroupV2", "checks.checkly.crossplane.io", "checkly_check_group_v2", "V2 check group with enhanced configuration options.", "checks", "checkgroupv2"),
    ("checks.checkly.crossplane.io_heartbeats", "Heartbeat", "checks.checkly.crossplane.io", "checkly_heartbeat", "Heartbeat (cron/ping) check that expects periodic pings.", "checks", "heartbeat"),
    ("checks.checkly.crossplane.io_heartbeatmonitors", "HeartbeatMonitor", "checks.checkly.crossplane.io", "checkly_heartbeat_monitor", "Monitor that watches for periodic heartbeat pings.", "checks", "heartbeatmonitor"),
    ("checks.checkly.crossplane.io_tcpchecks", "TCPCheck", "checks.checkly.crossplane.io", "checkly_tcp_check", "TCP connectivity check against a host and port.", "checks", "tcpcheck"),
    ("checks.checkly.crossplane.io_tcpmonitors", "TCPMonitor", "checks.checkly.crossplane.io", "checkly_tcp_monitor", "TCP port monitor with configurable thresholds.", "checks", "tcpmonitor"),
    ("checks.checkly.crossplane.io_urlmonitors", "URLMonitor", "checks.checkly.crossplane.io", "checkly_url_monitor", "URL availability monitor with response assertions.", "checks", "urlmonitor"),
    ("checks.checkly.crossplane.io_dnsmonitors", "DNSMonitor", "checks.checkly.crossplane.io", "checkly_dns_monitor", "DNS resolution monitor for domain records.", "checks", "dnsmonitor"),
    ("checks.checkly.crossplane.io_icmpmonitors", "ICMPMonitor", "checks.checkly.crossplane.io", "checkly_icmp_monitor", "ICMP (ping) monitor for host reachability.", "checks", "icmpmonitor"),
    ("checks.checkly.crossplane.io_playwrightchecksuites", "PlaywrightCheckSuite", "checks.checkly.crossplane.io", "checkly_playwright_check_suite", "Playwright test suite that runs browser-based checks.", "checks", "playwrightchecksuite"),
    ("checks.checkly.crossplane.io_playwrightcodebundles", "PlaywrightCodeBundle", "checks.checkly.crossplane.io", "checkly_playwright_code_bundle", "Uploadable Playwright code bundle for browser checks.", "checks", "playwrightcodebundle"),
    ("alerts.checkly.crossplane.io_alertchannels", "AlertChannel", "alerts.checkly.crossplane.io", "checkly_alert_channel", "Notification channel (email, Slack, SMS, webhook, PagerDuty, etc.).", "alerts", "alertchannel"),
    ("alerts.checkly.crossplane.io_maintenancewindows", "MaintenanceWindow", "alerts.checkly.crossplane.io", "checkly_maintenance_windows", "Scheduled maintenance window that suppresses alerts.", "alerts", "maintenancewindow"),
    ("infra.checkly.crossplane.io_snippets", "Snippet", "infra.checkly.crossplane.io", "checkly_snippet", "Reusable code snippet for setup/teardown scripts.", "infra", "snippet"),
    ("infra.checkly.crossplane.io_environmentvariables", "EnvironmentVariable", "infra.checkly.crossplane.io", "checkly_environment_variable", "Global environment variable available to all checks.", "infra", "environmentvariable"),
    ("infra.checkly.crossplane.io_privatelocations", "PrivateLocation", "infra.checkly.crossplane.io", "checkly_private_location", "Private check runner location inside your network.", "infra", "privatelocation"),
    ("infra.checkly.crossplane.io_clientcertificates", "ClientCertificate", "infra.checkly.crossplane.io", "checkly_client_certificate", "Client TLS certificate for mTLS check endpoints.", "infra", "clientcertificate"),
    ("statuspage.checkly.crossplane.io_dashboards", "Dashboard", "statuspage.checkly.crossplane.io", "checkly_dashboard", "Public monitoring dashboard showing check results.", "statuspage", "dashboard"),
    ("statuspage.checkly.crossplane.io_statuspages", "StatusPage", "statuspage.checkly.crossplane.io", "checkly_status_page", "Public status page with service health indicators.", "statuspage", "statuspage"),
    ("statuspage.checkly.crossplane.io_statuspageservices", "StatusPageService", "statuspage.checkly.crossplane.io", "checkly_status_page_service", "Service entry linked to a status page.", "statuspage", "statuspageservice"),
    ("triggers.checkly.crossplane.io_triggerchecks", "TriggerCheck", "triggers.checkly.crossplane.io", "checkly_trigger_check", "Webhook trigger that runs a single check on demand.", "triggers", "triggercheck"),
    ("triggers.checkly.crossplane.io_triggergroups", "TriggerGroup", "triggers.checkly.crossplane.io", "checkly_trigger_group", "Webhook trigger that runs a check group on demand.", "triggers", "triggergroup"),
]


def extract_spec_fields(crd_file):
    """Extract top-level forProvider fields from CRD OpenAPI schema."""
    try:
        with open(crd_file) as f:
            crd = yaml.safe_load(f)
    except Exception:
        return ""

    versions = crd.get("spec", {}).get("versions", [])
    if not versions:
        return ""

    schema = versions[0].get("schema", {}).get("openAPIV3Schema", {})
    spec_props = (
        schema.get("properties", {})
        .get("spec", {})
        .get("properties", {})
        .get("forProvider", {})
        .get("properties", {})
    )
    if not spec_props:
        return ""

    rows = []
    for name in sorted(spec_props.keys()):
        prop = spec_props[name]
        ptype = prop.get("type", "object")
        if "items" in prop:
            inner = prop["items"].get("type", "object")
            ptype = f"array<{inner}>"
        desc = prop.get("description", "")
        if len(desc) > 120:
            desc = desc[:117] + "..."
        rows.append(f"| `{name}` | `{ptype}` | {desc} |")

    return "\n".join(rows)


def extract_first_document(example_file):
    """Extract the first YAML document from a multi-document file."""
    try:
        with open(example_file) as f:
            lines = []
            for i, line in enumerate(f):
                if i > 0 and line.strip() == "---":
                    break
                lines.append(line.rstrip())
            return "\n".join(lines)
    except Exception:
        return ""


count = 0
for crd_base, kind, group, tf_resource, description, eg_group, eg_file in RESOURCES:
    crd_file = os.path.join(CRD_DIR, f"{crd_base}.yaml")
    example_path = os.path.join(EXAMPLES_DIR, eg_group, "v1alpha1", f"{eg_file}.yaml")
    out_file = os.path.join(OUT_DIR, f"{eg_file}.md")

    # Terraform docs resource name (strip checkly_ prefix)
    tf_doc_name = tf_resource.removeprefix("checkly_")

    # Cluster-scoped group variant
    cluster_group = group.replace("checkly.crossplane", "checkly.m.crossplane")

    # Extract fields and example
    spec_fields = extract_spec_fields(crd_file)
    if not spec_fields:
        spec_fields = "| | | *See the CRD for full schema details.* |"

    example_yaml = extract_first_document(example_path)
    if not example_yaml:
        example_yaml = f"# See examples-generated/{eg_group}/v1alpha1/{eg_file}.yaml"

    content = f"""# {kind}

{description}

## Details

| | |
|---|---|
| **API Group** | `{group}/v1alpha1` |
| **Kind** | `{kind}` |
| **Terraform Resource** | [`{tf_resource}`](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/{tf_doc_name}) |

## `spec.forProvider` fields

| Field | Type | Description |
|-------|------|-------------|
{spec_fields}

## Example

```yaml
{example_yaml}
```

## Notes

- Both **namespace-scoped** (`{group}/v1alpha1`) and **cluster-scoped** (`{cluster_group}/v1alpha1`) variants are available.
- See the [Terraform docs](https://registry.terraform.io/providers/checkly/checkly/latest/docs/resources/{tf_doc_name}) for detailed field descriptions.
"""

    with open(out_file, "w") as f:
        f.write(content)
    count += 1

print(f"Generated {count} resource pages in {OUT_DIR}")
