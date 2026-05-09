# Go Workspace Layout

When you install Go on WSL or Linux, you may see a `go` directory in your home folder with `bin`, `pkg`, and `src`.

## `bin`

- Holds compiled executables installed by Go.
- Commonly used for command-line tools.
- If `$(go env GOPATH)/bin` is on your `PATH`, you can run those tools directly.

## `pkg`

- Holds compiled package cache and build artifacts.
- Go can reuse this data to speed up builds.
- You usually do not edit this folder by hand.

## `src`

- Old GOPATH-style source code location.
- In older Go workflows, your packages lived under this folder.
- In modern Go, modules are the normal workflow, so this folder is much less important.

## Modern Go behavior

- With modules, your project can live anywhere on disk as long as it has a `go.mod` file.
- You do not have to place module-based projects under `src`.

## Helpful distinction

- The `go` folder in your home directory is usually your workspace or `GOPATH`.
- The Go installation itself is usually elsewhere, often `/usr/local/go`.

## Simple mental model

- `bin` = installed commands
- `pkg` = compiled cache/artifacts
- `src` = source code location from the older workflow
