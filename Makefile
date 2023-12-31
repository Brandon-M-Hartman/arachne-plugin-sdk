ifeq ($(OS),Windows_NT)
	SHELL := cmd
	OP = build_on_windows
else
	SHELL := /bin/bash
	OP = build_on_linux
endif

build: build-demo build-plugin

update:
	@$(MAKE) -C cmd/demo update-mods $(OP)
	@$(MAKE) -C cmd/arachne-plugin update-mods build	

build-demo:
	@$(MAKE) -C cmd/demo $(OP)

build-plugin:
	@$(MAKE) -C cmd/arachne-plugin build