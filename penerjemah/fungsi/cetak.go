package fungsi

import (
	"errors"
	"fmt"
	"indoscript/penerjemah/jenis"
)

func Cetak(argument []interface{}) (interface{}, error) {
	for _, arg := range argument {
		switch v := arg.(type) {
		case *jenis.Bilangan:
			fmt.Printf("%v", v.Angka)
			break

		case *jenis.Teks:
			fmt.Printf("%v", v.Teks)

		default:
			return nil, errors.New("Tidak dapat mencetak jenis ini")
		}
	}

	return nil, nil
}

func CetakBr(argument []interface{}) (interface{}, error) {
	hasil, err := Cetak(argument)
	if err != nil {
		return nil, err
	}
	println()

	return hasil, nil
}