package runtime

import (
	"fmt"
	"indoscript/utils"
)

type KesalahanRuntime struct {
	utils.BasePosisi
	pesan string
}

func (k *KesalahanRuntime) Error() string {

	return fmt.Sprint("[RUNTIME] [B", k.Baris, "K", k.Kolom, "] ", k.pesan)
}
