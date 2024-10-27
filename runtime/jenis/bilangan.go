package jenis

import (
	"errors"
	"fmt"
	"indoscript/lexer"
	"math"
)

type Bilangan struct {
	Angka float64
}

func (b *Bilangan) Operasi(targetBil *Bilangan, op lexer.JenisToken) error {

	switch op {
	case lexer.T_TAMBAH:
		b.Angka = b.Angka + targetBil.Angka
		return nil

	case lexer.T_KURANG:
		b.Angka = b.Angka - targetBil.Angka
		return nil

	case lexer.T_KALI:
		b.Angka = b.Angka * targetBil.Angka
		return nil

	case lexer.T_BAGI:
		b.Angka = b.Angka / targetBil.Angka
		return nil

	case lexer.T_PANGKAT:
		b.Angka = math.Pow(b.Angka, targetBil.Angka)
		return nil
	}

	return errors.New(fmt.Sprint("Operasi tak terduga: ", op))
}

func (b *Bilangan) OperasiBoolean(targetBil *Bilangan, op lexer.JenisToken) (*Boolean, error) {

	var hasil bool
	switch op {
	case lexer.T_KDAR:
		hasil = b.Angka < targetBil.Angka

	case lexer.T_LDAR:
		hasil = b.Angka > targetBil.Angka

	case lexer.T_KDARSD:
		hasil = b.Angka <= targetBil.Angka

	case lexer.T_LDARSD:
		hasil = b.Angka >= targetBil.Angka

	case lexer.T_SDSD:
		hasil = b.Angka == targetBil.Angka

	default:
		return nil, errors.New(fmt.Sprint("Operasi tak terduga: ", op))
	}

	return &Boolean{
		Isi: hasil,
	}, nil

}
