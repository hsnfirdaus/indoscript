package jenis

import (
	"errors"
	"fmt"
	"indoscript/lekser"
)

type Boolean struct {
	Isi bool
}

func (b *Boolean) OperasiBoolean(targetBoolean *Boolean, op lekser.JenisToken) (*Boolean, error) {

	var hasil bool
	switch op {
	case lekser.T_SDSD:
		hasil = b.Isi == targetBoolean.Isi

	case lekser.T_TDSD:
		hasil = b.Isi != targetBoolean.Isi

	default:
		return nil, errors.New(fmt.Sprint("Operasi BOOLEAN tak terduga: ", op))
	}

	return &Boolean{
		Isi: hasil,
	}, nil

}
