package lekser

import "slices"

func (lek *Lekser) Tokenisasi() ([]Token, *KesalahanLekser) {
	var tokenToken []Token = make([]Token, 0)

	for lek.karakterSaatIni != "" {

		if slices.Contains(ANGKA, lek.karakterSaatIni) {
			tokenToken = append(tokenToken, lek.tokenisasiAngka())
			continue
		} else if slices.Contains(HURUF, lek.karakterSaatIni) {
			tokenToken = append(tokenToken, lek.tokenisasiKataKunci())
			continue
		} else {

			switch lek.karakterSaatIni {
			case "\t", " ", "\n":

			case "'", "\"":
				tokenToken = append(tokenToken, lek.tokenisasiTeks())
				continue

			case "=":
				tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_SD, "=", T_SDSD))
				continue

			case "<":
				tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_KDAR, "=", T_KDARSD))
				continue

			case ">":
				tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_LDAR, "=", T_LDARSD))
				continue

			case "!":
				tokenToken = append(tokenToken, lek.tokenisasiLanjut(T_TIDAK, "=", T_TDSD))
				continue

			case "+":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_TAMBAH,
					BasisPosisi: lek.basisPosisi(),
				})

			case "-":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_KURANG,
					BasisPosisi: lek.basisPosisi(),
				})

			case "*":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_KALI,
					BasisPosisi: lek.basisPosisi(),
				})
			case "/":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_BAGI,
					BasisPosisi: lek.basisPosisi(),
				})

			case "^":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_PANGKAT,
					BasisPosisi: lek.basisPosisi(),
				})

			case "(":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_BKURUNG,
					BasisPosisi: lek.basisPosisi(),
				})

			case ")":
				tokenToken = append(tokenToken, Token{
					Jenis:       T_TKURUNG,
					BasisPosisi: lek.basisPosisi(),
				})

			case ";":
				tokenToken = append(tokenToken, Token{
					BasisPosisi: lek.basisPosisi(),
					Jenis:       T_TK,
				})

			case ",":
				tokenToken = append(tokenToken, Token{
					BasisPosisi: lek.basisPosisi(),
					Jenis:       T_KOMA,
				})

			case "{":
				tokenToken = append(tokenToken, Token{
					BasisPosisi: lek.basisPosisi(),
					Jenis:       T_BKURAWAL,
				})

			case "}":
				tokenToken = append(tokenToken, Token{
					BasisPosisi: lek.basisPosisi(),
					Jenis:       T_TKURAWAL,
				})

			default:
				return nil, &KesalahanLekser{
					BasisPosisi: lek.basisPosisi(),
					detail:      "Karakter tidak valid: " + lek.karakterSaatIni,
				}

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
