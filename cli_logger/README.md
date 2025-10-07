# Logger

This Logger CLI is a library exporting Logger level types and Logger functions.
Some functions are kept private to avoid exposing too much. In order to keep the
function signatures intuitive, the CLI has domain-specific types.

```go
type Level byte

const (
 LevelDebug Level = iota
 LevelInfo
 LevelWarn
 LevelError
 LevelFatal
)
```

Within the `New()` method, the CLI implements the required parameters while
setting sensible defaults through the
[functional options pattern](./pocketlog/options.go). The tests rely on
closed-box testing to validate behavior rather than internal details.

```zsh
go test ./pocketlog
```

## Go Doc

```zsh
# print the documentation for the package found in doc.go
go doc pocketlog

# print the documentation for the Logger type and its functions
go doc pocketlog.Logger

# print the documentation for a specific function
go doc pocketlog.Logger.Debugf
```
