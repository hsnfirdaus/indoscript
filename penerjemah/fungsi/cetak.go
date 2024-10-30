package fungsi

import (
	"errors"
	"fmt"
	"indoscript/penerjemah/jenis"
)

func Cetak(fnKeluaran func(string), argument []interface{}) (interface{}, error) {
	teks := ""
	for _, arg := range argument {
		switch v := arg.(type) {
		case *jenis.Bilangan:
			teks = teks + fmt.Sprintf("%v", v.Angka)
			break

		case *jenis.Teks:
			teks = teks + fmt.Sprintf("%v", v.Teks)

		default:
			return nil, errors.New("Tidak dapat mencetak jenis ini")
		}
	}

	fnKeluaran(teks)

	return nil, nil
}

func CetakBr(fnKeluaran func(string), argument []interface{}) (interface{}, error) {
	hasil, err := Cetak(fnKeluaran, argument)
	if err != nil {
		return nil, err
	}
	newLn := &jenis.Teks{
		Teks: "\n",
	}
	hasil, err = Cetak(fnKeluaran, []interface{}{newLn})
	if err != nil {
		return nil, err
	}

	return hasil, nil
}
