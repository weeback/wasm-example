//go:build js && wasm
// +build js,wasm

package main

import "syscall/js"

func main() {
	var (
		q = make(chan struct{})
	)
	println("WASM Go Initialized")
	//
	js.Global().Set("add", js.FuncOf(addFunc))
	//
	<-q
}
