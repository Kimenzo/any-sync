---
name: proto-schema-and-generated-code-update
description: Workflow command scaffold for proto-schema-and-generated-code-update in any-sync.
allowed_tools: ["Bash", "Read", "Write", "Grep", "Glob"]
---

# /proto-schema-and-generated-code-update

Use this workflow when working on **proto-schema-and-generated-code-update** in `any-sync`.

## Goal

Update protobuf schema files and regenerate Go code for protocol buffers across multiple modules.

## Common Files

- `**/*.proto`
- `**/*.pb.go`
- `**/*_vtproto.pb.go`

## Suggested Sequence

1. Understand the current state and failure mode before editing.
2. Make the smallest coherent change that satisfies the workflow goal.
3. Run the most relevant verification for touched files.
4. Summarize what changed and what still needs review.

## Typical Commit Signals

- Edit one or more .proto files (e.g., add/remove fields, messages, or services).
- Regenerate Go code using protoc or buf (resulting in changes to .pb.go and sometimes .vtproto.pb.go files).
- Update related model or builder files if necessary.
- Commit all updated .proto and generated .pb.go files.

## Notes

- Treat this as a scaffold, not a hard-coded script.
- Update the command if the workflow evolves materially.