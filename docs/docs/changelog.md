# Changelog

All notable changes to this project are documented here. This project follows [Semantic Versioning](https://semver.org).

For the full commit history, see [GitHub Releases](https://github.com/sanmoh-hombal/provider-checkly/releases).

## v0.1.0 (Pre-release)

Initial release of provider-checkly.

### Added

- 23 managed resources across 5 API groups: checks, alerts, infra, status pages, and triggers.
- Both cluster-scoped and namespace-scoped variants for all resources.
- Generated with Upjet from [`checkly/checkly`](https://github.com/checkly/terraform-provider-checkly) Terraform provider v1.22.0.
- Crossplane v2.x conformance — 56/56 tests passing.
- E2E test suite with kuttl.
- 5 example scenarios: ecommerce, multi-region, status page, Playwright, and GitOps setup.
