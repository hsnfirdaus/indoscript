package jenis

import (
	"errors"
	"fmt"
	"indoscript/lekser"
)

type Teks struct {
	Teks string
}

func (t *Teks) OperasiBoolean(targetTeks *Teks, op lekser.JenisToken) (*Boolean, error) {

	var hasil bool
	switch op {
	case lekser.T_SDSD:
		hasil = t.Teks == targetTeks.Teks

	case lekser.T_TDSD:
		hasil = t.Teks != targetTeks.Teks

	default:
		return nil, errors.New(fmt.Sprint("Operasi TEKS tak terduga: ", op))
	}

	return &Boolean{
		Isi: hasil,
	}, nil

}
