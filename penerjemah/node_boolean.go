package penerjemah

import (
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeBoolean(node *pengurai.NodeBoolean) (*jenis.Boolean, *KesalahanPenerjemah) {
	return &jenis.Boolean{Isi: node.Isi}, nil
}
