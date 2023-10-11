ifeq ($(OS),Windows_NT)
	SHELL := cmd
	OP = build_on_windows

else
	SHELL := /bin/bash
	OP = build_on_linux
	
endif

build: build-binary

build-binary:
	@echo "----------"
	@echo "Building Plugin with TinyGo"
	@echo "----------"
	tinygo build -o ../../bin/wasm/arachne-plugin.wasm -scheduler=none --no-debug -target=wasi ./

