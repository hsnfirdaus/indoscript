package fungsi

import (
	"bufio"
	"indoscript/penerjemah/jenis"
	"os"
)

func Masukan() (*jenis.Teks, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	teks := scanner.Text()

	return &jenis.Teks{
		Teks: teks,
	}, nil
}
