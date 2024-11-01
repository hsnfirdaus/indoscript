package pengurai

import "indoscript/lekser"

func (p *Pengurai) matBanding() (interface{}, *TokenTakTerduga) {
	if p.tokenSaatIni.Jenis == lekser.T_TIDAK {
		p.maju()

		hasil, err := p.matBanding()
		if err != nil {
			return nil, err
		}

		return &NodeTidak{
			Node: hasil,
		}, nil
	}
	return p.operasi(p.matExpr, p.matExpr, lekser.TOKEN_PERBANDINGAN)

}
