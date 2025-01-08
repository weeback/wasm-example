//go:build js && wasm
// +build js,wasm

package main

import "syscall/js"

func addFunc(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		return js.Error{
			Value: js.ValueOf("parameter invalid"),
		}
	}
	return args[0].Int() + args[1].Int()
}
