package pengurai

import "indoscript/lekser"

func (p *Pengurai) faktor() (interface{}, *TokenTakTerduga) {
	tok := p.tokenSaatIni

	if tok == nil {
		return nil, nil
	}

	if tok.Jenis == lekser.T_TAMBAH || tok.Jenis == lekser.T_KURANG {
		p.maju()
		bil, err := p.pangkat()
		if err != nil {
			return nil, err
		}
		return &NodeOperasiUner{
			BasisPosisi: p.basisPosisi(),
			Token:       tok.Jenis,
			Node:        bil,
		}, nil
	}
	return p.pangkat()
}
