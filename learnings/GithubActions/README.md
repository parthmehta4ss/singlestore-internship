# GitHub Actions — quick notes

Docs: https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax

YAML primer: http://learnxinyminutes.com/yaml/

## Basics
- CI (Continuous Integration): run on pushes/PRs to build and test changes.
- CD (Continuous Deployment): after CI passes, build artifacts (binaries/images) and deploy.

Minimal sample workflow (`.github/workflows/ci.yml`):

```yaml
name: CI
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run tests
        run: echo "Running tests..."
```

What this does: on push or PR, GitHub runs the `build` job on an Ubuntu runner, checks out the code, and runs the test step.

Starter examples: https://github.com/actions/starter-workflows

Keep it simple: start with one workflow for CI, expand jobs and steps as needed.

## Major topics to explore
- **Triggers:** `push`, `pull_request`, `workflow_dispatch`, `schedule` (cron).
- **Workflows, jobs, steps:** structure, `needs` (job dependencies), and `if` conditionals.
- **Runners:** GitHub-hosted vs self-hosted runners and runner labels.
- **Actions:** marketplace actions, custom actions, composite actions, and reusable workflows.
- **Matrix & strategy:** run jobs across multiple OSes, languages, or versions.
- **Secrets & env:** `secrets`, environment variables, protected environments and required reviewers.
- **Caching & artifacts:** speed up builds with cache; upload/download artifacts for jobs.
- **Service containers:** use databases or services during job runs (containers/containers.services).
- **Outputs & artifacts:** pass outputs between jobs and persist build results.
- **Permissions & security:** token permissions, OIDC, secret scanning, and least-privilege workflows.
- **Concurrency & workflow control:** `concurrency`, `cancel-in-progress`, and rate-limiting considerations.
- **Logging & debugging:** workflow logs, step outputs, and annotations for failures.
- **Marketplace & community:** find vetted actions and prefer well-maintained ones.
SELECT * FROM users WHERE id IN (1,2,3)