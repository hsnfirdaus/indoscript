package penerjemah

import (
	"indoscript/pengurai"
	"indoscript/utils"
)

func (p *Penerjemah) NodeJika(node *pengurai.NodeJika) (interface{}, *KesalahanPenerjemah) {
	banding, err := p.panggilNodeBoolean(node.Kondisi)
	if err != nil {
		return nil, &KesalahanPenerjemah{
			BasisPosisi: utils.BasisPosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: err.Error(),
		}
	}

	if banding.Isi {
		hasil, err := p.panggilNode(node.NodeNode)
		if err != nil {
			return nil, err
		}

		return hasil, nil
	}

	return nil, nil
}
