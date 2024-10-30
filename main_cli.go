//go:build !wasm
// +build !wasm

package main

import (
	"errors"
	"fmt"
	"os"
)

func mulaiInterpreter() {
	args := os.Args[1:]
	if len(args) < 1 {
		println("Anda harus memasukan path file sebagai argumen!")
		os.Exit(1)
	}
	filePath := args[0]

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		println("File \"" + filePath + "\" tidak ada!")
		os.Exit(1)
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		println("File \"" + filePath + "\" tidak dapat diakses!")
		os.Exit(1)
	}

	jalankan(string(file), func(s string) {
		fmt.Print(s)
	})
}
