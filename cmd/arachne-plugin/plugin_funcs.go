package main

import (
	"fmt"

	wapc "github.com/wapc/wapc-guest-tinygo"
)

/*
-------------------------
Functions that you want to expose for your plugin should go here. You can of course structure this like any Go program, this structure is just a suggestion.
-------------------------
*/

// hello will callback the host and return the payload
func Hello(payload []byte) ([]byte, error) {

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

// goodbye will callback the host and return the payload
func Goodbye(payload []byte) ([]byte, error) {

	//This will print in the host processes stdout
	fmt.Print("goodbye called")

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
	msg := "Goodbye, " + string(nameBytes) + ", also: " + string(sseBytes)

	// Return the message in byte format
	return []byte(msg), nil
}
