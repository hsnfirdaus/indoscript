package pengurai

import (
	"indoscript/lekser"
)

func (p *Pengurai) expr() (interface{}, *TokenTakTerduga) {
	tok := p.tokenSaatIni

	if p.tokenSaatIni.Jenis == lekser.T_TEKS {
		p.maju()
		return &NodeTeks{
			BasisPosisi: tok.BasisPosisi,
			Teks:        tok.Isi.(string),
		}, nil
	}

	return p.operasiDanAtau(p.matBanding, p.matBanding)

}
