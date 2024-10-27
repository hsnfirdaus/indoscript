package pengurai

import (
	"indoscript/lekser"
	"indoscript/utils"
)

func (p *Pengurai) deklarasi() (interface{}, *TokenTakTerduga) {
	tok := p.tokenSaatIni
	if p.tokenSaatIni.Jenis == lekser.T_KATKUN {
		if tok.Isi == lekser.KK_VAR {
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
				BasisPosisi: utils.BasisPosisi{
					Baris: tok.Baris,
					Kolom: tok.Kolom,
				},
				NamaVariabel: namaVariabel.(string),
				Node:         expr,
			}, nil

		}

		if tok.Isi == lekser.KK_JIKA {
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

			return &NodeJika{
				Kondisi:  compExpr,
				NodeNode: nodeNode,
			}, nil
		}

		if tok.Isi == lekser.KK_FUNGSI {
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
				NamaFungsi:  namaFungsi,
				NamaArgumen: daftarNamaArgument,
				NodeNode:    nodeNode,
			}, nil
		}

		if tok.Isi == lekser.KK_BALIKAN {
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
				BasisPosisi: utils.BasisPosisi{
					Baris: tok.Baris,
					Kolom: tok.Kolom,
				},
				Node: expr,
			}, nil
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
