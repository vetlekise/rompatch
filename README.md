# [Project Name]

A Go [package/tool/library] for [describe the primary action, e.g., interacting with the X API, parsing Y data].

A more detailed description of the project, its purpose, and what problems it solves.

## Setup

1. Initialise the module:
   ```sh
   go mod init github.com/<username>/<project-name>
   ```

2. Set the binary name in `Taskfile.yaml`:
   ```yaml
   vars:
     BINARY: <project-name>
   ```

3. Update `main.go` — replace the package comment with a description of your tool.

4. Fill in this `README.md`:
   - Replace `[Project Name]` with your project name.
   - Update installation URLs (`<username>/<project-name>`).
   - Write your own **Usage** section.

5. Delete this **Setup** section.

## Features

- **Automated CI**: Formatting, vetting, linting, vulnerability scanning, and tests on every pull request via GitHub Actions.
- **Static Analysis**: Enforced with [staticcheck](https://staticcheck.dev).
- **Vulnerability Scanning**: Via [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck).
- **Automated Releases**: Cross-platform binaries built and published on tag push using [GoReleaser](https://goreleaser.com).
- **Conventional Commits**: PR titles validated with [action-semantic-pull-request](https://github.com/amannn/action-semantic-pull-request).
- **Task Runner**: Common development tasks defined in [Taskfile](https://taskfile.dev).

## Installation

**Using Go:**

```sh
go install github.com/<username>/<project-name>@latest
```

**Pre-built binaries:**

Download the latest release for your platform from the [GitHub Releases](https://github.com/<username>/<project-name>/releases) page.

## Usage

<!-- Describe how to use the tool/library here. -->