package main

import (
	"errors"
	"indoscript/lexer"
	"indoscript/parser"
	"indoscript/runtime"
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
	lex := lexer.Baru(string(file))
	tokenToken, lexErr := lex.Tokenisasi()
	if lexErr != nil {
		println(lexErr.Error())
		os.Exit(1)
	}

	par := parser.Baru(tokenToken)
	ast, parErr := par.Parse()
	if parErr != nil {
		println(parErr.Error())
		os.Exit(1)
	}
	rt := runtime.Runtime{}
	rt.Baru()
	terErr := rt.Jalankan(ast)
	if terErr != nil {
		println(terErr.Error())
		os.Exit(1)
	}
}
