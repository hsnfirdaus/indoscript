package pengurai

import (
	"indoscript/lekser"
	"indoscript/utils"
)

type Pengurai struct {
	daftarToken  []lekser.Token
	tokenSaatIni *lekser.Token
	indeks       int
}

func PenguraiBaru(daftarToken []lekser.Token) Pengurai {
	pengurai := Pengurai{
		daftarToken:  daftarToken,
		tokenSaatIni: nil,
		indeks:       -1,
	}
	pengurai.maju()

	return pengurai
}

func (p *Pengurai) Urai() (*NodeAkar, *TokenTakTerduga) {
	nodeNode := make([]interface{}, 0)

	for p.tokenSaatIni != nil {
		if p.tokenSaatIni.Jenis == lekser.T_ADF {
			break
		}
		node, err := p.deklarasi()
		if err != nil {
			return nil, err
		}

		nodeNode = append(nodeNode, node)

	}

	return &NodeAkar{
		BasisPosisi: utils.BasisPosisi{
			Baris: 0,
			Kolom: 0,
		},
		NodeNode: nodeNode,
	}, nil
}
