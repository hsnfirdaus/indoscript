//go:build !wasm
// +build !wasm

package fungsi

import (
	"bufio"
	"indoscript/penerjemah/jenis"
	"os"
)

func Masukan(fnKeluaran func(string), argument []interface{}) (*jenis.Teks, error) {
	Cetak(fnKeluaran, argument)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	teks := scanner.Text()

	return &jenis.Teks{
		Teks: teks,
	}, nil
}
