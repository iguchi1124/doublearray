GO111MODULE_ON := GO111MODULE=on
BUILD_FILE := bin/doublearray

.PHONY: default
default: setup clean build

.PHONY: setup
setup:
	@mkdir -p bin

.PHONY: build
build:
	@$(GO111MODULE_ON) go build -o $(BUILD_FILE)
	@if test -e $(BUILD_FILE);\
	then\
		echo "Success";\
	else\
		echo "Failure";\
	fi

.PHONY: test
test:
	@$(GO111MODULE_ON) go test

.PHONY: benchmark
benchmark:
	@$(GO111MODULE_ON) go test -bench . -benchmem

.PHONY: clean
clean:
	@rm -f $(BUILD_FILE)

.PHONY: lint
lint:
	@golint -set_exit_status
