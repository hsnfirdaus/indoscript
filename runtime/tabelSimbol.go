package runtime

import (
	"indoscript/parser"
)

type ButirFungsi struct {
	namaArgument []string
	nodeAkar     *parser.NodeAkar
}

type TabelSimbol struct {
	variabel map[string]interface{}
	fungsi   map[string]ButirFungsi
	induk    *TabelSimbol
}

func (ts *TabelSimbol) Baru(induk *TabelSimbol) {
	ts.variabel = make(map[string]interface{})
	ts.fungsi = make(map[string]ButirFungsi)
	ts.induk = induk
}

func (ts *TabelSimbol) ambilVar(pengenal string) (interface{}, error) {
	val, ok := ts.variabel[pengenal]
	if !ok && ts.induk != nil {
		val, err := ts.induk.ambilVar(pengenal)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return val, nil
}

func (ts *TabelSimbol) aturVar(pengenal string, isi interface{}) {
	ts.variabel[pengenal] = isi
}

func (ts *TabelSimbol) ambilFung(pengenal string) (*ButirFungsi, error) {
	val, ok := ts.fungsi[pengenal]
	if !ok && ts.induk != nil {
		val, err := ts.induk.ambilFung(pengenal)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	return &val, nil
}

func (ts *TabelSimbol) aturFung(pengenal string, namaAgumen []string, nodeAkar *parser.NodeAkar) {
	ts.fungsi[pengenal] = ButirFungsi{
		namaArgument: namaAgumen,
		nodeAkar:     nodeAkar,
	}

}
