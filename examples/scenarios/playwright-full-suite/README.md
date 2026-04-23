# Scenario: Playwright Full Suite

A Playwright code bundle paired with a check suite and a trigger, enabling
browser-based end-to-end testing triggered on demand or on a schedule.

## What Gets Created

| Resource | Kind | Purpose |
|----------|------|---------|
| `e2e-bundle` | PlaywrightCodeBundle | Pre-built Playwright test archive |
| `e2e-suite` | PlaywrightCheckSuite | Runs the bundle on a schedule from eu-west-1 |
| `e2e-trigger` | TriggerCheck | Webhook URL to trigger the suite on demand (e.g., from CI) |

## Before

Browser tests run only in CI. There is no continuous synthetic monitoring of
user-facing flows, and no way to trigger a check from a deployment pipeline.

## After

Playwright tests run every 10 minutes as a Checkly check suite. A trigger URL
lets CI pipelines kick off a check after every deploy. Results are visible in the
Checkly dashboard alongside API checks.

## Usage

```bash
kubectl apply -k examples/scenarios/playwright-full-suite/
```

> **Note:** The `PlaywrightCodeBundle` references a pre-built archive file.
> Replace `./hello-world.tar.gz` with the path to your actual bundle.
