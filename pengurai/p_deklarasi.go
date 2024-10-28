package pengurai

import (
	"indoscript/lekser"
)

func (p *Pengurai) deklarasi() (interface{}, *TokenTakTerduga) {
	tok := p.tokenSaatIni
	if p.tokenSaatIni.Jenis == lekser.T_KATKUN {
		switch tok.Isi {
		case lekser.KK_VAR:
			return p.deklarasiVar()
		case lekser.KK_JIKA:
			return p.deklarasiJika()
		case lekser.KK_FUNGSI:
			return p.deklarasiFungsi()
		case lekser.KK_BALIKAN:
			return p.deklarasiBalikan()

		}
	}

	if tok.Jenis == lekser.T_PENGENAL {
		hasil, err := p.expr()
		if err != nil {
			return nil, err
		}
		err = p.tkMaju()
		if err != nil {
			return nil, err
		}
		return hasil, nil
	}

	return nil, &TokenTakTerduga{
		BasisPosisi: p.basisPosisi(),
		diharapkan:  []lekser.JenisToken{lekser.T_KATKUN},
		ditemukan:   p.tokenSaatIni.Jenis,
	}
}

func (p *Pengurai) deklarasiVar() (interface{}, *TokenTakTerduga) {
	posisi := p.basisPosisi()

	p.maju()
	if p.tokenSaatIni.Jenis != lekser.T_PENGENAL {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_PENGENAL},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}

	namaVariabel := p.tokenSaatIni.Isi

	p.maju()

	if p.tokenSaatIni.Jenis != lekser.T_SD {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_SD},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}

	p.maju()

	expr, err := p.expr()
	if err != nil {
		return nil, err
	}

	err = p.tkMaju()
	if err != nil {
		return nil, err
	}

	return &NodeAturVariabel{
		BasisPosisi:  posisi,
		NamaVariabel: namaVariabel.(string),
		Node:         expr,
	}, nil
}

func (p *Pengurai) deklarasiJika() (interface{}, *TokenTakTerduga) {
	posisi := p.basisPosisi()

	p.maju()
	if p.tokenSaatIni.Jenis != lekser.T_BKURUNG {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_BKURUNG},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}
	p.maju()

	compExpr, err := p.expr()
	if err != nil {
		return nil, err
	}

	if p.tokenSaatIni.Jenis != lekser.T_TKURUNG {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_TKURUNG},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}
	p.maju()

	if p.tokenSaatIni.Jenis != lekser.T_BKURAWAL {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_BKURAWAL},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}

	nodeNode, err := p.uraiKurawal()
	if err != nil {
		return nil, err
	}

	var daftarKasusLain []kasusLain = make([]kasusLain, 0)
	var lainLain *NodeAkar = nil

	for p.tokenSaatIni.Jenis == lekser.T_KATKUN && p.tokenSaatIni.Isi.(string) == lekser.KK_LAIN {
		p.maju()

		if p.tokenSaatIni.Jenis == lekser.T_KATKUN && p.tokenSaatIni.Isi.(string) == lekser.KK_JIKA {
			p.maju()
			if p.tokenSaatIni.Jenis != lekser.T_BKURUNG {
				return nil, &TokenTakTerduga{
					BasisPosisi: p.basisPosisi(),
					diharapkan:  []lekser.JenisToken{lekser.T_BKURUNG},
					ditemukan:   p.tokenSaatIni.Jenis,
				}
			}
			p.maju()

			lainJikaExpr, err := p.expr()
			if err != nil {
				return nil, err
			}

			if p.tokenSaatIni.Jenis != lekser.T_TKURUNG {
				return nil, &TokenTakTerduga{
					BasisPosisi: p.basisPosisi(),
					diharapkan:  []lekser.JenisToken{lekser.T_TKURUNG},
					ditemukan:   p.tokenSaatIni.Jenis,
				}
			}
			p.maju()

			if p.tokenSaatIni.Jenis != lekser.T_BKURAWAL {
				return nil, &TokenTakTerduga{
					BasisPosisi: p.basisPosisi(),
					diharapkan:  []lekser.JenisToken{lekser.T_BKURAWAL},
					ditemukan:   p.tokenSaatIni.Jenis,
				}
			}

			lainJikaNode, err := p.uraiKurawal()
			if err != nil {
				return nil, err
			}

			daftarKasusLain = append(daftarKasusLain, kasusLain{
				Kondisi:  lainJikaExpr,
				NodeNode: lainJikaNode,
			})
		} else {
			if p.tokenSaatIni.Jenis != lekser.T_BKURAWAL {
				return nil, &TokenTakTerduga{
					BasisPosisi: p.basisPosisi(),
					diharapkan:  []lekser.JenisToken{lekser.T_BKURAWAL},
					ditemukan:   p.tokenSaatIni.Jenis,
				}
			}

			lainLainNode, err := p.uraiKurawal()
			if err != nil {
				return nil, err
			}
			lainLain = lainLainNode
			break
		}
	}

	return &NodeJika{
		BasisPosisi: posisi,
		Kondisi:     compExpr,
		NodeNode:    nodeNode,
		KasusLain:   daftarKasusLain,
		LainLain:    lainLain,
	}, nil
}

