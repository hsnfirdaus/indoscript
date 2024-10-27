package pengurai

import (
	"fmt"
	"indoscript/lekser"
	"indoscript/utils"
)

type TokenTakTerduga struct {
	utils.BasisPosisi
	diharapkan []lekser.JenisToken
	ditemukan  lekser.JenisToken
}

func (t *TokenTakTerduga) Error() string {

	return fmt.Sprint("[PENGURAI ", t.Baris, ":", t.Kolom, "] Token tak terduga, diharapkan \"", t.diharapkan, "\", ditemukan \"", t.ditemukan, "\"!")
}
