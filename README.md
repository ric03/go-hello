# Go Playground

- tinkering with `go` ;)
- Querying the kubernetes API
- Executing git commands

# Useful commands

Run file `go run cmd/hello.go`

Help `go help`

# Dependencies

Dependencies are listed in `go.mod`

Add a package 

    go get <package>

To automatically update dependencies use

    go mod tidy

List available dependency updates (this also list the indirect dependencies)

    go list -m -u all

To update the dependencies run

    go get -u all
