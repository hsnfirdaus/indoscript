package lekser

import "testing"

func TestTokenisasiAngkaBulat(t *testing.T) {
	lek := LekserBaru("120000 xw")

	token := lek.tokenisasiAngka()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 0 {
		t.Error("Kolom bukan 0! ", token.Kolom)
	}

	if token.Jenis != T_BUL {
		t.Error("Jenis token bukan T_BUL! ", token.Jenis)
	}

	if token.Isi != float64(120000) {
		t.Error("Isi bukanlah 120000! ", token.Isi)
	}
}

func TestTokenisasiAngkaDesimal(t *testing.T) {
	lek := LekserBaru("12.35 xw")

	token := lek.tokenisasiAngka()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 0 {
		t.Error("Kolom bukan 0! ", token.Kolom)
	}

	if token.Jenis != T_DES {
		t.Error("Jenis token bukan T_DES! ", token.Jenis)
	}

	if token.Isi != float64(12.35) {
		t.Error("Isi bukanlah 12.35! ", token.Isi)
	}
}

func TestTokenisasiAngkaNest(t *testing.T) {
	lek := LekserBaru("12.35 10000")

	lek.tokenisasiAngka()

	lek.maju()

	token := lek.tokenisasiAngka()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 6 {
		t.Error("Kolom bukan 6! ", token.Kolom)
	}

	if token.Jenis != T_BUL {
		t.Error("Jenis token bukan T_BUL! ", token.Jenis)
	}

	if token.Isi != float64(10000) {
		t.Error("Isi bukanlah 10000! ", token.Isi)
	}
}
