package penerjemah

import (
	"indoscript/lekser"
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
	"indoscript/utils"
	"slices"
)

func (p *Penerjemah) nodeOperasi(node *pengurai.NodeOperasi) (interface{}, *KesalahanPenerjemah) {
	kiri, err := p.panggilNodeBilangan(node.NodeKiri)
	if err != nil {
		return nil, err
	}
	kanan, err := p.panggilNodeBilangan(node.NodeKanan)
	if err != nil {
		return nil, err
	}

	if slices.Contains(lekser.TOKEN_PERBANDINGAN, node.Operasi) {
		res, er := kiri.OperasiBoolean(kanan, node.Operasi)
		if er != nil {
			return nil, &KesalahanPenerjemah{
				BasisPosisi: utils.BasisPosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: er.Error(),
			}
		}

		return res, nil
	}

	bil, er := kiri.Operasi(kanan, node.Operasi)
	if er != nil {
		return nil, &KesalahanPenerjemah{
			BasisPosisi: utils.BasisPosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: er.Error(),
		}
	}

	return bil, nil
}
func (p *Penerjemah) nodeOperasiUner(node *pengurai.NodeOperasiUner) (*jenis.Bilangan, *KesalahanPenerjemah) {
	angka, err := p.panggilNodeBilangan(node.Node)
	if err != nil {
		return nil, err
	}

	if node.Token == lekser.T_KURANG {
		bil, err := angka.Operasi(&jenis.Bilangan{Angka: -1}, lekser.T_KALI)
		if err != nil {
			return nil, &KesalahanPenerjemah{
				BasisPosisi: utils.BasisPosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: err.Error(),
			}
		}
		return bil, nil
	}

	return angka, nil
}

func (p *Penerjemah) nodeOperasiDanAtau(node *pengurai.NodeOperasiDanAtau) (*jenis.Boolean, *KesalahanPenerjemah) {
	kiri, err := p.panggilNodeBoolean(node.NodeKiri)
	if err != nil {
		return nil, err
	}
	kanan, err := p.panggilNodeBoolean(node.NodeKanan)
	if err != nil {
		return nil, err
	}

	switch node.Operasi {
	case lekser.KK_ATAU:
		return &jenis.Boolean{
			Isi: kiri.Isi || kanan.Isi,
		}, nil
	case lekser.KK_DAN:
		return &jenis.Boolean{
			Isi: kiri.Isi && kanan.Isi,
		}, nil
	}

	return nil, &KesalahanPenerjemah{
		BasisPosisi: utils.BasisPosisi{
			Baris: node.Baris,
			Kolom: node.Kolom,
		},
		pesan: "Operasi " + node.Operasi + " tidak valid!",
	}
}
