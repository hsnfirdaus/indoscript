package pengurai

import "indoscript/lekser"

func (p *Pengurai) term() (interface{}, *TokenTakTerduga) {
	return p.operasi(p.faktor, p.faktor, []lekser.JenisToken{lekser.T_KALI, lekser.T_BAGI})
}
