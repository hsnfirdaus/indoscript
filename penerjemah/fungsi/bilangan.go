package fungsi

import (
	"errors"
	"indoscript/penerjemah/jenis"
	"reflect"
	"strconv"
)

func KeBilangan(jenisLain interface{}) (*jenis.Bilangan, error) {
	switch v := jenisLain.(type) {
	case *jenis.Bilangan:
		return v, nil

	case *jenis.Teks:
		angka, err := strconv.ParseFloat(v.Teks, 64)
		if err != nil {
			return nil, errors.New("Gagal melakukan konversi ke bilangan: angka tidak valid!")
		}
		return &jenis.Bilangan{
			Angka: angka,
		}, nil

	case *jenis.Boolean:
		var angka float64
		if v.Isi {
			angka = 1
		} else {
			angka = 0
		}
		return &jenis.Bilangan{
			Angka: angka,
		}, nil

	default:
		return nil, errors.New("Tidak dapat mengkonversi \"" + reflect.ValueOf(v).String() + "\" menjadi bilangan!")
	}

}
