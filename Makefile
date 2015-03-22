# default environment variables
PORT?=3000
HOST?=127.0.0.1
MONGO?=localhost
MONGODB?=toot-collect

.PHONY: build run deps

# `make build` compiles the static binary executable
build:
	@GOPATH="$(PWD)/ext" go build

./clerk:
	@$(MAKE) build

# `make run` executes the compiled binary
run: ./toot-collect
	@HOST=$(HOST) PORT=$(PORT) MONGO=$(MONGO) ./toot-collect

# `make deps` rebuilds the .deps/ directory by resolving all dependencies
deps:
	@rm -rf "$(PWD)/ext/*"
	@go list -f '{{range .Imports}}{{.}}|{{end}}{{range .TestImports}}{{.}}|{{end}}{{range .XTestImports}}{{.}}|{{end}}' ./... | \
		awk -F '|' '{ for (i = 1; i <= NF; i++) if ($$i != "") print $$i }' | sort | uniq | grep --invert-match '^\.' | \
		GOPATH="$(PWD)/.deps" xargs go get -v -d -u

# `make lint` formats all Go files in the project
lint:
	@go fmt ./...

# `make go` runs and builds the project
go:
	@HOST=$(HOST) PORT=$(PORT) MONGO=$(MONGO) MONGODB=$(MONGODB) GOPATH="$(PWD)/ext" go run toot-collect.go
