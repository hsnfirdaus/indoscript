package penerjemah

import (
	"fmt"
	"indoscript/penerjemah/fungsi"
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
	"indoscript/utils"
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

	default:
		fung, _ := p.ts.ambilFung(node.NamaFungsi)
		if fung == nil {
			return nil, &KesalahanPenerjemah{
				BasisPosisi: utils.BasisPosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: "Tidak ditemukan fungsi dengan nama \"" + node.NamaFungsi + "\"",
			}
		}

		hasil, err := p.lakukanPanggilFungsiTS(fung, hasilArgumen)
		if err != nil {
			return nil, &KesalahanPenerjemah{
				BasisPosisi: utils.BasisPosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: "Fungsi \"" + node.NamaFungsi + "\": " + err.Error(),
			}
		}
		return hasil, nil

	}

	if err != nil {
		return nil, &KesalahanPenerjemah{
			BasisPosisi: utils.BasisPosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: err.Error(),
		}
	}

	return hasilPanggil, nil
}

func (p *Penerjemah) lakukanPanggilFungsiTS(fung *ButirFungsi, argumen []interface{}) (interface{}, *KesalahanPenerjemah) {
	if len(argumen) != len(fung.namaArgument) {
		return nil, &KesalahanPenerjemah{
			pesan: fmt.Sprint("Argumen dibutuhkan ", len(fung.namaArgument), " diberikan ", len(argumen)),
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
