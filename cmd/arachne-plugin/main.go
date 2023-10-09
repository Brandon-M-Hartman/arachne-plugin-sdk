package main

import (
	"fmt"

	wapc "github.com/wapc/wapc-guest-tinygo"
)

func main() {
	// Register echo and fail functions
	wapc.RegisterFunctions(wapc.Functions{
		"hello": hello,
	})
}

// hello will callback the host and return the payload
func hello(payload []byte) ([]byte, error) {

	//This will print in the host processes stdout
	fmt.Print("hello called")

	// Make a host call to capitalize the name.
	nameBytes, err := wapc.HostCall("", "example", "capitalize", payload)
	if err != nil {
		return nil, err
	}

	// Make a second host call to try out the other case
	sseBytes, err2 := wapc.HostCall("", "example", "say something else", payload)
	if err2 != nil {
		return nil, err2
	}

	// Format the message.
	msg := "Hello there, " + string(nameBytes) + ", also: " + string(sseBytes)

	// Return the message in byte format
	return []byte(msg), nil
}
