## Arachne Plugin SDK

This repository contains the Software Development Kit (SDK) for creating plugins for the Arachne configuration automation platform. The SDK provides a set of tools and libraries that developers can use to develop their own plugins for Arachne.

### Getting Started

The `cmd/arachne-plugin` directory contains a sample plugin written in Go, which serves as a practical example.

In this example, WebAssembly (WASM) and the WebAssembly System Interface (WASI) are used to create a portable binary that can be loaded and executed by the Arachne platform. The Go code is compiled to a WASM+WASI binary using TinyGo, a Go compiler designed for small places that supports compiling Go programs to WASM. While Go, as of 1.21rc2, supports compiling to WASI, waPC-Go and wasmzero/wasmtime/wasmedge don't support GO-compiled WASI modules terribly well, so we are still using TinyGo for now for plugins.

The `cmd/arachne-plugin/main.go` file is the entry point of the plugin. It registers the functions that the plugin exposes and implements the required functions for the Arachne platform.

The `cmd/arachne-plugin/custom_plugin_funcs.go` file contains the functions that the plugin exposes. These functions can be invoked by the Arachne platform and can make host calls.

The `cmd/arachne-plugin/configuration.go` file contains the configuration constants for the plugin, such as the plugin name, developer identity, plugin URL, plugin version, and description.

The `pkg/shared/arachne-plugin/scaffold/scaffold.go` file contains the structure of the plugin system and the host function to handle calls from the wasm module.

The `cmd/demo/main.go` and `cmd/demo/configuration.go` files are used for testing the plugin. They load the plugin and handle the logging.

The Makefile in the root directory contains commands for building and running the plugin. The build target uses TinyGo to compile the Go code to a WASM binary. 

Here's a brief overview of how to use this example to create your own plugins:

1. Write your plugin logic in Go, following the structure of the example plugin in `cmd/arachne-plugin`. Define your own host functions and invoke functions in the WASM module as needed.

2. Compile your Go code to a WASM binary using TinyGo. You can use the build target in the Makefile as a reference.

3. Test your plugin by executing the WASM binary using the demo application in `cmd/demo`. You can compile `cmd/demo/main.go`, or you can just use "go run"

Remember to comply with the licenses of any third-party software you use in your plugin, as described in the "Third Party Licenses" section of the README.

### Contributing

We welcome contributions to the Arachne Plugin SDK! If you have any bug reports, feature requests, or pull requests, please submit them to the [Arachne Plugin SDK GitHub repository](https://github.com/getarachne/arachne-plugin-sdk).

### License

This project is licensed under the Mozilla Public License Version 2.0 - see the [LICENSE](LICENSE) file for details.

### Third Party Licenses

This project uses the following open source software:

- [wapc-go](https://github.com/wapc/wapc-go): Licensed under the [Apache License 2.0](https://github.com/wapc/wapc-go/blob/master/LICENSE.txt)
