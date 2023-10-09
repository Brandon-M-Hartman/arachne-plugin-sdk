ifeq ($(OS),Windows_NT)
	SHELL := cmd
	OP = build_on_windows
else
	SHELL := /bin/bash
	OP = build_on_linux
endif

build-wasm: build-demo build-plugin

build-demo:
	@$(MAKE) -C cmd/demo $(OP)

build-plugin:
	@$(MAKE) -C cmd/arachne_plugin build
