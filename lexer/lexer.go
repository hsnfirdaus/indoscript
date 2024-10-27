package lexer

import (
	"indoscript/utils"
	"slices"
	"strconv"
)

type Lexer struct {
	teks            string
	indeks          int
	baris           int
	kolom           int
	karakterSaatIni string
}

func Baru(teks string) Lexer {
	lex := Lexer{
		teks:            teks,
		indeks:          -1,
		baris:           1,
		kolom:           -1,
		karakterSaatIni: "",
	}
	lex.maju()

	return lex
}

var ANGKA = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var HURUF = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
var HURUF_ANGKA = append(ANGKA, HURUF...)
var PENGENAL_VALID = append(HURUF_ANGKA, []string{"_"}...)

const DES_SEPARATOR = "."

func (lex *Lexer) ambilBasePosisi() utils.BasePosisi {
	return utils.BasePosisi{
		Baris: lex.baris,
		Kolom: lex.kolom,
	}
}
func (lex *Lexer) maju() {
	lex.indeks = lex.indeks + 1
	if lex.indeks < len(lex.teks) {
		lex.karakterSaatIni = string(lex.teks[lex.indeks])
	} else {
		lex.karakterSaatIni = ""
	}
	if lex.karakterSaatIni == "\n" {
		lex.baris = lex.baris + 1
		lex.kolom = 0
	} else {
		lex.kolom = lex.kolom + 1
	}
}

func (lex *Lexer) Tokenisasi() ([]Token, *Kesalahan) {
	var tokenToken []Token = make([]Token, 0)

	for lex.karakterSaatIni != "" {
		if lex.karakterSaatIni == "\t" || lex.karakterSaatIni == " " || lex.karakterSaatIni == "\n" {

		} else if slices.Contains(ANGKA, lex.karakterSaatIni) {
			tokenToken = append(tokenToken, lex.tokenisasiAngka())
			continue
		} else if slices.Contains(HURUF, lex.karakterSaatIni) {
			tokenToken = append(tokenToken, lex.tokenisasiKataKunci())
			continue
		} else if lex.karakterSaatIni == "'" || lex.karakterSaatIni == "\"" {
			tokenToken = append(tokenToken, lex.tokenisasiTeks())
			continue
		} else if lex.karakterSaatIni == "=" {
			tokenToken = append(tokenToken, lex.tokenisasiDenganSelanjutnya(T_SD, "=", T_SDSD))
			continue
		} else if lex.karakterSaatIni == "<" {
			tokenToken = append(tokenToken, lex.tokenisasiDenganSelanjutnya(T_KDAR, "=", T_KDARSD))
			continue
		} else if lex.karakterSaatIni == ">" {
			tokenToken = append(tokenToken, lex.tokenisasiDenganSelanjutnya(T_LDAR, "=", T_LDARSD))
			continue
		} else if lex.karakterSaatIni == "!" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_TIDAK,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == "+" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_TAMBAH,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == "-" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_KURANG,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == "*" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_KALI,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == "/" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_BAGI,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == "^" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_PANGKAT,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == "(" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_BKURUNG,
				BasePosisi: lex.ambilBasePosisi(),
			})
		} else if lex.karakterSaatIni == ")" {
			tokenToken = append(tokenToken, Token{
				Jenis:      T_TKURUNG,
				BasePosisi: lex.ambilBasePosisi(),
			})

		} else if lex.karakterSaatIni == ";" {
			tokenToken = append(tokenToken, Token{
				BasePosisi: lex.ambilBasePosisi(),
				Jenis:      T_TK,
			})
		} else if lex.karakterSaatIni == "," {
			tokenToken = append(tokenToken, Token{
				BasePosisi: lex.ambilBasePosisi(),
				Jenis:      T_KOMA,
			})
		} else if lex.karakterSaatIni == "{" {
			tokenToken = append(tokenToken, Token{
				BasePosisi: lex.ambilBasePosisi(),
				Jenis:      T_BKURAWAL,
			})
		} else if lex.karakterSaatIni == "}" {
			tokenToken = append(tokenToken, Token{
				BasePosisi: lex.ambilBasePosisi(),
				Jenis:      T_TKURAWAL,
			})
		} else {
			return tokenToken, &Kesalahan{BasePosisi: lex.ambilBasePosisi(), detail: "Karakter tidak valid: " + lex.karakterSaatIni}
		}

		lex.maju()
	}

	tokenToken = append(tokenToken, Token{
		Jenis:      T_ADF,
		BasePosisi: lex.ambilBasePosisi(),
	})

	return tokenToken, nil
}

func (lex *Lexer) tokenisasiAngka() Token {
	angkaTeks := ""
	jmlKoma := 0
	baris := lex.baris
	kolom := lex.kolom

	for lex.karakterSaatIni != "" && (slices.Contains(ANGKA, lex.karakterSaatIni) || lex.karakterSaatIni == DES_SEPARATOR) {
		if lex.karakterSaatIni == DES_SEPARATOR {
			if jmlKoma == 1 {
				break
			}
			jmlKoma = jmlKoma + 1
			angkaTeks = angkaTeks + "."
		} else {
			angkaTeks = angkaTeks + lex.karakterSaatIni
		}
		lex.maju()
	}

	if jmlKoma == 0 {
		valInt, _ := strconv.Atoi(angkaTeks)
		return Token{
			Jenis: T_BUL,
			Isi:   float64(valInt),
			BasePosisi: utils.BasePosisi{
				Baris: baris,
				Kolom: kolom,
			},
		}
	} else {
		valFloat, _ := strconv.ParseFloat(angkaTeks, 64)
		return Token{
			Jenis: T_DES,
			Isi:   valFloat,
			BasePosisi: utils.BasePosisi{
				Baris: baris,
				Kolom: kolom,
			},
		}
	}
}

func (lex *Lexer) tokenisasiKataKunci() Token {
	pengenal := ""
	baris := lex.baris
	kolom := lex.kolom

	for slices.Contains(PENGENAL_VALID, lex.karakterSaatIni) {
		pengenal = pengenal + lex.karakterSaatIni
		lex.maju()
	}

	var jenisToken JenisToken
	if slices.Contains(KATKUN, pengenal) {
		jenisToken = T_KATKUN
	} else {
		jenisToken = T_PENGENAL
	}

	return Token{
		Jenis: jenisToken,
		Isi:   pengenal,
		BasePosisi: utils.BasePosisi{
			Baris: baris,
			Kolom: kolom,
		},
	}
}

func (lex *Lexer) tokenisasiTeks() Token {
	teks := ""
	pemisah := lex.karakterSaatIni

	lex.maju()
	escaped := false
	for {
		teks = teks + lex.karakterSaatIni
		lex.maju()
		if lex.karakterSaatIni == "\\" {
			escaped = true
			lex.maju()
			continue
		}
		if lex.karakterSaatIni == pemisah && !escaped {
			break
		}
		if escaped {
			escaped = false
		}
	}
	lex.maju()

	return Token{
		Jenis: T_TEKS,
		Isi:   teks,
	}
}

func (lex *Lexer) tokenisasiDenganSelanjutnya(jenisAsli JenisToken, karakterSelanjutnya string, jenisKarakterSelanjutnya JenisToken) Token {
	basePos := lex.ambilBasePosisi()
	lex.maju()
	if lex.karakterSaatIni == karakterSelanjutnya {
		lex.maju()
		return Token{
			BasePosisi: basePos,
			Jenis:      jenisKarakterSelanjutnya,
		}
	}

	return Token{
		BasePosisi: basePos,
		Jenis:      jenisAsli,
	}
}
