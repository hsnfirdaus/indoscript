package lekser

import "testing"

func TestMaju(t *testing.T) {
	str := "a9-1n \n#"

	lek := LekserBaru(str)
	if lek.indeks != 0 {
		t.Error("Indeks pertama bukan 0!")
	}

	if lek.karakterSaatIni != "a" {
		t.Error("Karakter pertama tidak valid!")
	}

	if lek.baris != 1 {
		t.Error("Baris pertama bukanlah 1!")
	}

	if lek.kolom != 0 {
		t.Error("Kolom pertama bukanlah 0!")
	}

	for i := 0; i < 6; i++ {
		lek.maju()
	}

	if lek.indeks != 6 {
		t.Error("Indeks bukan 6!", lek.indeks)
	}

	if lek.karakterSaatIni != "\n" {
		t.Error("Karakter bukan \\n!", lek.karakterSaatIni)
	}

	if lek.baris != 2 {
		t.Error("Baris bukanlah 2!", lek.baris)
	}

	if lek.kolom != -1 {
		t.Error("Kolom bukanlah -1!", lek.kolom)
	}

	lek.maju()

	if lek.indeks != 7 {
		t.Error("Indeks bukan 7!", lek.indeks)
	}

	if lek.karakterSaatIni != "#" {
		t.Error("Karakter bukan #!", lek.karakterSaatIni)
	}

	if lek.baris != 2 {
		t.Error("Baris bukanlah 2!", lek.baris)
	}

	if lek.kolom != 0 {
		t.Error("Kolom bukanlah 0!", lek.kolom)
	}
}
