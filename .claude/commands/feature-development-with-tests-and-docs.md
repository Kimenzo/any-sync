---
name: feature-development-with-tests-and-docs
description: Workflow command scaffold for feature-development-with-tests-and-docs in any-sync.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /feature-development-with-tests-and-docs

Use this workflow when working on **feature-development-with-tests-and-docs** in `any-sync`.

## Goal

Add a new feature or major improvement, including implementation, tests, and documentation.

## Common Files

- `**/*.go`
- `**/*_test.go`
- `docs/**/*.md`
- `README.md`
- `go.mod`
- `go.sum`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Implement the feature in multiple new or existing .go files.
- Add or update test files (e.g., *_test.go) to cover the new functionality.
- Add or update documentation files (e.g., docs/plans/*.md, README.md).
- Update go.mod and go.sum if new dependencies are required.
- Commit all related implementation, test, and documentation files.

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.