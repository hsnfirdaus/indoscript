package penerjemah

import (
	"indoscript/penerjemah/jenis"
	"indoscript/pengurai"
	"reflect"
)

type Penerjemah struct {
	ts *TabelSimbol
}

func PenerjemahBaru() Penerjemah {
	p := Penerjemah{}
	ts := TabelSimbolBaru(nil)
	p.ts = &ts

	return p
}

func (p *Penerjemah) buatAnak() Penerjemah {
	ts := TabelSimbolBaru(p.ts)
	pb := Penerjemah{
		ts: &ts,
	}

	return pb
}

func (p *Penerjemah) Jalankan(node *pengurai.NodeAkar) *KesalahanPenerjemah {
	_, err := p.panggilNode(node)
	return err
}

func (p *Penerjemah) panggilNode(node interface{}) (interface{}, *KesalahanPenerjemah) {

	switch v := node.(type) {
	case *pengurai.NodeAkar:
		return p.nodeAkar(v)
	case *pengurai.NodeBilangan:
		return p.nodeBilangan(v)

	case *pengurai.NodeTeks:
		return p.nodeTeks(v)

	case *pengurai.NodeBoolean:
		return p.nodeBoolean(v)

	case *pengurai.NodeOperasi:
		return p.nodeOperasi(v)

	case *pengurai.NodeOperasiUner:
		return p.nodeOperasiUner(v)

	case *pengurai.NodeOperasiDanAtau:
		return p.nodeOperasiDanAtau(v)

	case *pengurai.NodeAturVariabel:
		return p.nodeAturVariabel(v)

	case *pengurai.NodeAksesVariabel:
		return p.nodeAksesVariabel(v)

	case *pengurai.NodeAturFungsi:
		return p.nodeAturFungsi(v)

	case *pengurai.NodePanggilFungsi:
		return p.nodePanggilFungsi(v)

	case *pengurai.NodeJika:
		return p.NodeJika(v)

	case *pengurai.NodeBalikan:
		return p.nodeBalikan(v)

	default:
		jenisNode := reflect.TypeOf(v)
		return nil, &KesalahanPenerjemah{
			pesan: "Node \"" + jenisNode.String() + "\" tak dikenali!",
		}
	}
}

func (p *Penerjemah) panggilNodeBilangan(node interface{}) (*jenis.Bilangan, *KesalahanPenerjemah) {
	isi, err := p.panggilNode(node)
	if err != nil {
		return nil, err
	}

	isiBil, ok := isi.(*jenis.Bilangan)
	if !ok {
		return nil, &KesalahanPenerjemah{
			// TODO: How to get baris kolom without assert one by one
			pesan: "Bukan merupakan bilangan!",
		}
	}

	return isiBil, nil
}

func (p *Penerjemah) panggilNodeBoolean(node interface{}) (*jenis.Boolean, *KesalahanPenerjemah) {
	isi, err := p.panggilNode(node)
	if err != nil {
		return nil, err
	}

	isiBool, ok := isi.(*jenis.Boolean)
	if !ok {
		jenisAsli := reflect.TypeOf(isi)
		return nil, &KesalahanPenerjemah{
			// TODO: How to get baris kolom without assert one by one
			pesan: jenisAsli.String() + " bukan merupakan boolean!",
		}
	}

	return isiBool, nil
}
func (p *Penerjemah) panggilNodeTeks(node interface{}) (*jenis.Teks, *KesalahanPenerjemah) {
	isi, err := p.panggilNode(node)
	if err != nil {
		return nil, err
	}

	isiTeks, ok := isi.(*jenis.Teks)
	if !ok {
		jenisAsli := reflect.TypeOf(isi)
		return nil, &KesalahanPenerjemah{
			// TODO: How to get baris kolom without assert one by one
			pesan: jenisAsli.String() + " bukan merupakan boolean!",
		}
	}

	return isiTeks, nil
}
