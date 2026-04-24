<p align="center">
  <br />
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/crossplane/artwork/master/logo/logo-horizontal-whitetext.png">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/crossplane/artwork/master/logo/logo-horizontal-bluetext.svg">
    <img alt="Crossplane" src="https://raw.githubusercontent.com/crossplane/artwork/master/logo/logo-horizontal-bluetext.svg" width="360">
  </picture>
  <br /><br />
</p>

<h1 align="center">
  provider-checkly
</h1>

<h4 align="center">
  GitOps-native synthetic monitoring &mdash; manage <a href="https://checklyhq.com">Checkly</a> resources<br />
  as Kubernetes custom resources via <a href="https://crossplane.io">Crossplane</a>.
</h4>

<p align="center">
  <a href="https://github.com/sanmoh-hombal/provider-checkly/actions/workflows/ci.yml"><img alt="CI" src="https://img.shields.io/github/actions/workflow/status/sanmoh-hombal/provider-checkly/ci.yml?branch=main&style=for-the-badge&label=CI&logo=githubactions&logoColor=white" /></a>
  <a href="https://github.com/sanmoh-hombal/provider-checkly/actions/workflows/e2e.yaml"><img alt="E2E" src="https://img.shields.io/github/actions/workflow/status/sanmoh-hombal/provider-checkly/e2e.yaml?branch=main&style=for-the-badge&label=E2E&logo=githubactions&logoColor=white" /></a>
  <a href="https://codecov.io/gh/sanmoh-hombal/provider-checkly"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/sanmoh-hombal/provider-checkly?style=for-the-badge&logo=codecov&logoColor=white" /></a>
  <a href="https://github.com/sanmoh-hombal/provider-checkly/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/sanmoh-hombal/provider-checkly?include_prereleases&sort=semver&style=for-the-badge&label=release&logo=github&color=blue" /></a>
  <a href="https://goreportcard.com/report/github.com/sanmoh-hombal/provider-checkly"><img alt="Go Report" src="https://img.shields.io/badge/go%20report-A+-brightgreen?style=for-the-badge&logo=go&logoColor=white" /></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/github/license/sanmoh-hombal/provider-checkly?style=for-the-badge&color=blue" /></a>
</p>

<br />

<p align="center">
  <a href="#-quickstart">Quickstart</a> &bull;
  <a href="#-resources">Resources</a> &bull;
  <a href="#-examples">Examples</a> &bull;
  <a href="#-contributing">Contributing</a> &bull;
  <a href="#-documentation">Documentation</a>
</p>

<br />

---

<br />

## Overview

