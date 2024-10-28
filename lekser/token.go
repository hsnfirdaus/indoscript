package lekser

import "indoscript/utils"

type JenisToken string

const (
	T_BUL JenisToken = "BUL"
	T_DES JenisToken = "DES"

	T_TAMBAH  JenisToken = "TAMBAH"
	T_KURANG  JenisToken = "KURANG"
	T_KALI    JenisToken = "KALI"
	T_BAGI    JenisToken = "BAGI"
	T_PANGKAT JenisToken = "PANGKAT"

	T_BKURUNG  JenisToken = "BKURUNG" // Buka Kurung
	T_TKURUNG  JenisToken = "TKURUNG" // Tutup Kurung
	T_BKURAWAL JenisToken = "BKURAWAL"
	T_TKURAWAL JenisToken = "TKURAWAL"

	T_KDAR   JenisToken = "KDAR"   // Kurang Dari
	T_LDAR   JenisToken = "LDAR"   // Lebih Dari
	T_KDARSD JenisToken = "KDARSD" // Kurang Dari Sama Dengan
	T_LDARSD JenisToken = "LDARSD" // Lebih Dari Sama Dengan
	T_SDSD   JenisToken = "SDSD"   // Sama Dengan Sama Dengan
	T_TDSD   JenisToken = "TDSD"   // Tidak Sama Dengan
	T_TIDAK  JenisToken = "TIDAK"

	T_KATKUN   JenisToken = "KATKUN"   // Kata Kunci / keyword
	T_PENGENAL JenisToken = "PENGENAL" // Identifier
	T_SD       JenisToken = "SD"       // Samadengan (EQ)
	T_TK       JenisToken = "TK"       // Titik Koma
	T_KOMA     JenisToken = "KOMA"
	T_TEKS     JenisToken = "TEKS"
	T_ADF      JenisToken = "ADF" // Akhir Dari File

)

var TOKEN_PERBANDINGAN = []JenisToken{
	T_KDAR,
	T_LDAR,
	T_KDARSD,
	T_LDARSD,
	T_SDSD,
	T_TDSD,
}

type Token struct {
	utils.BasisPosisi
	Jenis JenisToken
	Isi   interface{}
}
