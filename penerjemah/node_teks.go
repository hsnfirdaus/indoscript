package penerjemah

import (
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeTeks(node *pengurai.NodeTeks) (*jenis.Teks, *KesalahanPenerjemah) {
	return &jenis.Teks{
		Teks: node.Teks,
	}, nil
}
