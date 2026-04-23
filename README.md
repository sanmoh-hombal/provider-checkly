<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/crossplane/crossplane/master/banner.png">
    <img alt="provider-checkly" src="https://raw.githubusercontent.com/crossplane/crossplane/master/banner.png" width="600">
  </picture>
</p>

<h1 align="center">provider-checkly</h1>

<p align="center">
  <em>GitOps-native synthetic monitoring — manage <a href="https://checklyhq.com">Checkly</a> resources as Kubernetes custom resources via <a href="https://crossplane.io">Crossplane</a>.</em>
</p>

<p align="center">
  <a href="https://github.com/sanmoh-hombal/provider-checkly/actions/workflows/e2e.yaml"><img alt="E2E" src="https://github.com/sanmoh-hombal/provider-checkly/actions/workflows/e2e.yaml/badge.svg" /></a>
  <a href="https://github.com/sanmoh-hombal/provider-checkly/actions/workflows/pr-title.yaml"><img alt="PR Title" src="https://github.com/sanmoh-hombal/provider-checkly/actions/workflows/pr-title.yaml/badge.svg" /></a>
  <a href="https://github.com/sanmoh-hombal/provider-checkly/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/sanmoh-hombal/provider-checkly?include_prereleases&sort=semver&label=release" /></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/github/license/sanmoh-hombal/provider-checkly" /></a>
  <a href="https://goreportcard.com/report/github.com/sanmoh-hombal/provider-checkly"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/sanmoh-hombal/provider-checkly" /></a>
</p>

---

## What is provider-checkly?

