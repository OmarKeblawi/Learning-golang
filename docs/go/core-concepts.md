# Go Core Concepts

## Module

- A module is a project-level boundary in Go.
- It is defined by `go.mod`.
- In this repo, the module is `go_tutorials`.
- A module can contain many packages.

## Package

- A package is the unit Go compiles and imports.
- In normal Go code, one directory contains one package for a build.
- All `.go` files in the same directory with the same `package` declaration are compiled together as one package.
- A package is how Go groups related code.

## Package vs folder

- A folder usually maps to one package.
- The folder name and package name do not have to match, but they usually do.
- One directory cannot normally contain two different packages in the same build.

## `package main`

- `package main` marks code as an executable program package.
- If a package is named `main` and contains `func main()`, Go can build it as a runnable binary.
- Other files in the same directory that also say `package main` become part of the same executable package.
- Those files can define helper functions, types, and variables that `main.go` can call.
- Only one `func main()` should exist in package `main`.

## Reusable library

- A reusable library is code meant to be imported and used by other packages.
- It is usually not run directly.
- Standard library packages like `fmt` are examples of reusable packages.

## Compilation model

- Go compiles a package as one combined unit.
- This is similar to how several C source files can contribute to one program.
- It is not one binary per file.
- The package is the boundary for compilation and reuse.

## Import path

- A package is what other Go code imports.
- The import path is usually based on the module name plus the package location.
- Packages live inside modules.

## Quick mental model

- module = project and dependency boundary
- package = directory of Go files compiled together
- file = one source file inside a package
- `package main` + `func main()` = executable entry point

## Notes from this repo

- `go.mod` defines the module.
- `cmd/tut1/main.go` belongs to `package main`.
- `cmd/tut1` can hold more Go files with `package main` and they will compile into the same program.
