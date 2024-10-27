package lekser

func (lek *Lekser) tokenisasiTeks() Token {
	teks := ""
	pemisah := lek.karakterSaatIni
	posisi := lek.basisPosisi()

	lek.maju()
	escaped := false
	for {
		teks = teks + lek.karakterSaatIni
		lek.maju()
		if lek.karakterSaatIni == "\\" {
			escaped = true
			lek.maju()
			continue
		}
		if lek.karakterSaatIni == pemisah && !escaped {
			break
		}
		if escaped {
			escaped = false
		}
	}
	lek.maju()

	return Token{
		BasisPosisi: posisi,
		Jenis:       T_TEKS,
		Isi:         teks,
	}
}
