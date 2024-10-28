package jenis

import (
	"errors"
	"fmt"
	"indoscript/lekser"
	"math"
)

type Bilangan struct {
	Angka float64
}

func (b *Bilangan) Operasi(targetBil *Bilangan, op lekser.JenisToken) (*Bilangan, error) {
	var hasil float64

	switch op {
	case lekser.T_TAMBAH:
		hasil = b.Angka + targetBil.Angka

	case lekser.T_KURANG:
		hasil = b.Angka - targetBil.Angka

	case lekser.T_KALI:
		hasil = b.Angka * targetBil.Angka

	case lekser.T_BAGI:
		hasil = b.Angka / targetBil.Angka

	case lekser.T_PANGKAT:
		hasil = math.Pow(b.Angka, targetBil.Angka)

	default:
		return nil, errors.New(fmt.Sprint("Operasi tak terduga: ", op))

	}

	return &Bilangan{
		Angka: hasil,
	}, nil
}

func (b *Bilangan) OperasiBoolean(targetBil *Bilangan, op lekser.JenisToken) (*Boolean, error) {

	var hasil bool
	switch op {
	case lekser.T_KDAR:
		hasil = b.Angka < targetBil.Angka

	case lekser.T_LDAR:
		hasil = b.Angka > targetBil.Angka

	case lekser.T_KDARSD:
		hasil = b.Angka <= targetBil.Angka

	case lekser.T_LDARSD:
		hasil = b.Angka >= targetBil.Angka

	case lekser.T_SDSD:
		hasil = b.Angka == targetBil.Angka

	case lekser.T_TDSD:
		hasil = b.Angka != targetBil.Angka

	default:
		return nil, errors.New(fmt.Sprint("Operasi BILANGAN tak terduga: ", op))
	}

	return &Boolean{
		Isi: hasil,
	}, nil

}
