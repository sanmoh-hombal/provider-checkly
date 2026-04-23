# Scenario: Status Page for a SaaS Product

A public status page backed by four services, each with its own health check.
Customers visit the status page to see real-time availability of core services.

## What Gets Created

| Resource | Kind | Purpose |
|----------|------|---------|
| `saas-status-page` | StatusPage | Public status page at `status-saas-crossplane` |
| `svc-api` | StatusPageService | "API" service on the status page |
| `svc-dashboard` | StatusPageService | "Dashboard" service on the status page |
| `svc-auth` | StatusPageService | "Authentication" service on the status page |
| `svc-webhooks` | StatusPageService | "Webhooks" service on the status page |
| `check-api` | Check (API) | Health check backing the API service |
| `check-dashboard` | Check (API) | Health check backing the Dashboard service |
| `check-auth` | Check (API) | Health check backing the Auth service |
| `check-webhooks` | Check (API) | Health check backing the Webhooks service |

## Before

No public status page. Customers open support tickets to ask "is the service
down?" during incidents.

## After

A branded status page shows real-time health of four core services. Each service
is backed by an automated health check. Customers self-serve during incidents.

## Usage

```bash
kubectl apply -k examples/scenarios/status-page-for-saas/
```
