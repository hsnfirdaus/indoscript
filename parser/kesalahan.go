package parser

import (
	"fmt"
	"indoscript/lexer"
	"indoscript/utils"
)

type TokenTakTerduga struct {
	utils.BasePosisi
	diharapkan []lexer.JenisToken
	ditemukan  lexer.JenisToken
}

func (t *TokenTakTerduga) Error() string {

	return fmt.Sprint("[PARSER] [B", t.Baris, "K", t.Kolom, "] Token tak terduga, diharapkan \"", t.diharapkan, "\" ditemukan \"", t.ditemukan, "\"")
}
