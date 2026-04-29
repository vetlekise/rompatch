# Contributing

## Requirements

- [Go](https://go.dev/dl/) 1.26+
- [Task](https://taskfile.dev/installation/)
- [staticcheck](https://staticcheck.dev): `go install honnef.co/go/tools/cmd/staticcheck@latest`
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck): `go install golang.org/x/vuln/cmd/govulncheck@latest`

> [!NOTE]
> Go installs binaries to `$GOPATH/bin` (default: `$HOME/go/bin`). Make sure this is in your `$PATH`:
> ```sh
> # macOS example:
> echo 'export PATH="$PATH:$HOME/go/bin"' >> ~/.zshrc && source ~/.zshrc
> ```

## Development

This project uses [Task](https://taskfile.dev) to automate common development workflows locally. The main ones are:

| Command | Description |
|---|---|
| `task check` | Run all checks: fmt, vet, lint, vuln, and test |
| `task build` | Run all checks and build binary to `bin/` |

For the full list, see [Taskfile.yaml](Taskfile.yaml).

## Pull Requests

This project uses **squash merges**, so the PR title becomes the single commit message on `main` and is used to generate the changelog on release.

PR titles must follow the [Conventional Commits](https://www.conventionalcommits.org/) format:

```
<type>: <description>

feat(ui): Add `Button` component
^    ^    ^
|    |    |__ Subject
|    |_______ Scope
|____________ Type
```

Common types: `feat`, `fix`, `docs`, `chore`, `refactor`, `test`.
