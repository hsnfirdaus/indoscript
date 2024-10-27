package lekser

import "testing"

func TestTokenisasiKatKun(t *testing.T) {
	lek := LekserBaru("var jika")

	token := lek.tokenisasiKataKunci()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 0 {
		t.Error("Kolom bukan 0! ", token.Kolom)
	}

	if token.Jenis != T_KATKUN {
		t.Error("Jenis token bukan T_KATKUN! ", token.Jenis)
	}

	if token.Isi != "var" {
		t.Error("Isi bukanlah 'var'! ", token.Isi)
	}

	lek.maju()

	token = lek.tokenisasiKataKunci()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 4 {
		t.Error("Kolom bukan 4! ", token.Kolom)
	}

	if token.Jenis != T_KATKUN {
		t.Error("Jenis token bukan T_KATKUN! ", token.Jenis)
	}

	if token.Isi != "jika" {
		t.Error("Isi bukanlah 'jika'! ", token.Isi)
	}
}

func TestTokenisasiKatKunPengenal(t *testing.T) {
	lek := LekserBaru("contoh_variabel")

	token := lek.tokenisasiKataKunci()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 0 {
		t.Error("Kolom bukan 0! ", token.Kolom)
	}

	if token.Jenis != T_PENGENAL {
		t.Error("Jenis token bukan T_PENGENAL! ", token.Jenis)
	}

	if token.Isi != "contoh_variabel" {
		t.Error("Isi bukanlah 'contoh_variabel'! ", token.Isi)
	}
}

func TestTokenisasiKatKunCampuran(t *testing.T) {
	lek := LekserBaru("var nama_variabel")

	token := lek.tokenisasiKataKunci()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 0 {
		t.Error("Kolom bukan 0! ", token.Kolom)
	}

	if token.Jenis != T_KATKUN {
		t.Error("Jenis token bukan T_KATKUN! ", token.Jenis)
	}

	if token.Isi != "var" {
		t.Error("Isi bukanlah 'var'! ", token.Isi)
	}

	lek.maju()

	token = lek.tokenisasiKataKunci()

	if token.Baris != 1 {
		t.Error("Baris bukan 1! ", token.Baris)
	}
	if token.Kolom != 4 {
		t.Error("Kolom bukan 4! ", token.Kolom)
	}

	if token.Jenis != T_PENGENAL {
		t.Error("Jenis token bukan T_PENGENAL! ", token.Jenis)
	}

	if token.Isi != "nama_variabel" {
		t.Error("Isi bukanlah 'nama_variabel'! ", token.Isi)
	}
}
