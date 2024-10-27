package penerjemah

import (
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeBalikan(node *pengurai.NodeBalikan) (*jenis.Balikan, *KesalahanPenerjemah) {
	hasil, err := p.panggilNode(node.Node)
	if err != nil {
		return nil, err
	}

	return &jenis.Balikan{
		Isi: hasil,
	}, nil
}
