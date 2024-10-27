package lexer

import (
	"fmt"
	"indoscript/utils"
)

type Kesalahan struct {
	utils.BasePosisi
	detail string
}

func (k *Kesalahan) Error() string {
	return fmt.Sprint("[LEXER] [B", k.Baris, "] [K", k.Kolom, "] ", k.detail)
}
