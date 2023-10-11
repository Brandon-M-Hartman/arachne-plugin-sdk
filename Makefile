ifeq ($(OS),Windows_NT)
	SHELL := cmd
	OP = build_on_windows

else
	SHELL := /bin/bash
	OP = build_on_linux

endif

build-wasm: update-mods build-demo build-plugin

update-mods:
	@echo "Updating go.mod files"
	@go mod tidy
	go get -u ./...
	cd cmd/demo
	@go mod tidy
	go get -u ./...
	cd cmd/arachne-plugin
	@go mod tidy
	go get -u ./...

build-demo:
	@$(MAKE) -C cmd/demo $(OP)

build-plugin:
	@$(MAKE) -C cmd/arachne-plugin build