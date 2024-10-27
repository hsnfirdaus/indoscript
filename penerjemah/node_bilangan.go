package penerjemah

import (
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
)

func (p *Penerjemah) nodeBilangan(node *pengurai.NodeBilangan) (*jenis.Bilangan, *KesalahanPenerjemah) {
	isi := node.Token.Isi.(float64)
	return &jenis.Bilangan{Angka: isi}, nil
}
