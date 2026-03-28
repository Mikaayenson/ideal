# Ideal

Opinionated **Go application starter**: a working CLI you can delete or extend, with layout and tooling that scale from prototype to production.

**What you get**

- **Layout**: [`cmd/`](https://github.com/golang-standards/project-layout) for binaries, [`internal/`](https://go.dev/doc/go1.4#internalpackages) for app code (not importable by other modules).
- **Config**: [12-factor](https://12factor.net) settings via environment (`IDEAL_*`), parsed with [`caarlos0/env`](https://github.com/caarlos0/env).
- **Logging**: [`log/slog`](https://pkg.go.dev/log/slog) with text or JSON (level names match [slog levels](https://pkg.go.dev/log/slog#Level)).
- **Quality**: `go vet`, tests with race + shuffle, [`golangci-lint` v2](https://golangci-lint.run/), and [`govulncheck`](https://go.dev/doc/security/vuln/) in CI.
- **Maintenance**: [Dependabot](.github/dependabot.yml) for Go modules and GitHub Actions.

## Layout

| Path | Role |
|------|------|
| `cmd/ideal/` | Thin `main`: flags, config, wiring, signal-aware context |
| `internal/config/` | Typed env config |
| `internal/logging/` | slog setup (text or JSON) |
| `internal/greetings/` | Example domain package + tests |
| `internal/version/` | Build metadata (`-version`, logs); falls back to [debug.ReadBuildInfo](https://pkg.go.dev/runtime/debug#ReadBuildInfo) |

**Toolchain:** optional [mise](https://mise.jdx.dev/) config in [`mise.toml`](mise.toml) pins a Go release for contributors (`mise install`).

## Requirements

- [Go](https://go.dev/dl/) **1.24+** (see `go.mod`)

## Quick start

```bash
git clone <your-repo-url> ideal
cd ideal
go run ./cmd/ideal
```

Override config with environment variables (see [`.env.example`](.env.example)):

```bash
export IDEAL_USERNAME=Ada
export IDEAL_LOG_LEVEL=debug
go run ./cmd/ideal
```

## First-time fork checklist

1. Set the **module path** in `go.mod` and fix `import` paths (see below).
2. Update **`LICENSE`** copyright for your project.
3. Rename the **`ideal`** binary: `cmd/ideal` → `cmd/<yourapp>`, update `Makefile` `BINARY`, Dockerfile paths, and CI if needed.

## Replace the module path

1. Edit the first line of `go.mod` (`module …`).
2. Replace imports of `github.com/stryker/ideal` across the tree.

**macOS (BSD sed)**

```bash
rg -l 'github.com/stryker/ideal' --glob '*.go' | xargs sed -i '' 's|github.com/stryker/ideal|github.com/you/repo|g'
```

**Linux (GNU sed)**

```bash
rg -l 'github.com/stryker/ideal' --glob '*.go' | xargs sed -i 's|github.com/stryker/ideal|github.com/you/repo|g'
```

## Commands

Run `make help` for a short summary of every target.

| Command | Description |
|---------|-------------|
| `go test ./...` | Run tests |
| `make vet` | `go vet ./...` |
| `make test` | Tests with race detector and shuffle |
| `make cover` | Coverage report → `coverage.html` |
| `make build` | Build `bin/ideal` with version from `git describe` |
| `make vuln` | Run `govulncheck` (pinned version, same as CI) |
| `make lint` | golangci-lint v2 (install separately; see below) |
| `make fmt` | `golangci-lint fmt` |
| `make check` | `vet` + `test` + `lint` + `build` |
| `make docker-build` | Image with `VERSION` build-arg from git |
| `make install-lint` | Optional: install pinned golangci-lint into `GOPATH/bin` |

### Installing golangci-lint locally

- **Homebrew**: `brew install golangci-lint` (ensure v2; `golangci-lint version`).
- **Make**: `make install-lint` installs the same **v2.11.4** pin used in [`.github/workflows/ci.yml`](.github/workflows/ci.yml).

### Live reload (optional)

```bash
go install github.com/air-verse/air@latest
air
```

## Continuous integration

[`.github/workflows/ci.yml`](.github/workflows/ci.yml) runs on pushes and PRs to `main`:

- `go vet`, release-mode `go build`, and `go test -race -shuffle`
- `govulncheck` on `./...`
- golangci-lint **v2.11.4**

## Docker

```bash
docker build --build-arg VERSION="$(git describe --tags --always --dirty 2>/dev/null || echo dev)" -t ideal:local .
docker run --rm -e IDEAL_USERNAME=Ada ideal:local
```

(`make docker-build` passes `VERSION`, `COMMIT`, and `BUILD_DATE` for you.)

## Extending this template (2026+)

Ideas that stay out of the minimal core but are common for real services—add what you need, delete what you do not.

| Area | Libraries / approaches |
|------|-------------------------|
| **CLI** | [spf13/cobra](https://github.com/spf13/cobra), [alecthomas/kong](https://github.com/alecthomas/kong) for subcommands and flags at scale |
| **HTTP** | Start with `net/http` + graceful `Shutdown`; then [go-chi/chi](https://github.com/go-chi/chi), [labstack/echo](https://github.com/labstack/echo), or [gofiber/fiber](https://github.com/gofiber/fiber) if you want a router/framework |
| **Config** | Keep env-first; add file overlays with [knadh/koanf](https://github.com/knadh/koanf). Local-only `.env`: [joho/godotenv](https://github.com/joho/godotenv) (do not rely on it in production) |
| **Data** | [jackc/pgx](https://github.com/jackc/pgx) + [sqlc](https://sqlc.dev/), or [ent](https://entgo.io/) / [gorm](https://gorm.io/) if you prefer an ORM |
| **Observability** | [OpenTelemetry Go](https://go.opentelemetry.io/otel/) for traces and metrics; expose `pprof` and `expvar` behind auth in production |
| **API contracts** | [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen), [buf](https://buf.build/) + Connect/gRPC |
| **Resilience** | [sony/gobreaker](https://github.com/sony/gobreaker), `golang.org/x/sync/errgroup`, context deadlines |
| **Testing** | [stretchr/testify](https://github.com/stretchr/testify) (assertions), [testcontainers-go](https://golang.testcontainers.org/) (integration), [go.uber.org/mock](https://github.com/uber-go/mock) or `go generate` with `mockgen` |
| **Releases** | [GoReleaser](https://goreleaser.com/) for archives, images, SBOMs, and changelog |
| **Policy / security** | Keep `govulncheck` and golangci-lint; consider [OpenSSF Scorecard](https://scorecard.dev/) and dependency review on PRs |

## References

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Organizing a Go module](https://go.dev/doc/modules/layout)
- [Effective Go](https://go.dev/doc/effective_go)
- [Vulnerability management](https://go.dev/doc/security/vuln/)

## License

See [LICENSE](LICENSE).
