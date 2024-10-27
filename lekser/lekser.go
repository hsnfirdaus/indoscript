package lekser

type Lekser struct {
	teks            string
	indeks          int
	baris           int
	kolom           int
	karakterSaatIni string
}

func LekserBaru(teks string) Lekser {
	lek := Lekser{
		teks:            teks,
		indeks:          -1,
		baris:           1,
		kolom:           -1,
		karakterSaatIni: "",
	}
	lek.maju()

	return lek
}

var ANGKA = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var HURUF = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}
var HURUF_ANGKA = append(ANGKA, HURUF...)
var PENGENAL_VALID = append(HURUF_ANGKA, []string{"_"}...)

const DES_SEPARATOR = "."
