//go:build wasm
// +build wasm

package utils

import (
	"syscall/js"
)

func CetakError(pesan string) {
	global := js.Global()
	global.Call("alert", pesan)
}
