package pengurai

import "indoscript/lekser"

func (p *Pengurai) pangkat() (interface{}, *TokenTakTerduga) {
	return p.operasi(p.atom, p.faktor, []lekser.JenisToken{lekser.T_PANGKAT})
}
