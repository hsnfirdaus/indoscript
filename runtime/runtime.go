package runtime

import (
	"fmt"
	"indoscript/lexer"
	"indoscript/parser"
	"indoscript/runtime/fungsi"
	"indoscript/runtime/jenis"
	"indoscript/utils"
	"reflect"
	"slices"
)

type Runtime struct {
	ts *TabelSimbol
}

func (r *Runtime) Baru() {
	ts := TabelSimbol{}
	ts.Baru(nil)
	r.ts = &ts
}

func (r *Runtime) anakRuntime() Runtime {
	ts := TabelSimbol{}
	ts.Baru(r.ts)
	rt := Runtime{
		ts: &ts,
	}

	return rt
}

func (r *Runtime) Jalankan(node *parser.NodeAkar) *KesalahanRuntime {
	_, err := r.panggilNode(node)
	return err
}

func (r *Runtime) panggilNode(node interface{}) (interface{}, *KesalahanRuntime) {

	switch v := node.(type) {
	case *parser.NodeAkar:
		return r.nodeAkar(v)
	case *parser.NodeBilangan:
		return r.nodeBilangan(v)

	case *parser.NodeOperasi:
		return r.nodeOperasi(v)

	case *parser.NodeOperasiUner:
		return r.nodeOperasiUner(v)

	case *parser.NodeAturVariabel:
		return r.nodeAturVariabel(v)

	case *parser.NodeAksesVariabel:
		return r.nodeAksesVariabel(v)

	case *parser.NodeAturFungsi:
		return r.nodeAturFungsi(v)

	case *parser.NodePanggilFungsi:
		return r.nodePanggilFungsi(v)

	case *parser.NodeTeks:
		return r.nodeTeks(v)

	case *parser.NodeJika:
		return r.NodeJika(v)

	case *parser.NodeBalikan:
		return r.nodeBalikan(v)

	default:
		jenisNode := reflect.TypeOf(v)
		return nil, &KesalahanRuntime{
			pesan: "Node \"" + jenisNode.String() + "\" tak dikenali!",
		}
	}
}

func (r *Runtime) panggilNodeBilangan(node interface{}) (*jenis.Bilangan, *KesalahanRuntime) {
	isi, err := r.panggilNode(node)
	if err != nil {
		return nil, err
	}

	isiBil, ok := isi.(*jenis.Bilangan)
	if !ok {
		return nil, &KesalahanRuntime{
			// TODO: How to get baris kolom without assert one by one
			pesan: "Bukan merupakan bilangan!",
		}
	}

	return isiBil, nil
}

func (r *Runtime) panggilNodeBoolean(node interface{}) (*jenis.Boolean, *KesalahanRuntime) {
	isi, err := r.panggilNode(node)
	if err != nil {
		return nil, err
	}

	isiBool, ok := isi.(*jenis.Boolean)
	if !ok {
		jenisAsli := reflect.TypeOf(isi)
		return nil, &KesalahanRuntime{
			// TODO: How to get baris kolom without assert one by one
			pesan: jenisAsli.String() + " bukan merupakan boolean!",
		}
	}

	return isiBool, nil
}

func (r *Runtime) nodeAkar(node *parser.NodeAkar) (interface{}, *KesalahanRuntime) {
	rt := r.anakRuntime()
	results := make([]interface{}, 0)

	for _, node := range node.NodeNode {
		balikan, apakahBalikan := node.(*parser.NodeBalikan)
		if apakahBalikan {
			return rt.nodeBalikan(balikan)
		}
		hasil, err := rt.panggilNode(node)
		if err != nil {
			return nil, err
		}
		bal, apBal := hasil.(*jenis.Balikan)
		if apBal {
			return bal, nil
		}
		results = append(results, hasil)
	}

	return results, nil
}

func (r *Runtime) nodeBilangan(node *parser.NodeBilangan) (*jenis.Bilangan, *KesalahanRuntime) {
	isi := node.Token.Isi.(float64)
	return &jenis.Bilangan{Angka: isi}, nil
}
func (r *Runtime) nodeOperasi(node *parser.NodeOperasi) (interface{}, *KesalahanRuntime) {
	kiri, err := r.panggilNodeBilangan(node.NodeKiri)
	if err != nil {
		return nil, err
	}
	kanan, err := r.panggilNodeBilangan(node.NodeKanan)
	if err != nil {
		return nil, err
	}

	if slices.Contains(lexer.TOKEN_PERBANDINGAN, node.Operasi) {
		res, er := kiri.OperasiBoolean(kanan, node.Operasi)
		if er != nil {
			return nil, &KesalahanRuntime{
				BasePosisi: utils.BasePosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: er.Error(),
			}
		}

		return res, nil
	}

	er := kiri.Operasi(kanan, node.Operasi)
	if er != nil {
		return nil, &KesalahanRuntime{
			BasePosisi: utils.BasePosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: er.Error(),
		}
	}

	return kiri, nil
}
func (r *Runtime) nodeOperasiUner(node *parser.NodeOperasiUner) (*jenis.Bilangan, *KesalahanRuntime) {
	angka, err := r.panggilNodeBilangan(node.Node)
	if err != nil {
		return nil, err
	}

	if node.Token == lexer.T_KURANG {
		err := angka.Operasi(&jenis.Bilangan{Angka: -1}, lexer.T_KALI)
		if err != nil {
			return nil, &KesalahanRuntime{
				BasePosisi: utils.BasePosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: err.Error(),
			}
		}
	}

	return angka, nil
}

