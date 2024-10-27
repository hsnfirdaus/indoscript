package pengurai

import (
	"indoscript/lekser"
	"slices"
)

func (p *Pengurai) operasi(fnKiri func() (interface{}, *TokenTakTerduga), fnKanan func() (interface{}, *TokenTakTerduga), opr []lekser.JenisToken) (interface{}, *TokenTakTerduga) {
	var kiri interface{}
	kiri, err := fnKiri()
	if err != nil {
		return nil, err
	}

	for p.tokenSaatIni != nil {
		if !slices.Contains(opr, p.tokenSaatIni.Jenis) {
			break
		}
		tokOp := p.tokenSaatIni.Jenis
		p.maju()
		kanan, err := fnKanan()
		if err != nil {
			return nil, err
		}

		kiri = &NodeOperasi{
			BasisPosisi: p.basisPosisi(),
			NodeKiri:    kiri,
			Operasi:     tokOp,
			NodeKanan:   kanan,
		}
	}

	return kiri, nil
}

func (p *Pengurai) operasiDanAtau(fnKiri func() (interface{}, *TokenTakTerduga), fnKanan func() (interface{}, *TokenTakTerduga)) (interface{}, *TokenTakTerduga) {
	var kiri interface{}
	kiri, err := fnKiri()
	if err != nil {
		return nil, err
	}

	for p.tokenSaatIni != nil {
		if p.tokenSaatIni.Jenis != lekser.T_KATKUN || !slices.Contains([]string{lekser.KK_ATAU, lekser.KK_DAN}, p.tokenSaatIni.Isi.(string)) {
			break
		}
		tokOp := p.tokenSaatIni.Isi.(string)
		p.maju()
		kanan, err := fnKanan()
		if err != nil {
			return nil, err
		}

		kiri = &NodeOperasiDanAtau{
			BasisPosisi: p.basisPosisi(),
			NodeKiri:    kiri,
			Operasi:     tokOp,
			NodeKanan:   kanan,
		}
	}

	return kiri, nil
}
