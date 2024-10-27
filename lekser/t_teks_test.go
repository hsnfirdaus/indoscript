package lekser

import "testing"

func TestTokenisasiTeks(t *testing.T) {
	lek := LekserBaru("'ini teks'")

	token := lek.tokenisasiTeks()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 0 {
		t.Error("Kolom bukan 0! ", token.Kolom)
	}

	if token.Jenis != T_TEKS {
		t.Error("Jenis token bukan T_TEKS! ", token.Jenis)
	}

	if token.Isi != "ini teks" {
		t.Error("Isi bukanlah 'ini teks'! ", token.Isi)
	}
}

func TestTokenisasiTeksKutip(t *testing.T) {
	lek := LekserBaru("\"ini teks\"")

	token := lek.tokenisasiTeks()

	if token.Isi != "ini teks" {
		t.Error("Isi bukanlah 'ini teks'! ", token.Isi)
	}
}

func TestTokenisasiTeksEscape(t *testing.T) {
	lek := LekserBaru("'Ma\\'af'")

	token := lek.tokenisasiTeks()

	if token.Isi != "Ma'af" {
		t.Error("Isi bukanlah 'Ma'af'! ", token.Isi)
	}

	lek = LekserBaru("'Ma\\\\'af'")
	token = lek.tokenisasiTeks()

	if token.Isi != "Ma\\'af" {
		t.Error("Isi bukanlah 'Ma\\'af'! ", token.Isi)
	}
}
