package parser

import (
	"indoscript/lexer"
	"indoscript/utils"
	"slices"
)

type Parser struct {
	TokenToken   []lexer.Token
	TokenSaatIni *lexer.Token
	Indeks       int
}

func Baru(tokenToken []lexer.Token) Parser {
	newPar := Parser{
		TokenToken:   tokenToken,
		TokenSaatIni: nil,
		Indeks:       -1,
	}
	newPar.maju()

	return newPar
}

func (par *Parser) ambilBasePosisi() utils.BasePosisi {
	return par.TokenSaatIni.BasePosisi
}

func (par *Parser) maju() {
	par.Indeks = par.Indeks + 1
	if par.Indeks < len(par.TokenToken) {
		par.TokenSaatIni = &par.TokenToken[par.Indeks]
	} else {
		par.TokenSaatIni = nil
	}
}

func (par *Parser) ambilTokenSelanjutnya() *lexer.Token {
	indeks := par.Indeks + 1
	if indeks < len(par.TokenToken) {
		return &par.TokenToken[indeks]
	}

	return nil
}

func (par *Parser) TKMaju() *TokenTakTerduga {

	if par.TokenSaatIni == nil {
		return &TokenTakTerduga{}
	} else if par.TokenSaatIni.Jenis != lexer.T_TK {
		return &TokenTakTerduga{
			BasePosisi: utils.BasePosisi{},
		}
	}

	par.maju()

	return nil
}

func (par *Parser) Parse() (*NodeAkar, *TokenTakTerduga) {
	nodeNode := make([]interface{}, 0)

	for par.TokenSaatIni != nil {
		if par.TokenSaatIni.Jenis == lexer.T_ADF {
			break
		}
		node, err := par.deklarasi()
		if err != nil {
			return nil, err
		}

		nodeNode = append(nodeNode, node)

	}

	return &NodeAkar{nodeNode}, nil
}

func (par *Parser) akarKurawal() (*NodeAkar, *TokenTakTerduga) {
	par.maju()
	nodeNode := make([]interface{}, 0)

	for par.TokenSaatIni != nil {
		if par.TokenSaatIni.Jenis == lexer.T_TKURAWAL {
			par.maju()
			break
		}
		node, err := par.deklarasi()
		if err != nil {
			return nil, err
		}

		nodeNode = append(nodeNode, node)

	}

	return &NodeAkar{nodeNode}, nil
}

func (par *Parser) operasi(fnKiri func() (interface{}, *TokenTakTerduga), fnKanan func() (interface{}, *TokenTakTerduga), opr []lexer.JenisToken) (interface{}, *TokenTakTerduga) {
	var kiri interface{}
	kiri, err := fnKiri()
	if err != nil {
		return nil, err
	}

	for par.TokenSaatIni != nil {
		if !slices.Contains(opr, par.TokenSaatIni.Jenis) {
			break
		}
		tokOp := par.TokenSaatIni.Jenis
		par.maju()
		kanan, err := fnKanan()
		if err != nil {
			return nil, err
		}

		kiri = &NodeOperasi{
			BasePosisi: par.ambilBasePosisi(),
			NodeKiri:   kiri,
			Operasi:    tokOp,
			NodeKanan:  kanan,
		}
	}

	return kiri, nil
}

