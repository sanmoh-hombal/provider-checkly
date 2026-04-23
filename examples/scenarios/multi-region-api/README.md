# Scenario: Multi-Region API Monitoring

A check group with 5 API checks spread across multiple regions, plus a weekly
maintenance window to suppress alerts during planned deployments.

## What Gets Created

| Resource | Kind | Purpose |
|----------|------|---------|
| `api-checks` | CheckGroup | Shared config for all API endpoint checks |
| `api-users` | Check (API) | GET /api/users returns 200 |
| `api-products` | Check (API) | GET /api/products returns 200 |
| `api-orders` | Check (API) | GET /api/orders returns 200 |
| `api-search` | Check (API) | GET /api/search?q=test returns 200 |
| `api-health` | Check (API) | GET /healthz returns 200 |
| `deploy-window` | MaintenanceWindow | Suppresses alerts every Sunday 02:00-04:00 UTC |

## Before

Each API endpoint is monitored ad-hoc. No shared configuration. Deployments
trigger false-positive alerts.

## After

All five endpoints share a single group configuration with consistent locations,
tags, and concurrency settings. The weekly maintenance window prevents alert noise
during planned Sunday deployments.

## Usage

```bash
kubectl apply -k examples/scenarios/multi-region-api/
```
