package penerjemah

import (
	"fmt"
	"indoscript/utils"
)

type KesalahanPenerjemah struct {
	utils.BasisPosisi
	pesan string
}

func (k *KesalahanPenerjemah) Error() string {

	return fmt.Sprint("[PENERJEMAH ", k.Baris, ":", k.Kolom, "] ", k.pesan)
}
