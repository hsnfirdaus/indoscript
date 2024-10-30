//go:build !wasm
// +build !wasm

package utils

import "os"

func CetakError(pesan string) {
	println(pesan)
	os.Exit(1)
}
