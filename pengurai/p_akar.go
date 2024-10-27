package pengurai

import "indoscript/lekser"

func (p *Pengurai) uraiKurawal() (*NodeAkar, *TokenTakTerduga) {
	posisi := p.basisPosisi()
	p.maju()
	nodeNode := make([]interface{}, 0)

	for p.tokenSaatIni != nil {
		if p.tokenSaatIni.Jenis == lekser.T_TKURAWAL {
			p.maju()
			break
		}
		node, err := p.deklarasi()
		if err != nil {
			return nil, err
		}

		nodeNode = append(nodeNode, node)

	}

	return &NodeAkar{
		BasisPosisi: posisi,
		NodeNode:    nodeNode,
	}, nil
}
