package lekser

import (
	"slices"
	"strconv"
)

func (lek *Lekser) tokenisasiAngka() Token {
	angkaTeks := ""
	jmlKoma := 0
	posisi := lek.basisPosisi()

	for lek.karakterSaatIni != "" && (slices.Contains(ANGKA, lek.karakterSaatIni) || lek.karakterSaatIni == DES_SEPARATOR) {
		if lek.karakterSaatIni == DES_SEPARATOR {
			if jmlKoma == 1 {
				break
			}
			jmlKoma = jmlKoma + 1
			angkaTeks = angkaTeks + "."
		} else {
			angkaTeks = angkaTeks + lek.karakterSaatIni
		}
		lek.maju()
	}

	if jmlKoma == 0 {
		valInt, _ := strconv.Atoi(angkaTeks)
		return Token{
			Jenis:       T_BUL,
			Isi:         float64(valInt),
			BasisPosisi: posisi,
		}
	} else {
		valFloat, _ := strconv.ParseFloat(angkaTeks, 64)
		return Token{
			Jenis:       T_DES,
			Isi:         valFloat,
			BasisPosisi: posisi,
		}
	}
}
