# Troubleshooting

## Common issues

### Provider not becoming healthy

**Symptom:** `kubectl get providers` shows the provider stuck in `Installing` or `Unhealthy`.

**Steps:**

1. Check provider pod logs:

    ```bash
    kubectl -n crossplane-system logs -l pkg.crossplane.io/revision -c package-runtime --tail=50
    ```

2. Verify the provider package is accessible:

    ```bash
    crossplane xpkg list
    ```

3. Ensure Crossplane v2.x is installed — this provider does not support v1.x.

### Resources stuck in "Creating" or not syncing

**Symptom:** `kubectl get <resource>` shows `SYNCED=False` or `READY=False`.

**Steps:**

1. Describe the resource for events and conditions:

    ```bash
    kubectl describe check <name> -n <namespace>
    ```

2. Check the provider logs for API errors:

    ```bash
    kubectl -n crossplane-system logs -l pkg.crossplane.io/revision -c package-runtime --tail=100 | grep -i error
    ```

3. Verify your `ProviderConfig` credentials are correct — a 401 response means an invalid API key.

### ProviderConfig not found

**Symptom:** Resources show `cannot find referenced ProviderConfig`.

**Fix:** Ensure a `ProviderConfig` named `default` exists, or set `spec.providerConfigRef.name` on each resource to match your config name.

### "Forbidden" or 403 errors

**Symptom:** Provider logs show 403 responses from the Checkly API.

**Cause:** The API key lacks permissions for the requested resource type, or the `account_id` doesn't match.

**Fix:** Verify the API key and account ID in your credentials secret. Regenerate the API key from the Checkly dashboard if needed.

### Namespace-scoped vs cluster-scoped confusion

**Symptom:** `kubectl get checks` returns nothing, but resources exist.

**Cause:** You may be using the wrong API group. Namespace-scoped and cluster-scoped resources use different API groups.

**Fix:** Check both:

```bash
# Namespace-scoped
kubectl get checks.checks.checkly.crossplane.io -n <namespace>

# Cluster-scoped
kubectl get checks.checks.checkly.m.crossplane.io
```

## Getting help

- [GitHub Issues](https://github.com/crossplane-contrib/provider-checkly/issues) — bug reports and feature requests.
- [Crossplane Slack](https://slack.crossplane.io/) — community support.
- [Checkly documentation](https://www.checklyhq.com/docs/) — Checkly API and feature reference.
