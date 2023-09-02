# Change this to your own Go module
GO_MODULE := hello-protobuf


.PHONY: tidy
tidy:
	go mod tidy

SHELL := $(if $(filter $(OS),Windows_NT),cmd.exe,/bin/sh)

.PHONY: clean
clean:
	$(if $(filter $(OS),Windows_NT), \
		if exist "protogen" rmdir /s /q protogen, \
		rm -rf ./protogen)


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \


.PHONY: build
build: clean protoc tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run


.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
