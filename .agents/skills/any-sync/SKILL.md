```markdown
# any-sync Development Patterns

> Auto-generated skill from repository analysis

## Overview

This skill provides a comprehensive guide to the development patterns, coding conventions, and common workflows used in the `any-sync` Go codebase. It covers file organization, code style, commit practices, and step-by-step instructions for frequent development tasks such as updating protocol buffers, adding features, managing dependencies, and maintaining test compatibility.

## Coding Conventions

- **Language:** Go
- **Framework:** Go standard library (no major external frameworks detected)
- **File Naming:** Use `snake_case` for file names.
  - Example: `sync_manager.go`, `user_acl.go`
- **Import Style:** Use relative imports within the module.
  - Example:
    ```go
    import (
        "context"
        "github.com/yourorg/any-sync/internal/model"
    )
    ```
- **Export Style:** Use named exports for functions, types, and constants.
  - Example:
    ```go
    // Exported function
    func NewSyncManager() *SyncManager { ... }
    ```
- **Commit Messages:** Freeform, sometimes prefixed (e.g., `build:`, `acl:`), average length ~57 characters.

## Workflows

### Proto Schema and Generated Code Update
**Trigger:** When you change a `.proto` file to add, remove, or update protocol messages or services.  
**Command:** `/update-proto`

1. Edit one or more `.proto` files (e.g., add/remove fields, messages, or services).
2. Regenerate Go code using `protoc` or `buf`:
   ```sh
   protoc --go_out=. --go-grpc_out=. path/to/file.proto
   ```
3. Update related model or builder files if necessary.
4. Commit all updated `.proto`, `.pb.go`, and `*_vtproto.pb.go` files.

**Files Involved:**  
- `**/*.proto`  
- `**/*.pb.go`  
- `**/*_vtproto.pb.go`  

### Feature Development with Tests and Docs
**Trigger:** When implementing a new major feature or protocol (e.g., WebTransport, AddSeq).  
**Command:** `/new-feature`

1. Implement the feature in new or existing `.go` files.
2. Add or update test files (`*_test.go`) to cover the new functionality.
   ```go
   func TestNewFeature(t *testing.T) {
       // test logic here
   }
   ```
3. Add or update documentation (`docs/plans/*.md`, `README.md`).
4. Update `go.mod` and `go.sum` if new dependencies are required.
5. Commit all related implementation, test, and documentation files.

**Files Involved:**  
- `**/*.go`  
- `**/*_test.go`  
- `docs/**/*.md`  
- `README.md`  
- `go.mod`, `go.sum`  

### Dependency Bump via Dependabot
**Trigger:** When a new version of a dependency is released and `go.mod`/`go.sum` are updated (often via Dependabot).  
**Command:** `/bump-dependency`

1. Update `go.mod` and `go.sum` to reflect new dependency versions.
2. Commit `go.mod` and `go.sum` with a message referencing the dependency and new version.

**Files Involved:**  
- `go.mod`  
- `go.sum`  

### Test Fix or Compatibility Update
**Trigger:** When a new Go version is released or a test fails due to environment changes.  
**Command:** `/fix-test`

1. Edit or add `*_test.go` files to fix compatibility or test logic.
2. Update `Makefile` or CI scripts if necessary.
3. Commit test and/or `Makefile` changes.

**Files Involved:**  
- `**/*_test.go`  
- `Makefile`  

## Testing Patterns

- **Test Files:** Named with the `_test.go` suffix (e.g., `sync_manager_test.go`).
- **Framework:** Standard Go testing (`testing` package).
- **Example Test:**
  ```go
  import "testing"

  func TestSyncManager_Run(t *testing.T) {
      // Arrange
      mgr := NewSyncManager()
      // Act
      err := mgr.Run()
      // Assert
      if err != nil {
          t.Fatalf("expected no error, got %v", err)
      }
  }
  ```
- **Test Location:** Tests are placed alongside implementation files.

## Commands

| Command           | Purpose                                                         |
|-------------------|-----------------------------------------------------------------|
| /update-proto     | Update protobuf schema and regenerate Go code                   |
| /new-feature      | Start a new feature with implementation, tests, and documentation|
| /bump-dependency  | Update Go module dependencies                                   |
| /fix-test         | Fix or update tests for compatibility or failures               |
```