`provider-checkly` is a [Crossplane](https://crossplane.io) provider that reconciles Kubernetes custom resources into [Checkly](https://checklyhq.com) objects — checks, groups, alert channels, dashboards, status pages, and more — via the Checkly REST API.

Commit a YAML manifest, open a PR, merge, and the check exists. Delete the resource and it's cleaned up automatically. Your entire synthetic-monitoring configuration lives in Git, reviewed and versioned like any other infrastructure.

Under the hood the provider is generated with [Upjet](https://github.com/crossplane/upjet) from the official [Terraform provider for Checkly](https://github.com/checkly/terraform-provider-checkly) (`v1.22.0`), so every field and resource stays in sync with upstream — no per-field drift, no manual mapping.

## Installation

> **Prerequisites:** A running Kubernetes cluster with [Crossplane](https://docs.crossplane.io/latest/software/install/) **v2.x** installed.

```bash
# Install the provider
crossplane xpkg install provider ghcr.io/sanmoh-hombal/provider-checkly:v0.1.0
```

Or apply the install manifest directly:

```bash
kubectl apply -f https://raw.githubusercontent.com/sanmoh-hombal/provider-checkly/main/examples/install.yaml
```

## Quickstart

Get from zero to a managed Checkly API check in four steps.

### 1. Create a cluster (skip if you already have one)

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
crossplane xpkg install provider ghcr.io/sanmoh-hombal/provider-checkly:v0.1.0
```

Wait for the provider to become healthy:

```bash
kubectl get providers -w
```

### 4. Configure credentials and create a Check

Create the Checkly credentials secret:

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

Apply the `ProviderConfig`:

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

Create your first API check:

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

Apply everything and verify:

```bash
kubectl apply -f secret.yaml -f providerconfig.yaml -f check.yaml

# Watch the check become ready
kubectl get checks -w
```

## Resources

23 Checkly resources are available, organized across five API groups.

### `checks.checkly.crossplane.io/v1alpha1`

| Crossplane Kind | Terraform Resource | Example |
|---|---|---|
| `Check` | `checkly_check` | [`examples/namespaced/checks/check.yaml`](examples/namespaced/checks/check.yaml) |
| `CheckGroup` | `checkly_check_group` | [`examples/namespaced/checks/checkgroup.yaml`](examples/namespaced/checks/checkgroup.yaml) |
| `CheckGroupV2` | `checkly_check_group_v2` | [`examples/namespaced/checks/checkgroupv2.yaml`](examples/namespaced/checks/checkgroupv2.yaml) |
| `Heartbeat` | `checkly_heartbeat` | [`examples/namespaced/checks/heartbeat.yaml`](examples/namespaced/checks/heartbeat.yaml) |
| `HeartbeatMonitor` | `checkly_heartbeat_monitor` | [`examples/namespaced/checks/heartbeatmonitor.yaml`](examples/namespaced/checks/heartbeatmonitor.yaml) |
| `TCPCheck` | `checkly_tcp_check` | [`examples/namespaced/checks/tcpcheck.yaml`](examples/namespaced/checks/tcpcheck.yaml) |
| `TCPMonitor` | `checkly_tcp_monitor` | [`examples/namespaced/checks/tcpmonitor.yaml`](examples/namespaced/checks/tcpmonitor.yaml) |
| `URLMonitor` | `checkly_url_monitor` | [`examples/namespaced/checks/urlmonitor.yaml`](examples/namespaced/checks/urlmonitor.yaml) |
| `DNSMonitor` | `checkly_dns_monitor` | [`examples/namespaced/checks/dnsmonitor.yaml`](examples/namespaced/checks/dnsmonitor.yaml) |
| `ICMPMonitor` | `checkly_icmp_monitor` | [`examples/namespaced/checks/icmpmonitor.yaml`](examples/namespaced/checks/icmpmonitor.yaml) |
| `PlaywrightCheckSuite` | `checkly_playwright_check_suite` | [`examples/namespaced/checks/playwrightchecksuite.yaml`](examples/namespaced/checks/playwrightchecksuite.yaml) |
| `PlaywrightCodeBundle` | `checkly_playwright_code_bundle` | [`examples/namespaced/checks/playwrightcodebundle.yaml`](examples/namespaced/checks/playwrightcodebundle.yaml) |

### `alerts.checkly.crossplane.io/v1alpha1`

| Crossplane Kind | Terraform Resource | Example |
|---|---|---|
| `AlertChannel` | `checkly_alert_channel` | [`examples/namespaced/alerts/alertchannel.yaml`](examples/namespaced/alerts/alertchannel.yaml) |
| `MaintenanceWindow` | `checkly_maintenance_windows` | [`examples/namespaced/alerts/maintenancewindow.yaml`](examples/namespaced/alerts/maintenancewindow.yaml) |

### `infra.checkly.crossplane.io/v1alpha1`

| Crossplane Kind | Terraform Resource | Example |
|---|---|---|
| `Snippet` | `checkly_snippet` | [`examples/namespaced/infra/snippet.yaml`](examples/namespaced/infra/snippet.yaml) |
| `EnvironmentVariable` | `checkly_environment_variable` | [`examples/namespaced/infra/envvar.yaml`](examples/namespaced/infra/envvar.yaml) |
| `PrivateLocation` | `checkly_private_location` | [`examples/namespaced/infra/privatelocation.yaml`](examples/namespaced/infra/privatelocation.yaml) |
| `ClientCertificate` | `checkly_client_certificate` | [`examples/namespaced/infra/clientcertificate.yaml`](examples/namespaced/infra/clientcertificate.yaml) |

### `statuspage.checkly.crossplane.io/v1alpha1`

| Crossplane Kind | Terraform Resource | Example |
|---|---|---|
| `Dashboard` | `checkly_dashboard` | [`examples/namespaced/statuspage/dashboard.yaml`](examples/namespaced/statuspage/dashboard.yaml) |
| `StatusPage` | `checkly_status_page` | [`examples/namespaced/statuspage/statuspage.yaml`](examples/namespaced/statuspage/statuspage.yaml) |
| `StatusPageService` | `checkly_status_page_service` | [`examples/namespaced/statuspage/service.yaml`](examples/namespaced/statuspage/service.yaml) |

### `triggers.checkly.crossplane.io/v1alpha1`

| Crossplane Kind | Terraform Resource | Example |
|---|---|---|
| `TriggerCheck` | `checkly_trigger_check` | [`examples/namespaced/triggers/triggercheck.yaml`](examples/namespaced/triggers/triggercheck.yaml) |
| `TriggerGroup` | `checkly_trigger_group` | [`examples/namespaced/triggers/triggergroup.yaml`](examples/namespaced/triggers/triggergroup.yaml) |

### `checkly.crossplane.io/v1beta1`

| Kind | Purpose |
|---|---|
| `ProviderConfig` | Credentials and API endpoint configuration |
| `ProviderConfigUsage` | Tracks which managed resources reference a ProviderConfig |

## Cluster vs Namespaced Scope

The provider generates **both** cluster-scoped and namespace-scoped variants of every resource:

- **Cluster-scoped** — `cluster.checks.checkly.crossplane.io/v1alpha1` — classic Crossplane single-tenant model.
- **Namespace-scoped** — `namespaced.checks.checkly.crossplane.io/v1alpha1` — multi-tenant isolation per namespace.

All examples and documentation default to **namespaced** resources, as this is the direction the Crossplane ecosystem is heading. Use cluster-scoped variants when you need backward compatibility or a single-tenant setup.

## Terraform Provider Fallback

For Checkly resources or configurations not yet wrapped by this provider, you can use [`provider-terraform`](https://github.com/upbound/provider-terraform) as a fallback. This lets you write raw HCL blocks as Crossplane resources until native support is added here.

## Documentation

| Resource | Link |
|---|---|
| Example manifests | [`examples/`](examples/) |
| Conformance reports | [`docs/conformance/`](docs/conformance/) |
| Maintainer guides | [`docs/maintainers/`](docs/maintainers/) |
| Checkly API reference | [checklyhq.com/docs/api](https://www.checklyhq.com/docs/api/) |
| Crossplane docs | [docs.crossplane.io](https://docs.crossplane.io/) |
| Upstream Terraform provider | [registry.terraform.io/providers/checkly/checkly](https://registry.terraform.io/providers/checkly/checkly/latest/docs) |

## Contributing

Contributions are welcome! Please read the [Contributing Guide](build/CONTRIBUTING.md) before submitting a pull request.

All commits must include a `Signed-off-by` line (DCO). Use `git commit -s` to add it automatically.

## Releases and Support

| Version | Status | Crossplane | Terraform Provider |
|---|---|---|---|
| `v0.1.0` | Pre-release | v2.x | `checkly/checkly` v1.22.0 |

- **Releases** follow [Semantic Versioning](https://semver.org). See the [Releases](https://github.com/sanmoh-hombal/provider-checkly/releases) page for changelogs.
- **Issues**: Report bugs and request features via [GitHub Issues](https://github.com/sanmoh-hombal/provider-checkly/issues).
- **Security**: Report vulnerabilities privately per the [Security Policy](SECURITY.md).

## License

This project is licensed under the [Apache 2.0 License](LICENSE).

---

<p align="center">
  Built with <a href="https://github.com/crossplane/upjet">Upjet</a> · Powered by <a href="https://crossplane.io">Crossplane</a> · Monitors with <a href="https://checklyhq.com">Checkly</a>
</p>
