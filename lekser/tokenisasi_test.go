package lekser

import (
	"testing"
)

func TestTokenisasiLanjut(t *testing.T) {
	lek := LekserBaru(" >= ")
	lek.maju()

	token := lek.tokenisasiLanjut(T_LDAR, "=", T_LDARSD)

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 1 {
		t.Error("Kolom bukan 1! ", token.Kolom)
	}

	if token.Jenis != T_LDARSD {
		t.Error("Jenis token bukan T_LDARSD! ", token.Jenis)
	}

}

func TestTokenisasiLanjutAsli(t *testing.T) {
	lek := LekserBaru(" = ")
	lek.maju()

	token := lek.tokenisasiLanjut(T_SD, "=", T_SDSD)

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 1 {
		t.Error("Kolom bukan 1! ", token.Kolom)
	}

	if token.Jenis != T_SD {
		t.Error("Jenis token bukan T_SD! ", token.Jenis)
	}
}
