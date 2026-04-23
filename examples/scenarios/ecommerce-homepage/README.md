# Scenario: E-Commerce Homepage Monitoring

Monitor an e-commerce homepage with both an API health check and a browser check,
grouped together, with three alert channels (email, Slack, PagerDuty).

## What Gets Created

| Resource | Kind | Purpose |
|----------|------|---------|
| `homepage-group` | CheckGroup | Groups both checks with shared locations and tags |
| `homepage-api` | Check (API) | Verifies the homepage returns HTTP 200 |
| `homepage-browser` | Check (BROWSER) | Loads the homepage in a real browser and asserts the title |
| `alert-email` | AlertChannel | Sends failure/recovery alerts via email |
| `alert-slack` | AlertChannel | Posts to a Slack channel on failure/recovery |
| `alert-pagerduty` | AlertChannel | Pages the on-call engineer via PagerDuty |

## Before

No monitoring exists for the storefront. Outages go undetected until customers
report them.

## After

The homepage is checked every 5 minutes from two regions. Any failure triggers
email, Slack, and PagerDuty alerts simultaneously. Recovery notifications confirm
when the issue is resolved.

## Usage

```bash
kubectl apply -k examples/scenarios/ecommerce-homepage/
```
