package lekser

import (
	"fmt"
	"indoscript/utils"
)

type KesalahanLekser struct {
	utils.BasisPosisi
	detail string
}

func (k *KesalahanLekser) Error() string {
	return fmt.Sprint("[LEKSER ", k.Baris, ":", k.Kolom, "] ", k.detail)
}
