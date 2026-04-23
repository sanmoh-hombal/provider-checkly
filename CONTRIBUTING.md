# Contributing to provider-checkly

Thank you for your interest in contributing! This guide walks you through
setting up your development environment, making changes, and opening a pull
request.

## Prerequisites

| Tool | Version | Notes |
|------|---------|-------|
| Go | 1.25+ | `go version` |
| Docker | latest stable | For `kind` clusters and image builds |
| kind | v0.20+ | Local Kubernetes clusters |
| kubectl | 1.28+ | Cluster interaction |
| kuttl | v0.15+ | E2E test harness |
| Crossplane CLI | v2.x | `crossplane xpkg build` |
| golangci-lint | v1.62+ | Linting |

## Setup

```bash
# 1. Fork and clone the repository
git clone https://github.com/<your-user>/provider-checkly.git
cd provider-checkly

# 2. Initialise submodules (pulls build helpers and the upstream TF provider)
make submodules

# 3. Generate CRDs, deepcopy methods, and cross-references
make generate

# 4. Build provider binaries
make build
```

After `make build` succeeds you have a working local binary. Open a trivial PR
(typo fix, doc improvement) to validate the full flow before tackling bigger
changes.

## Build and test

```bash
# Unit / integration tests
go test ./...

# Lint
golangci-lint run

# E2E smoke test (requires Docker for kind + CHECKLY_API_KEY)
scripts/e2e.sh

# Uptest (kuttl-based E2E against a real cluster)
make uptest
```

## Style

- **Format:** `gofmt` + `goimports`. CI will reject unformatted code.
- **Lint:** `golangci-lint run` — fix all findings before pushing.
- **Conventional commits:** Every commit message must follow
  [Conventional Commits](https://www.conventionalcommits.org/):

  ```
  feat(checks): configure checkly_check resource
  fix(alertchannel): correct external-name template
  docs: update CONTRIBUTING with setup steps
  ```

  The PR title linter enforces this on pull requests.

## Certificate of Origin (DCO)

By contributing you agree to the
[Developer Certificate of Origin](https://developercertificate.org/). Sign off
every commit with `-s`:

```bash
git commit -s -m 'feat(checks): add check resource'
```

A [DCO bot](https://github.com/apps/dco) enforces this on every pull request.
If you forgot `-s`, amend and force-push:

```bash
git commit --amend -s
git push --force-with-lease
```

## Adding a resource

> **Recommended:** Use the `crossplane-add-resource` Claude Code skill
> (`~/.claude/skills/crossplane-add-resource/SKILL.md`) to automate these
> steps. Invoke it with:
>
> ```
> /crossplane-add-resource
> ```
>
> The skill will prompt you for the required inputs and handle all file
> creation, code generation, and validation.

If you prefer to work manually, the steps are:

1. **Gather inputs.** You need:
   - `TF_RESOURCE` — Terraform resource name (e.g. `checkly_check`).
   - `KIND` — Crossplane kind (e.g. `Check`).
   - `SHORT_GROUP` — API short group (e.g. `checks`).
   - `EXTERNAL_NAME` — naming strategy: `identifier-from-provider` (default),
     `templated-string`, or `name-as-identifier`.
   - `REFERENCES` — cross-resource references (optional).
   - `SENSITIVE_FIELDS` — fields to mark sensitive (optional).
   - `CONNECTION_DETAILS` — fields to surface in connection secrets (optional).

2. **Check the resource isn't already registered:**

   ```bash
   grep -rn '"checkly_check"' config/ || true
   ```

3. **Add external-name entry** in `config/external_name.go`:

   ```go
   "checkly_check": config.IdentifierFromProvider,
   ```

4. **Add per-scope configurators** in both
   `config/cluster/<short_group>/config.go` and
   `config/namespaced/<short_group>/config.go`. Set `ShortGroup`, `Kind`,
   references, sensitive fields, and connection details.

5. **Regenerate:**

   ```bash
   make generate
   ```

6. **Write a registration test** in
   `config/cluster/<short_group>/<resource>_test.go` that asserts the resource
   is registered with the correct group, kind, references, and sensitive
   fields.

7. **Write a minimal example manifest** at
   `examples/namespaced/<short_group>/<resource>.yaml`. Use
   `examples-generated/` as a starting point for field values.

8. **Run checks:**

   ```bash
   go build ./...
   go test ./config/...
   kubectl apply --dry-run=client -f examples/namespaced/<short_group>/<resource>.yaml
   ```

9. **Commit:**

   ```bash
   git add config/ examples/
   git commit -s -m 'feat(<short_group>): configure <tf_resource>'
   ```

## Bumping the Checkly Terraform provider

1. Update `TERRAFORM_PROVIDER_VERSION` and
   `TERRAFORM_NATIVE_PROVIDER_BINARY` in the top-level `Makefile`.
2. Run `make generate` and review the diff for new/changed resources.
3. Add or update example manifests and tests for any changed resources.
4. Open a PR with the conventional commit prefix `build:` or `feat:`.

## Pull request checklist

- [ ] Conventional commit title (enforced by CI)
- [ ] DCO sign-off on all commits
- [ ] `make generate` is idempotent (running it twice produces no diff)
- [ ] `go test ./...` passes
- [ ] `golangci-lint run` clean
- [ ] Example manifests added/updated for new resources
- [ ] Manual E2E on sandbox (if behavioural change)

## Security issues

Do **not** open a public issue for security vulnerabilities. Follow the process
in [SECURITY.md](SECURITY.md).