func (p *Pengurai) deklarasiFungsi() (interface{}, *TokenTakTerduga) {
	posisi := p.basisPosisi()

	p.maju()
	if p.tokenSaatIni.Jenis != lekser.T_PENGENAL {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_PENGENAL},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}
	namaFungsi := p.tokenSaatIni.Isi.(string)
	p.maju()
	if p.tokenSaatIni.Jenis != lekser.T_BKURUNG {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_BKURUNG},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}
	p.maju()

	daftarNamaArgument := make([]string, 0)
	punyaArgumen := false
	if p.tokenSaatIni.Jenis != lekser.T_TKURUNG {
		punyaArgumen = true
	}

	for punyaArgumen {
		if p.tokenSaatIni.Jenis != lekser.T_PENGENAL {
			return nil, &TokenTakTerduga{
				BasisPosisi: p.basisPosisi(),
				diharapkan:  []lekser.JenisToken{lekser.T_PENGENAL},
				ditemukan:   p.tokenSaatIni.Jenis,
			}
		}
		pengenal := p.tokenSaatIni.Isi.(string)
		daftarNamaArgument = append(daftarNamaArgument, pengenal)
		p.maju()

		if p.tokenSaatIni.Jenis == lekser.T_TKURUNG {
			break
		} else if p.tokenSaatIni.Jenis == lekser.T_KOMA {
			p.maju()
		} else {
			return nil, &TokenTakTerduga{
				BasisPosisi: p.basisPosisi(),
				diharapkan:  []lekser.JenisToken{lekser.T_KOMA},
				ditemukan:   p.tokenSaatIni.Jenis,
			}
		}
	}

	p.maju()
	if p.tokenSaatIni.Jenis != lekser.T_BKURAWAL {
		return nil, &TokenTakTerduga{
			BasisPosisi: p.basisPosisi(),
			diharapkan:  []lekser.JenisToken{lekser.T_BKURAWAL},
			ditemukan:   p.tokenSaatIni.Jenis,
		}
	}
	nodeNode, err := p.uraiKurawal()
	if err != nil {
		return nil, err
	}

	return &NodeAturFungsi{
		BasisPosisi: posisi,
		NamaFungsi:  namaFungsi,
		NamaArgumen: daftarNamaArgument,
		NodeNode:    nodeNode,
	}, nil
}

func (p *Pengurai) deklarasiBalikan() (interface{}, *TokenTakTerduga) {
	posisi := p.basisPosisi()

	p.maju()

	expr, err := p.expr()
	if err != nil {
		return nil, err
	}

	err = p.tkMaju()
	if err != nil {
		return nil, err
	}

	return &NodeBalikan{
		BasisPosisi: posisi,
		Node:        expr,
	}, nil
}