func (r *Runtime) nodeAturVariabel(node *parser.NodeAturVariabel) (interface{}, *KesalahanRuntime) {
	isi, err := r.panggilNode(node.Node)
	if err != nil {
		return nil, err
	}

	r.ts.aturVar(node.NamaVariabel, isi)

	return isi, nil
}

func (r *Runtime) nodeAksesVariabel(node *parser.NodeAksesVariabel) (interface{}, *KesalahanRuntime) {
	isi, err := r.ts.ambilVar(node.NamaVariabel)
	if err != nil {
		return nil, &KesalahanRuntime{
			BasePosisi: utils.BasePosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: "Variabel tidak didefinisikan: " + node.NamaVariabel,
		}
	}

	return isi, nil
}

func (r *Runtime) nodeAturFungsi(node *parser.NodeAturFungsi) (interface{}, *KesalahanRuntime) {
	r.ts.aturFung(node.NamaFungsi, node.NamaArgumen, node.NodeNode)

	return nil, nil
}

func (r *Runtime) nodePanggilFungsi(node *parser.NodePanggilFungsi) (interface{}, *KesalahanRuntime) {
	var hasilArgumen = make([]interface{}, 0)

	for _, arg := range node.Argumen {
		hasil, err := r.panggilNode(arg)
		if err != nil {
			return nil, err
		}
		hasilArgumen = append(hasilArgumen, hasil)
	}

	var hasilPanggil interface{}
	var err error = nil

	switch node.NamaFungsi {
	case "cetak":
		hasilPanggil, err = fungsi.Cetak(hasilArgumen)

	case "cetakBr":
		hasilPanggil, err = fungsi.CetakBr(hasilArgumen)

	default:
		fung, _ := r.ts.ambilFung(node.NamaFungsi)
		if fung == nil {
			return nil, &KesalahanRuntime{
				BasePosisi: utils.BasePosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: "Tidak ditemukan fungsi dengan nama \"" + node.NamaFungsi + "\"",
			}
		}

		hasil, err := r.lakukanPanggilFungsiTS(fung, hasilArgumen)
		if err != nil {
			return nil, &KesalahanRuntime{
				BasePosisi: utils.BasePosisi{
					Baris: node.Baris,
					Kolom: node.Kolom,
				},
				pesan: "Fungsi \"" + node.NamaFungsi + "\": " + err.Error(),
			}
		}
		return hasil, nil

	}

	if err != nil {
		return nil, &KesalahanRuntime{
			BasePosisi: utils.BasePosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: err.Error(),
		}
	}

	return hasilPanggil, nil
}

func (r *Runtime) lakukanPanggilFungsiTS(fung *ButirFungsi, argumen []interface{}) (interface{}, *KesalahanRuntime) {
	if len(argumen) != len(fung.namaArgument) {
		return nil, &KesalahanRuntime{
			pesan: fmt.Sprint("Argumen dibutuhkan ", len(fung.namaArgument), " diberikan ", len(argumen)),
		}
	}
	rt := r.anakRuntime()
	for idk, nama := range fung.namaArgument {
		rt.ts.aturVar(nama, argumen[idk])
	}

	return rt.nodeAkar(fung.nodeAkar)
}

func (r *Runtime) nodeTeks(node *parser.NodeTeks) (*jenis.Teks, *KesalahanRuntime) {
	return &jenis.Teks{
		Teks: node.Teks,
	}, nil
}

func (r *Runtime) NodeJika(node *parser.NodeJika) (interface{}, *KesalahanRuntime) {
	banding, err := r.panggilNodeBoolean(node.Kondisi)
	if err != nil {
		return nil, &KesalahanRuntime{
			BasePosisi: utils.BasePosisi{
				Baris: node.Baris,
				Kolom: node.Kolom,
			},
			pesan: err.Error(),
		}
	}

	if banding.Isi {
		hasil, err := r.panggilNode(node.NodeNode)
		if err != nil {
			return nil, err
		}

		return hasil, nil
	}

	return nil, nil
}
func (r *Runtime) nodeBalikan(node *parser.NodeBalikan) (*jenis.Balikan, *KesalahanRuntime) {
	hasil, err := r.panggilNode(node.Node)
	if err != nil {
		return nil, err
	}

	return &jenis.Balikan{
		Isi: hasil,
	}, nil
}
