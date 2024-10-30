//go:build wasm
// +build wasm

package fungsi

import (
	"indoscript/penerjemah/jenis"
	"syscall/js"
)

func Masukan(fnKeluaran func(string), argument []interface{}) (*jenis.Teks, error) {
	window := js.Global()

	var label string

	Cetak(func(s string) {
		label = s
	}, argument)

	result := window.Call("prompt", label).String()

	teks := &jenis.Teks{
		Teks: result,
	}

	keluaran := append(argument, teks, &jenis.Teks{
		Teks: "\n",
	})

	Cetak(fnKeluaran, keluaran)

	return teks, nil
}
