package penerjemah

import (
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeAturVariabel(node *pengurai.NodeAturVariabel) (interface{}, *KesalahanPenerjemah) {
	isi, err := p.panggilNode(node.Node)
	if err != nil {
		return nil, err
	}

	p.ts.aturVar(node.NamaVariabel, isi)

	return isi, nil
}

func (p *Penerjemah) nodeAksesVariabel(node *pengurai.NodeAksesVariabel) (interface{}, *KesalahanPenerjemah) {
	isi, err := p.ts.ambilVar(node.NamaVariabel)
	if err != nil {
		return nil, &KesalahanPenerjemah{
			BasisPosisi: node.BasisPosisi,
			pesan:       "Variabel tidak didefinisikan: " + node.NamaVariabel,
		}
	}

	return isi, nil
}
