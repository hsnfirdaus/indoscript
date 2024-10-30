package main

import (
	"indoscript/lekser"
	"indoscript/penerjemah"
	"indoscript/pengurai"
	"indoscript/utils"
)

func main() {
	mulaiInterpreter()
}

func jalankan(kode string, fnKeluaran func(string)) {
	lek := lekser.LekserBaru(string(kode))
	tokenToken, lekErr := lek.Tokenisasi()
	if lekErr != nil {
		utils.CetakError(lekErr.Error())
	}

	par := pengurai.PenguraiBaru(tokenToken)
	ast, parErr := par.Urai()
	if parErr != nil {
		utils.CetakError(parErr.Error())
	}
	rt := penerjemah.PenerjemahBaru(fnKeluaran)
	terErr := rt.Jalankan(ast)
	if terErr != nil {
		utils.CetakError(terErr.Error())
	}
}
