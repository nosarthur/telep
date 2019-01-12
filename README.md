# telep

Run shell command(s) remotely.

## prerequisites

- protobuf
  - install pre-compiled binaries of [protobuf](https://github.com/protocolbuffers/protobuf/releases)
  - `go get -u github.com/golang/protobuf/protoc-gen-go`
- `go get -u google.golang.org/grpc`
- `go get -u github.com/spf13/cobra/cobra`

## TODO

- authentication
- wrap each command inside a `struct` for querying and control
