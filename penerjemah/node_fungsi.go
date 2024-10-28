package penerjemah

import (
	"fmt"
	"indoscript/penerjemah/fungsi"
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeAturFungsi(node *pengurai.NodeAturFungsi) (interface{}, *KesalahanPenerjemah) {
	p.ts.aturFung(node.NamaFungsi, node.NamaArgumen, node.NodeNode)

	return nil, nil
}

func (p *Penerjemah) nodePanggilFungsi(node *pengurai.NodePanggilFungsi) (interface{}, *KesalahanPenerjemah) {
	var hasilArgumen = make([]interface{}, 0)

	for _, arg := range node.Argumen {
		hasil, err := p.panggilNode(arg)
		if err != nil {
			return nil, err
		}
		hasilArgumen = append(hasilArgumen, hasil)
	}

	var hasilPanggil interface{}
	var err error = nil

	switch node.NamaFungsi {
	case "cetak":
		hasilPanggil, err = fungsi.Cetak(hasilArgumen)

	case "cetakBr":
		hasilPanggil, err = fungsi.CetakBr(hasilArgumen)

	case "masukan":
		errArg := p.pastikanJumlahArgumen(node, len(hasilArgumen), 0)
		if errArg != nil {
			return nil, errArg
		}
		hasilPanggil, err = fungsi.Masukan()

	case "keBilangan":
		errArg := p.pastikanJumlahArgumen(node, len(hasilArgumen), 1)
		if errArg != nil {
			return nil, errArg
		}
		hasilPanggil, err = fungsi.KeBilangan(hasilArgumen[0])

	default:
		fung, _ := p.ts.ambilFung(node.NamaFungsi)
		if fung == nil {
			return nil, &KesalahanPenerjemah{
				BasisPosisi: node.BasisPosisi,
				pesan:       "Tidak ditemukan fungsi dengan nama \"" + node.NamaFungsi + "\"",
			}
		}

		hasil, err := p.lakukanPanggilFungsiTS(fung, hasilArgumen)
		if err != nil {
			return nil, &KesalahanPenerjemah{
				BasisPosisi: node.BasisPosisi,
				pesan:       "Fungsi \"" + node.NamaFungsi + "\": " + err.pesan,
			}
		}
		return hasil, nil

	}

	if err != nil {
		return nil, &KesalahanPenerjemah{
			BasisPosisi: node.BasisPosisi,
			pesan:       err.Error(),
		}
	}

	return hasilPanggil, nil
}

func (p *Penerjemah) pastikanJumlahArgumen(node *pengurai.NodePanggilFungsi, jumlahDiatur int, jumlahSeharusnya int) *KesalahanPenerjemah {
	if jumlahDiatur != jumlahSeharusnya {
		return &KesalahanPenerjemah{
			BasisPosisi: node.BasisPosisi,
			pesan:       fmt.Sprint("Fungsi \"", node.NamaFungsi, "\": argumen dibutuhkan ", jumlahSeharusnya, " diberikan ", jumlahDiatur),
		}
	}
	return nil
}

func (p *Penerjemah) lakukanPanggilFungsiTS(fung *ButirFungsi, argumen []interface{}) (interface{}, *KesalahanPenerjemah) {
	if len(argumen) != len(fung.namaArgument) {
		return nil, &KesalahanPenerjemah{
			pesan: fmt.Sprint("argumen dibutuhkan ", len(fung.namaArgument), " diberikan ", len(argumen)),
		}
	}
	pAnak := p.buatAnak()
	for idk, nama := range fung.namaArgument {
		pAnak.ts.aturVar(nama, argumen[idk])
	}

	hasil, err := pAnak.nodeAkar(fung.nodeAkar)
	if err != nil {
		return nil, err
	}
	balikan, apakahBalikan := hasil.(*jenis.Balikan)
	if apakahBalikan {
		return balikan.Unwrap(), nil
	}
	return hasil, nil
}
