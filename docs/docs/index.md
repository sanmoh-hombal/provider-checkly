# provider-checkly

**GitOps-native synthetic monitoring** — manage [Checkly](https://checklyhq.com) resources as Kubernetes custom resources via [Crossplane](https://crossplane.io).

## Overview

`provider-checkly` is a Crossplane provider generated with [Upjet](https://github.com/crossplane/upjet) from the official [Terraform provider for Checkly](https://github.com/checkly/terraform-provider-checkly) (`v1.22.0`).

Commit a YAML manifest, open a PR, merge, and the check exists. Delete the resource and it's cleaned up automatically. Your entire synthetic-monitoring configuration lives in Git, reviewed and versioned like any other infrastructure.

## Features

- **23 managed resources** across checks, alerts, infrastructure, status pages, and triggers.
- **Cluster-scoped and namespace-scoped** variants for every resource.
- **Full Checkly API coverage** — API checks, browser checks, heartbeats, TCP/DNS/ICMP monitors, Playwright suites, alert channels, dashboards, status pages, and more.
- **GitOps-ready** — declarative, drift-detected, and reconciled continuously.

## Quick links

| | |
|---|---|
| [Quickstart](quickstart.md) | From zero to a managed Checkly check in four steps |
| [Resources](resources/index.md) | All 23 resource types with examples |
| [Configuration](configuration.md) | ProviderConfig, credentials, and tuning |
| [Troubleshooting](troubleshooting.md) | Common issues and fixes |
| [Conformance](conformance.md) | Crossplane conformance report |

## Requirements

- Kubernetes cluster
- [Crossplane](https://docs.crossplane.io/latest/software/install/) **v2.x**
- A [Checkly](https://checklyhq.com) account with an API key
