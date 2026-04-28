# Contributing

## Requirements

- [Go](https://go.dev/dl/) 1.26+
- [Task](https://taskfile.dev/installation/)
- [staticcheck](https://staticcheck.dev): `go install honnef.co/go/tools/cmd/staticcheck@latest`
- [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck): `go install golang.org/x/vuln/cmd/govulncheck@latest`

## Development

```sh
task build   # build binary to bin/
task check   # run fmt, vet, lint, vuln, and test
```

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
