//go:build wasm
// +build wasm

package main

import (
	"syscall/js"

	"github.com/sprkweb/finaplan-web/src/lib/go/internal/jsfinaplan"
)

func main() {
	js.Global().Set("InitFinaplan", js.FuncOf(jsfinaplan.InitPlan))

	<-make(chan struct{})
}
