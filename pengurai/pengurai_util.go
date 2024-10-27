package pengurai

import (
	"indoscript/lekser"
	"indoscript/utils"
)

func (p *Pengurai) basisPosisi() utils.BasisPosisi {
	return p.tokenSaatIni.BasisPosisi
}

func (p *Pengurai) maju() {
	p.indeks = p.indeks + 1
	if p.indeks < len(p.daftarToken) {
		p.tokenSaatIni = &p.daftarToken[p.indeks]
	} else {
		p.tokenSaatIni = nil
	}
}

func (p *Pengurai) tkMaju() *TokenTakTerduga {

	if p.tokenSaatIni == nil {
		return &TokenTakTerduga{}
	} else if p.tokenSaatIni.Jenis != lekser.T_TK {
		return &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_TK},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}

	p.maju()

	return nil
}
