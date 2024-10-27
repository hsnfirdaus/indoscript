package main

import (
	"errors"
	"indoscript/lekser"
	"indoscript/penerjemah"
	"indoscript/pengurai"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		println("Anda harus memasukan path file sebagai argumen!")
		os.Exit(1)
	}
	filePath := args[0]

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		println("File \"" + filePath + "\" tidak ada!")
		os.Exit(1)
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		println("File \"" + filePath + "\" tidak dapat diakses!")
		os.Exit(1)
	}
	lek := lekser.LekserBaru(string(file))
	tokenToken, lekErr := lek.Tokenisasi()
	if lekErr != nil {
		println(lekErr.Error())
		os.Exit(1)
	}

	par := pengurai.PenguraiBaru(tokenToken)
	ast, parErr := par.Urai()
	if parErr != nil {
		println(parErr.Error())
		os.Exit(1)
	}
	rt := penerjemah.PenerjemahBaru()
	terErr := rt.Jalankan(ast)
	if terErr != nil {
		println(terErr.Error())
		os.Exit(1)
	}
}
