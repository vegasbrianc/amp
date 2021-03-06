# Contributing

## Prerequisites

* [Docker 1.12.x](https://www.docker.com/products/docker)
* [Go](https://golang.org/dl/)
* [Glide](https://glide.sh)
* [Golint](https://github.com/golang/lint/golint)

## Directories

* `api/rpc` - contains the directories for specifying our rpc services and messages. They should be segregated by service and only contain *.proto files for definitions and *.pb.go files that are generated by the rpc rule in the Makefile (`make rpc`). Running `make install` or `make build` will automatically invoke the rpc rule.
* `api/server` - contains the implementation files for the api/rpc/* service interfaces, and also contains `server.go`, which registers each service with a server on port 50501.
* `api/cmd` - contains directories corresponding to the two generated executable binaries: `amp` (the CLI) and `amplifier` (the server daemon).
* `data` - backend datastores for configuration, state, stats, and logs.
* `vendor` - along with `glide.yaml` and `glide.yaml`, used for dependency management.

## Makefile

* `make` - print version/build info, run a check on your code, then build `amp` and `amplifier` in the project root.
* `make install` - (re)build the project and install the executable binaries (`amp` and `amplifier`) in `$GOPATH/bin`.
* `make check` - runs formatting and lint checks on the source. Make sure to run before submitting a PR.
* `make fmt` - format source files according to go conventions.
* `make clean` - cleans up auto-generated code and deletes the `amp` and `amplifier` executables from `$GOPATH/src`.
* `make test` - run automated tests.
* `make proto` - regenerate source files (*.pb.go) based on protocol buffer (*.proto) definition files using `protoc`.
* `make install-deps` - will reinstall all required packages under the `vendor` directory (same as `glide install`)
* `make update-deps` - will update all package dependencies (same as `glide update`).

## Package Management

The project uses [Glide](https://glide.sh/) for vendoring, which is an approach to locking down packages that you have tested for a specific
version. Glide isn't the only tool that supports vendoring, but it is one of the more popular ones. The Go community generally commits all the
vendor artifacts (the `vendor` directory and lock file (`glide.lock`), along with the dependency specification (`glide.yaml`) file).

When you want to add new dependencies to the project, run `glide get <package>`.

When you want to reinstall dependencies, run `glide install` (or `make install-deps`).

When you want to update to the latest versions, run `glide update` (or `make update-deps`). Run `make test` to ensure tests continue to pass after updates
before submitting a PR.

