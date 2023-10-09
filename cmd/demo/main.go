package main

import (
	"context"
	"fmt"
	"os"

	"github.com/wapc/wapc-go/engines/wazero"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/wapc/wapc-go"
)

// Main function of the program
func main() {

	// Check if the program has the correct number of arguments
	if len(os.Args) < 2 {
		fmt.Println("usage: hello <name>")
		return
	}

	// Get the name argument from the command line
	name := os.Args[1]

	// Create a new context
	ctx := context.Background()

	// Read the wasm plugin file
	guest, err := os.ReadFile("./plugin/arachne_plugin.wasm")

	// Check for errors in reading the file
	if err != nil {
		panic(err)
	}

	// Initialize the wasm engine
	engine := wazero.Engine()

	// Create a new module with the wasm engine, host function, wasm file and module configuration
	module, err := engine.New(ctx, host, guest, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})

	// Check for errors in creating the module
	if err != nil {
		panic(err)
	}

	// Ensure the module is closed at the end of the function
	defer module.Close(ctx)

	// Instantiate the module
	instance, err := module.Instantiate(ctx)

	// Check for errors in instantiating the module
	if err != nil {
		panic(err)
	}

	// Ensure the instance is closed at the end of the function
	defer instance.Close(ctx)

	// Invoke the "hello" function in the wasm module with the name as argument
	result, err := instance.Invoke(ctx, "hello", []byte(name))

	// Check for errors in invoking the function
	if err != nil {
		panic(err)
	}

	// Print the result of the function invocation
	fmt.Println(string(result))
}

// Host function to handle calls from the wasm module
func host(_ context.Context, binding, namespace, operation string, payload []byte) ([]byte, error) {
	// Route the payload to any custom functionality accordingly.
	// You can even route to other waPC plugins!
	switch namespace {
	case "example":

		switch operation {

		case "capitalize":
			// If the operation is "capitalize", capitalize the payload and return it
			name := string(payload)
			name = cases.Title(language.English).String(name)
			return []byte(name), nil

		case "say something else":
			// If the operation is "say something else", return a predefined string
			reply := "something else"
			return []byte(reply), nil
		}
	}
	// If no operation matches, return a default string
	return []byte("default"), nil
}
