package lekser

import "slices"

func (lek *Lekser) Tokenisasi() ([]Token, *KesalahanLekser) {
	var tokenToken []Token = make([]Token, 0)

	for lek.karakterSaatIni != "" {
		if lek.karakterSaatIni == "\t" || lek.karakterSaatIni == " " || lek.karakterSaatIni == "\n" {

		} else if slices.Contains(ANGKA, lek.karakterSaatIni) {
			tokenToken = append(tokenToken, lek.tokenisasiAngka())
			continue
		} else if slices.Contains(HURUF, lek.karakterSaatIni) {
			tokenToken = append(tokenToken, lek.tokenisasiKataKunci())
			continue
		} else if lek.karakterSaatIni == "'" || lek.karakterSaatIni == "\"" {
			tokenToken = append(tokenToken, lek.tokenisasiTeks())
			continue
		} else if lek.karakterSaatIni == "=" {
			tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_SD, "=", T_SDSD))
			continue
		} else if lek.karakterSaatIni == "<" {
			tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_KDAR, "=", T_KDARSD))
			continue
		} else if lek.karakterSaatIni == ">" {
			tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_LDAR, "=", T_LDARSD))
			continue
		} else if lek.karakterSaatIni == "!" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_TIDAK,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == "+" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_TAMBAH,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == "-" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_KURANG,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == "*" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_KALI,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == "/" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_BAGI,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == "^" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_PANGKAT,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == "(" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_BKURUNG,
				BasisPosisi: lek.basisPosisi(),
			})
		} else if lek.karakterSaatIni == ")" {
			tokenToken = append(tokenToken, Token{
				Jenis:       T_TKURUNG,
				BasisPosisi: lek.basisPosisi(),
			})

		} else if lek.karakterSaatIni == ";" {
			tokenToken = append(tokenToken, Token{
				BasisPosisi: lek.basisPosisi(),
				Jenis:       T_TK,
			})
		} else if lek.karakterSaatIni == "," {
			tokenToken = append(tokenToken, Token{
				BasisPosisi: lek.basisPosisi(),
				Jenis:       T_KOMA,
			})
		} else if lek.karakterSaatIni == "{" {
			tokenToken = append(tokenToken, Token{
				BasisPosisi: lek.basisPosisi(),
				Jenis:       T_BKURAWAL,
			})
		} else if lek.karakterSaatIni == "}" {
			tokenToken = append(tokenToken, Token{
				BasisPosisi: lek.basisPosisi(),
				Jenis:       T_TKURAWAL,
			})
		} else {
			return nil, &KesalahanLekser{
				BasisPosisi: lek.basisPosisi(),
				detail:      "Karakter tidak valid: " + lek.karakterSaatIni,
			}
		}

		lek.maju()
	}

	tokenToken = append(tokenToken, Token{
		Jenis:       T_ADF,
		BasisPosisi: lek.basisPosisi(),
	})

	return tokenToken, nil
}
func (lek *Lekser) tokenisasiLanjut(jenisAsli JenisToken, karakterSelanjutnya string, jenisKarakterSelanjutnya JenisToken) Token {
	posisi := lek.basisPosisi()
	lek.maju()
	if lek.karakterSaatIni == karakterSelanjutnya {
		lek.maju()
		return Token{
			BasisPosisi: posisi,
			Jenis:       jenisKarakterSelanjutnya,
		}
	}

	return Token{
		BasisPosisi: posisi,
		Jenis:       jenisAsli,
	}
}
