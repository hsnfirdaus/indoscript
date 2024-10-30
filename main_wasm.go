//go:build wasm
// +build wasm

package main

import "syscall/js"

func jalankanKode(this js.Value, arument []js.Value) interface{} {
	kode := arument[0].String()

	invoker := func(teks string) {
		arument[1].Invoke(teks)
	}

	jalankan(kode, invoker)

	return nil
}

func mulaiInterpreter() {
	ch := make(chan struct{}, 0)

	js.Global().Set("jalankanKode", js.FuncOf(jalankanKode))

	<-ch
}
