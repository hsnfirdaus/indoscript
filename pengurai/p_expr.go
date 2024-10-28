package pengurai

func (p *Pengurai) expr() (interface{}, *TokenTakTerduga) {

	return p.operasiDanAtau(p.matBanding, p.matBanding)

}
