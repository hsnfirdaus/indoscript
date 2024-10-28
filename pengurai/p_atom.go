package pengurai

import "indoscript/lekser"

func (p *Pengurai) atom() (interface{}, *TokenTakTerduga) {
	tok := p.tokenSaatIni

	if tok == nil {
		return nil, nil
	}

	if p.tokenSaatIni.Jenis == lekser.T_TEKS {
		p.maju()
		return &NodeTeks{
			BasisPosisi: tok.BasisPosisi,
			Teks:        tok.Isi.(string),
		}, nil
	} else if tok.Jenis == lekser.T_KATKUN {
		switch tok.Isi {
		case lekser.KK_BENAR:
			p.maju()
			return &NodeBoolean{
				BasisPosisi: p.basisPosisi(),
				Isi:         true,
			}, nil

		case lekser.KK_SALAH:
			p.maju()
			return &NodeBoolean{
				BasisPosisi: p.basisPosisi(),
				Isi:         false,
			}, nil
		}
	} else if tok.Jenis == lekser.T_PENGENAL {
		p.maju()

		if p.tokenSaatIni.Jenis == lekser.T_BKURUNG {
			p.maju()
			daftarArgument := make([]interface{}, 0)

			adaVar := false
			if p.tokenSaatIni.Jenis != lekser.T_TKURUNG {
				adaVar = true
			}
			for adaVar {
				argument, err := p.expr()
				if err != nil {
					return nil, err
				}
				daftarArgument = append(daftarArgument, argument)

				if p.tokenSaatIni.Jenis == lekser.T_KOMA {
					p.maju()
				} else {
					adaVar = false
					break
				}
			}
			if p.tokenSaatIni.Jenis != lekser.T_TKURUNG {
				return nil, &TokenTakTerduga{
					BasisPosisi: p.basisPosisi(),
					diharapkan:  []lekser.JenisToken{lekser.T_TKURUNG},
					ditemukan:   p.tokenSaatIni.Jenis,
				}
			}
			p.maju()

			return &NodePanggilFungsi{
				BasisPosisi: p.basisPosisi(),
				NamaFungsi:  tok.Isi.(string),
				Argumen:     daftarArgument,
			}, nil
		}

		return &NodeAksesVariabel{
			BasisPosisi:  p.basisPosisi(),
			NamaVariabel: tok.Isi.(string),
		}, nil
	} else if tok.Jenis == lekser.T_BUL || tok.Jenis == lekser.T_DES {
		p.maju()
		return &NodeBilangan{
			BasisPosisi: p.basisPosisi(),
			Token:       *tok,
		}, nil
	} else if tok.Jenis == lekser.T_BKURUNG {
		p.maju()
		expr, err := p.expr()
		if err != nil {
			return nil, err
		}
		if p.tokenSaatIni.Jenis == lekser.T_TKURUNG {
			p.maju()
			return expr, nil
		}

		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_BKURUNG},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}

	return nil, &TokenTakTerduga{
		BasisPosisi: p.basisPosisi(),
		diharapkan:  []lekser.JenisToken{lekser.T_BUL, lekser.T_DES},
		ditemukan:   p.tokenSaatIni.Jenis,
	}
}
