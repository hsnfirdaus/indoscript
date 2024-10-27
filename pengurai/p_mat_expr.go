package pengurai

import "indoscript/lekser"

func (p *Pengurai) matExpr() (interface{}, *TokenTakTerduga) {
	return p.operasi(p.term, p.term, []lekser.JenisToken{lekser.T_TAMBAH, lekser.T_KURANG})
}
