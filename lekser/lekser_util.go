package lekser

import "indoscript/utils"

func (lek *Lekser) basisPosisi() utils.BasisPosisi {
	return utils.BasisPosisi{
		Baris: lek.baris,
		Kolom: lek.kolom,
	}
}
func (lek *Lekser) maju() {
	lek.indeks = lek.indeks + 1
	if lek.indeks < len(lek.teks) {
		lek.karakterSaatIni = string(lek.teks[lek.indeks])
	} else {
		lek.karakterSaatIni = ""
	}
	if lek.karakterSaatIni == "\n" {
		lek.baris = lek.baris + 1
		lek.kolom = -1
	} else {
		lek.kolom = lek.kolom + 1
	}
}
