### Variables
PKG_NAME = fried-gophers

## Source config
SRC_FILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
SRC_PKGS = $(shell go list ./...)

## Testing config
TEST_TIMEOUT = 20s
COVER_OUT = coverage.txt


### Commands (targets):
## Prevent targeting filenames...
.PHONY: default all all-bench build \
				clean clean-all clean-lib \
				test test-v test-race bench \
				run check fmt 

## Default target when no arguments are given to make.
default: fmt check test
all: fmt check test-race

## Like all, but includes benchmarks.
all-bench: fmt check test-race bench


## Build the program.
build:
	@printf "Building..."; go build; echo "OK!"

clean-all: clean clean-lib

## Clean executables created by this package.
clean:
	@printf "Cleaning executables"
	@if [ -f "$(PKG_NAME)" ]; then rm $(PKG_NAME); fi;
	@XPATH="$$GOPATH/bin/$(PKG_NAME)"; if [ -f "$$XPATH" ]; then rm "$$XPATH"; fi
	@echo "...OK!"

## Clean libraries created by this package.
# clean-lib:
# 	@printf "Cleaning libraries"
# 	@for PKG in $(SRC_PKGS); do LIBPATH="$$GOPATH/pkg/$(PKG)"; \
# 		if [ -d $$LIBPATH ]; then rm -rf $$LIBPATH; fi; done
# 	@echo "...OK!"

## Build and run the program.
run: build
	@./$(PKG_NAME)

## Checks for formatting, linting, and suspicious code.
check:
## Check formatting...
	@printf "Check format:"
	@GOFMT_OUT="$(shell gofmt -l $(SRC_FILES))"; if [ -n "$$GOFMT_OUT" ]; then \
		printf "\n> [WARN] Fix formatting issues in the following files with "; \
		printf "'make fmt':\n$$GOFMT_OUT\n\n"; else echo " ...OK!"; fi
## Lint files...
	@printf "Check lint:"
	@GOLINT_OUT="$(shell for PKG in "$(SRC_PKGS)"; do golint $$PKG; done)"; \
		if [ -n "$$GOLINT_OUT" ]; then printf "\n" && \
		for PKG in "$$GOLINT_OUT"; do echo "> $$PKG"; done; printf "\n"; \
		else echo " ...OK!"; fi
## Check suspicious code...
	@printf "Check vet:"
	@GOVET_OUT="$(shell go vet 2>&1)"; if [ -n "$$GOVET_OUT" ]; \
		then printf "\n> [WARN] Fix suspicious code from 'go vet':\n"; \
		printf "$$GOVET_OUT\n\n"; else echo " ...OK!"; fi

## Reformats code according to 'gofmt'.
fmt:
	@echo Formatting:
	@GOFMT_OUT=$(shell gofmt -l -s -w $(SRC_FILES)); if [ -n "$$GOFMT_OUT" ]; \
		then for FILE in "$$GOFMT_OUT"; do printf "> $$FILE\n"; done; \
		else echo "> ...all files formmatted correctly!"; fi

## Testing commands:
GOTEST = go test --coverprofile=$(COVER_OUT) --timeout=$(TEST_TIMEOUT)
test:
	@echo "Testing..."
	@$(GOTEST) ./...
test-v:
	@echo "Testing (verbose)..."
	@$(GOTEST) -v ./...
test-race:
	@echo "Testing (race)..."
	@$(GOTEST) --race ./...
bench:
	@echo "Benchmarking..."
	@go test --run=a^ --bench=. --benchmem
