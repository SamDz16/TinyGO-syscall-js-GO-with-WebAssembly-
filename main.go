package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func add() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 2 {
			return args[0].Int() + args[1].Int()
		}
		return -1
	})
}

func capitalized() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 1 {
			return strings.ToUpper(args[0].String())
		}
		return -1
	})
}


func main() {
	fmt.Printf("Hello Web Assembly from Go!\n")

	js.Global().Set("add", add())
	js.Global().Set("capitalized", capitalized())
	
	
	<-make(chan struct{}, 0)
}