**provider-checkly** is a [Crossplane](https://crossplane.io) provider that reconciles Kubernetes custom resources into [Checkly](https://checklyhq.com) objects &mdash; checks, groups, alert channels, dashboards, status pages, and more.

```
Commit a YAML manifest  ->  Open a PR  ->  Merge  ->  Check exists in Checkly
Delete the resource      ->  Check is cleaned up automatically
```

Your entire synthetic-monitoring configuration lives in Git, reviewed and versioned like any other infrastructure.

Under the hood, the provider is generated with [Upjet](https://github.com/crossplane/upjet) from the official [Terraform provider for Checkly](https://github.com/checkly/terraform-provider-checkly) (`v1.22.0`), so every field and resource stays in sync with upstream &mdash; no per-field drift, no manual mapping.

<br />

## Installation

> **Prerequisites:** A running Kubernetes cluster with [Crossplane](https://docs.crossplane.io/latest/software/install/) **v2.x** installed.

```bash
crossplane xpkg install provider ghcr.io/sanmoh-hombal/provider-checkly:v0.1.1
```

Or apply the install manifest directly:

```bash
kubectl apply -f https://raw.githubusercontent.com/sanmoh-hombal/provider-checkly/main/examples/install.yaml
```

<br />

## Quickstart

Get from zero to a managed Checkly API check in four steps.

### 1. Create a cluster

> Skip if you already have one.

```bash
kind create cluster --name crossplane-demo
```

### 2. Install Crossplane

```bash
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane crossplane-stable/crossplane \
  --namespace crossplane-system \
  --create-namespace
```

### 3. Install provider-checkly

```bash
crossplane xpkg install provider ghcr.io/sanmoh-hombal/provider-checkly:v0.1.1
```

Wait for the provider to become healthy:

```bash
kubectl get providers -w
```

### 4. Configure credentials and create a Check

<details>
<summary><b>Create the Checkly credentials secret</b></summary>

```yaml
# secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: checkly-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<your-checkly-api-key>",
      "account_id": "<your-checkly-account-id>",
      "api_url": "https://api.checklyhq.com"
    }
```

</details>

<details>
<summary><b>Apply the ProviderConfig</b></summary>

```yaml
# providerconfig.yaml
apiVersion: checkly.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: checkly-creds
      key: credentials
```

</details>

<details>
<summary><b>Create your first API check</b></summary>

```yaml
# check.yaml
apiVersion: checks.checkly.crossplane.io/v1alpha1
kind: Check
metadata:
  name: homepage-api
  namespace: default
spec:
  forProvider:
    name: homepage-api
    activated: true
    frequency: 5
    type: API
    locations:
      - eu-west-1
      - us-east-1
    tags:
      - managed-by-crossplane
    request:
      - method: GET
        url: https://example.com
        assertion:
          - source: STATUS_CODE
            property: ""
            comparison: EQUALS
            target: "200"
  providerConfigRef:
    name: default
```

</details>

Apply everything and verify:

```bash
kubectl apply -f secret.yaml -f providerconfig.yaml -f check.yaml

# Watch the check become ready
kubectl get checks -w
```

<br />

## Resources

**23** Checkly resources are available, organized across **five** API groups.

<details open>
<summary><b>checks</b> &mdash; <code>checks.checkly.crossplane.io/v1alpha1</code></summary>

<br />

| Kind | Terraform Resource | Example |
|:-----|:-------------------|:--------|
| `Check` | `checkly_check` | [`check.yaml`](examples/namespaced/checks/check.yaml) |
| `CheckGroup` | `checkly_check_group` | [`checkgroup.yaml`](examples/namespaced/checks/checkgroup.yaml) |
| `CheckGroupV2` | `checkly_check_group_v2` | [`checkgroupv2.yaml`](examples/namespaced/checks/checkgroupv2.yaml) |
| `Heartbeat` | `checkly_heartbeat` | [`heartbeat.yaml`](examples/namespaced/checks/heartbeat.yaml) |
| `HeartbeatMonitor` | `checkly_heartbeat_monitor` | [`heartbeatmonitor.yaml`](examples/namespaced/checks/heartbeatmonitor.yaml) |
| `TCPCheck` | `checkly_tcp_check` | [`tcpcheck.yaml`](examples/namespaced/checks/tcpcheck.yaml) |
| `TCPMonitor` | `checkly_tcp_monitor` | [`tcpmonitor.yaml`](examples/namespaced/checks/tcpmonitor.yaml) |
| `URLMonitor` | `checkly_url_monitor` | [`urlmonitor.yaml`](examples/namespaced/checks/urlmonitor.yaml) |
| `DNSMonitor` | `checkly_dns_monitor` | [`dnsmonitor.yaml`](examples/namespaced/checks/dnsmonitor.yaml) |
| `ICMPMonitor` | `checkly_icmp_monitor` | [`icmpmonitor.yaml`](examples/namespaced/checks/icmpmonitor.yaml) |
| `PlaywrightCheckSuite` | `checkly_playwright_check_suite` | [`playwrightchecksuite.yaml`](examples/namespaced/checks/playwrightchecksuite.yaml) |
| `PlaywrightCodeBundle` | `checkly_playwright_code_bundle` | [`playwrightcodebundle.yaml`](examples/namespaced/checks/playwrightcodebundle.yaml) |

</details>

<details>
<summary><b>alerts</b> &mdash; <code>alerts.checkly.crossplane.io/v1alpha1</code></summary>

<br />

| Kind | Terraform Resource | Example |
|:-----|:-------------------|:--------|
| `AlertChannel` | `checkly_alert_channel` | [`alertchannel.yaml`](examples/namespaced/alerts/alertchannel.yaml) |
| `MaintenanceWindow` | `checkly_maintenance_windows` | [`maintenancewindow.yaml`](examples/namespaced/alerts/maintenancewindow.yaml) |

</details>

<details>
<summary><b>infra</b> &mdash; <code>infra.checkly.crossplane.io/v1alpha1</code></summary>

<br />

| Kind | Terraform Resource | Example |
|:-----|:-------------------|:--------|
| `Snippet` | `checkly_snippet` | [`snippet.yaml`](examples/namespaced/infra/snippet.yaml) |
| `EnvironmentVariable` | `checkly_environment_variable` | [`envvar.yaml`](examples/namespaced/infra/envvar.yaml) |
| `PrivateLocation` | `checkly_private_location` | [`privatelocation.yaml`](examples/namespaced/infra/privatelocation.yaml) |
| `ClientCertificate` | `checkly_client_certificate` | [`clientcertificate.yaml`](examples/namespaced/infra/clientcertificate.yaml) |

</details>

<details>
<summary><b>statuspage</b> &mdash; <code>statuspage.checkly.crossplane.io/v1alpha1</code></summary>

<br />

| Kind | Terraform Resource | Example |
|:-----|:-------------------|:--------|
| `Dashboard` | `checkly_dashboard` | [`dashboard.yaml`](examples/namespaced/statuspage/dashboard.yaml) |
| `StatusPage` | `checkly_status_page` | [`statuspage.yaml`](examples/namespaced/statuspage/statuspage.yaml) |
| `StatusPageService` | `checkly_status_page_service` | [`service.yaml`](examples/namespaced/statuspage/service.yaml) |

</details>

<details>
<summary><b>triggers</b> &mdash; <code>triggers.checkly.crossplane.io/v1alpha1</code></summary>

<br />

| Kind | Terraform Resource | Example |
|:-----|:-------------------|:--------|
| `TriggerCheck` | `checkly_trigger_check` | [`triggercheck.yaml`](examples/namespaced/triggers/triggercheck.yaml) |
| `TriggerGroup` | `checkly_trigger_group` | [`triggergroup.yaml`](examples/namespaced/triggers/triggergroup.yaml) |

</details>

<details>
<summary><b>provider</b> &mdash; <code>checkly.crossplane.io/v1beta1</code></summary>

<br />

| Kind | Purpose |
|:-----|:--------|
| `ProviderConfig` | Credentials and API endpoint configuration |
| `ProviderConfigUsage` | Tracks which managed resources reference a ProviderConfig |

</details>

<br />

## Cluster vs Namespaced Scope

The provider generates **both** cluster-scoped and namespace-scoped variants of every resource:

| Scope | API Group Pattern | Use Case |
|:------|:------------------|:---------|
| **Cluster** | `checks.checkly.crossplane.io` | Classic Crossplane single-tenant model |
| **Namespaced** | `checks.checkly.m.crossplane.io` | Multi-tenant isolation per namespace |

All examples default to **namespaced** resources, as this is the direction the Crossplane ecosystem is heading.

<br />

## Examples

Multi-resource scenario walkthroughs live in [`examples/scenarios/`](examples/scenarios/):

| Scenario | Description |
|:---------|:------------|
| [`ecommerce-homepage/`](examples/scenarios/ecommerce-homepage/) | API + browser check, grouped, with email/Slack/PagerDuty alerts |
| [`multi-region-api/`](examples/scenarios/multi-region-api/) | 5 API checks across regions with a weekly maintenance window |
| [`status-page-for-saas/`](examples/scenarios/status-page-for-saas/) | Status page + 4 services, each backed by health checks |
| [`playwright-full-suite/`](examples/scenarios/playwright-full-suite/) | Playwright code bundle + check suite + trigger |
| [`gitops-setup/`](examples/scenarios/gitops-setup/) | Kustomization composing ProviderConfig + snippet + 10 checks |

Each folder contains a README and can be applied with `kubectl apply -k`.

<br />

## Documentation

| Resource | Link |
|:---------|:-----|
| Example manifests | [`examples/`](examples/) |
| Conformance reports | [`docs/conformance/`](docs/conformance/) |
| Maintainer guides | [`docs/maintainers/`](docs/maintainers/) |
| Checkly API reference | [checklyhq.com/docs/api](https://www.checklyhq.com/docs/api/) |
| Crossplane docs | [docs.crossplane.io](https://docs.crossplane.io/) |
| Upstream TF provider | [registry.terraform.io](https://registry.terraform.io/providers/checkly/checkly/latest/docs) |

<br />

## Terraform Provider Fallback

For Checkly resources not yet wrapped by this provider, you can use [`provider-terraform`](https://github.com/upbound/provider-terraform) as a fallback &mdash; write raw HCL blocks as Crossplane resources until native support is added.

<br />

## Contributing

Contributions are welcome! Please read the [Contributing Guide](CONTRIBUTING.md) before submitting a pull request.

All commits must include a `Signed-off-by` line (DCO). Use `git commit -s` to add it automatically.

<br />

## Releases and Support

| Version | Status | Crossplane | Terraform Provider |
|:--------|:-------|:-----------|:-------------------|
| `v0.1.1` | Latest | v2.x | `checkly/checkly` v1.22.0 |
| `v0.1.0` | Pre-release | v2.x | `checkly/checkly` v1.22.0 |

Releases follow [Semantic Versioning](https://semver.org). See the [Releases](https://github.com/sanmoh-hombal/provider-checkly/releases) page for changelogs.

**Issues** &mdash; Report bugs and request features via [GitHub Issues](https://github.com/sanmoh-hombal/provider-checkly/issues).
**Security** &mdash; Report vulnerabilities privately per the [Security Policy](SECURITY.md).

<br />

## License

This project is licensed under the [Apache 2.0 License](LICENSE).

---

<p align="center">
  <sub>Built with <a href="https://github.com/crossplane/upjet">Upjet</a> &bull; Powered by <a href="https://crossplane.io">Crossplane</a> &bull; Monitors with <a href="https://checklyhq.com">Checkly</a></sub>
</p>
