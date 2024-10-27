package lekser

import (
	"slices"
)

func (lek *Lekser) tokenisasiKataKunci() Token {
	pengenal := ""
	posisi := lek.basisPosisi()

	for slices.Contains(PENGENAL_VALID, lek.karakterSaatIni) {
		pengenal = pengenal + lek.karakterSaatIni
		lek.maju()
	}

	var jenisToken JenisToken
	if slices.Contains(KATKUN, pengenal) {
		jenisToken = T_KATKUN
	} else {
		jenisToken = T_PENGENAL
	}

	return Token{
		Jenis:       jenisToken,
		Isi:         pengenal,
		BasisPosisi: posisi,
	}
}