func (par *Parser) deklarasi() (interface{}, *TokenTakTerduga) {
	tok := par.TokenSaatIni
	if par.TokenSaatIni.Jenis == lexer.T_KATKUN {
		if tok.Isi == lexer.KK_VAR {
			par.maju()
			if par.TokenSaatIni.Jenis != lexer.T_PENGENAL {
				return nil, &TokenTakTerduga{
					BasePosisi: par.ambilBasePosisi(),
					diharapkan: []lexer.JenisToken{lexer.T_PENGENAL},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}

			namaVariabel := par.TokenSaatIni.Isi

			par.maju()

			if par.TokenSaatIni.Jenis != lexer.T_SD {
				return nil, &TokenTakTerduga{
					BasePosisi: par.ambilBasePosisi(),
					diharapkan: []lexer.JenisToken{lexer.T_SD},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}

			par.maju()

			expr, err := par.expr()
			if err != nil {
				return nil, err
			}

			par.TKMaju()

			return &NodeAturVariabel{
				BasePosisi: utils.BasePosisi{
					Baris: tok.Baris,
					Kolom: tok.Kolom,
				},
				NamaVariabel: namaVariabel.(string),
				Node:         expr,
			}, nil

		}

		if tok.Isi == lexer.KK_JIKA {
			par.maju()
			if par.TokenSaatIni.Jenis != lexer.T_BKURUNG {
				return nil, &TokenTakTerduga{
					BasePosisi: utils.BasePosisi{},
					diharapkan: []lexer.JenisToken{lexer.T_BKURUNG},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}
			par.maju()

			compExpr, err := par.expr()
			if err != nil {
				return nil, err
			}

			if par.TokenSaatIni.Jenis != lexer.T_TKURUNG {
				return nil, &TokenTakTerduga{
					BasePosisi: utils.BasePosisi{},
					diharapkan: []lexer.JenisToken{lexer.T_TKURUNG},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}
			par.maju()

			if par.TokenSaatIni.Jenis != lexer.T_BKURAWAL {
				return nil, &TokenTakTerduga{
					BasePosisi: utils.BasePosisi{},
					diharapkan: []lexer.JenisToken{lexer.T_BKURAWAL},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}

			nodeNode, err := par.akarKurawal()
			if err != nil {
				return nil, err
			}

			return &NodeJika{
				Kondisi:  compExpr,
				NodeNode: nodeNode,
			}, nil
		}

		if tok.Isi == lexer.KK_FUNGSI {
			par.maju()
			if par.TokenSaatIni.Jenis != lexer.T_PENGENAL {
				return nil, &TokenTakTerduga{
					BasePosisi: utils.BasePosisi{},
					diharapkan: []lexer.JenisToken{lexer.T_PENGENAL},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}
			namaFungsi := par.TokenSaatIni.Isi.(string)
			par.maju()
			if par.TokenSaatIni.Jenis != lexer.T_BKURUNG {
				return nil, &TokenTakTerduga{
					BasePosisi: utils.BasePosisi{},
					diharapkan: []lexer.JenisToken{lexer.T_BKURUNG},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}
			par.maju()

			daftarNamaArgument := make([]string, 0)
			punyaArgumen := false
			if par.TokenSaatIni.Jenis != lexer.T_TKURUNG {
				punyaArgumen = true
			}

			for punyaArgumen {
				if par.TokenSaatIni.Jenis != lexer.T_PENGENAL {
					return nil, &TokenTakTerduga{
						BasePosisi: utils.BasePosisi{},
						diharapkan: []lexer.JenisToken{lexer.T_PENGENAL},
						ditemukan:  par.TokenSaatIni.Jenis,
					}
				}
				pengenal := par.TokenSaatIni.Isi.(string)
				daftarNamaArgument = append(daftarNamaArgument, pengenal)
				par.maju()

				if par.TokenSaatIni.Jenis == lexer.T_TKURUNG {
					break
				}
			}

			par.maju()
			if par.TokenSaatIni.Jenis != lexer.T_BKURAWAL {
				return nil, &TokenTakTerduga{
					BasePosisi: utils.BasePosisi{},
					diharapkan: []lexer.JenisToken{lexer.T_BKURAWAL},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}
			nodeNode, err := par.akarKurawal()
			if err != nil {
				return nil, err
			}

			return &NodeAturFungsi{
				NamaFungsi:  namaFungsi,
				NamaArgumen: daftarNamaArgument,
				NodeNode:    nodeNode,
			}, nil
		}

		if tok.Isi == lexer.KK_BALIKAN {
			par.maju()

			expr, err := par.expr()
			if err != nil {
				return nil, err
			}

			par.TKMaju()

			return &NodeBalikan{
				BasePosisi: utils.BasePosisi{
					Baris: tok.Baris,
					Kolom: tok.Kolom,
				},
				Node: expr,
			}, nil
		}
	}

	if tok.Jenis == lexer.T_PENGENAL {
		return par.expr()
	}

	return nil, &TokenTakTerduga{
		BasePosisi: par.ambilBasePosisi(),
		diharapkan: []lexer.JenisToken{lexer.T_KATKUN},
		ditemukan:  par.TokenSaatIni.Jenis,
	}
}

func (par *Parser) expr() (interface{}, *TokenTakTerduga) {
	tok := par.TokenSaatIni

	if par.TokenSaatIni.Jenis == lexer.T_TEKS {
		par.maju()
		return &NodeTeks{
			BasePosisi: tok.BasePosisi,
			Teks:       tok.Isi.(string),
		}, nil
	}

	if slices.Contains([]lexer.JenisToken{lexer.T_BUL, lexer.T_DES, lexer.T_TAMBAH, lexer.T_KURANG, lexer.T_PENGENAL}, par.TokenSaatIni.Jenis) {
		hasil, err := par.matBanding()
		if err != nil {
			return nil, err
		}

		if par.TokenSaatIni.Jenis == lexer.T_TK {
			par.maju()
		}

		return hasil, nil
	}

	return nil, nil
}

func (par *Parser) matBanding() (interface{}, *TokenTakTerduga) {
	if par.TokenSaatIni.Jenis == lexer.T_TIDAK {
		par.maju()

		hasil, err := par.matBanding()
		if err != nil {
			return nil, err
		}

		return &NodeTidak{
			Node: hasil,
		}, nil
	}
	return par.operasi(par.matExpr, par.matExpr, []lexer.JenisToken{lexer.T_KDAR, lexer.T_LDAR, lexer.T_KDARSD, lexer.T_LDARSD, lexer.T_SDSD})

}

func (par *Parser) matExpr() (interface{}, *TokenTakTerduga) {
	return par.operasi(par.term, par.term, []lexer.JenisToken{lexer.T_TAMBAH, lexer.T_KURANG})
}

func (par *Parser) term() (interface{}, *TokenTakTerduga) {
	return par.operasi(par.faktor, par.faktor, []lexer.JenisToken{lexer.T_KALI, lexer.T_BAGI})
}

func (par *Parser) faktor() (interface{}, *TokenTakTerduga) {
	tok := par.TokenSaatIni

	if tok == nil {
		return nil, nil
	}

	if tok.Jenis == lexer.T_TAMBAH || tok.Jenis == lexer.T_KURANG {
		par.maju()
		bil, err := par.pangkat()
		if err != nil {
			return nil, err
		}
		return &NodeOperasiUner{
			BasePosisi: par.ambilBasePosisi(),
			Token:      tok.Jenis,
			Node:       bil,
		}, nil
	}
	return par.pangkat()
}

func (par *Parser) pangkat() (interface{}, *TokenTakTerduga) {
	return par.operasi(par.atom, par.faktor, []lexer.JenisToken{lexer.T_PANGKAT})
}

func (par *Parser) atom() (interface{}, *TokenTakTerduga) {
	tok := par.TokenSaatIni

	if tok == nil {
		return nil, nil
	}

	if tok.Jenis == lexer.T_PENGENAL {
		par.maju()

		if par.TokenSaatIni.Jenis == lexer.T_BKURUNG {
			par.maju()
			daftarArgument := make([]interface{}, 0)

			adaVar := false
			if par.TokenSaatIni.Jenis != lexer.T_TKURUNG {
				adaVar = true
			}
			for adaVar {
				argument, err := par.expr()
				if err != nil {
					return nil, err
				}
				daftarArgument = append(daftarArgument, argument)

				if par.TokenSaatIni.Jenis == lexer.T_KOMA {
					par.maju()
				} else {
					adaVar = false
					break
				}
			}
			if par.TokenSaatIni.Jenis != lexer.T_TKURUNG {
				return nil, &TokenTakTerduga{
					BasePosisi: par.ambilBasePosisi(),
					diharapkan: []lexer.JenisToken{lexer.T_TKURUNG},
					ditemukan:  par.TokenSaatIni.Jenis,
				}
			}
			par.maju()

			return &NodePanggilFungsi{
				NamaFungsi: tok.Isi.(string),
				Argumen:    daftarArgument,
			}, nil
		}

		return &NodeAksesVariabel{
			BasePosisi:   par.ambilBasePosisi(),
			NamaVariabel: tok.Isi.(string),
		}, nil
	} else if tok.Jenis == lexer.T_BUL || tok.Jenis == lexer.T_DES {
		par.maju()
		return &NodeBilangan{
			BasePosisi: par.ambilBasePosisi(),
			Token:      *tok,
		}, nil
	} else if tok.Jenis == lexer.T_BKURUNG {
		par.maju()
		expr, err := par.expr()
		if err != nil {
			return nil, err
		}
		if par.TokenSaatIni.Jenis == lexer.T_TKURUNG {
			par.maju()
			return expr, nil
		}

		return nil, &TokenTakTerduga{
			BasePosisi: par.ambilBasePosisi(),
			diharapkan: []lexer.JenisToken{lexer.T_BKURUNG},
			ditemukan:  par.TokenSaatIni.Jenis,
		}
	}

	return nil, &TokenTakTerduga{
		BasePosisi: par.ambilBasePosisi(),
		diharapkan: []lexer.JenisToken{lexer.T_BUL, lexer.T_DES},
		ditemukan:  par.TokenSaatIni.Jenis,
	}
}
