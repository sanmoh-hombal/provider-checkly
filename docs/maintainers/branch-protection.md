# Branch Protection — `main`

Settings applied in **GitHub UI** → Settings → Branches → Branch protection rules.

## Rule: `main`

| Setting | Value |
|---------|-------|
| Require a pull request before merging | Yes |
| Required approvals | 1 |
| Require status checks to pass before merging | Yes |
| Required status checks | `reviewable`, `PR Title` |
| Require branches to be up-to-date before merging | Yes |
| Require linear history | Yes |
| Restrict who can push to matching branches | `sanmoh-hombal` |

## Why these settings

- **1 approval** — every change gets a second pair of eyes without creating a bottleneck on a small team.
- **Status checks** — `reviewable` is a summary job in `ci.yml` that gates on all CI jobs (lint, check-diff, unit-tests, local-deploy, report-breaking-changes). `PR Title` enforces conventional commits. Gating on `reviewable` instead of individual jobs means adding or renaming CI jobs won't break the branch protection rule.
- **Up-to-date branch** — prevents merging stale branches that haven't been rebased against the latest `main`.
- **Linear history** — squash/rebase only; keeps `git log` readable and `git bisect` practical.
- **Restricted pushes** — only `sanmoh-hombal` can push directly (for emergencies); everyone else goes through PRs.

## Applying changes

All settings are managed manually in the GitHub UI. If you need to update them:

1. Go to **Settings → Branches → `main`** in the GitHub repo.
2. Edit the branch protection rule.
3. Update this document to match.
