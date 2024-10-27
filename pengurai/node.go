package pengurai

import (
	"indoscript/lekser"
	"indoscript/utils"
)

type NodeAkar struct {
	utils.BasisPosisi
	NodeNode []interface{}
}

type NodeBilangan struct {
	utils.BasisPosisi
	Token lekser.Token
}

type NodeTeks struct {
	utils.BasisPosisi
	Teks string
}

type NodeBoolean struct {
	utils.BasisPosisi
	Isi bool
}

type NodeTidak struct {
	utils.BasisPosisi
	Node interface{}
}

type NodeOperasi struct {
	utils.BasisPosisi
	NodeKiri  interface{}
	Operasi   lekser.JenisToken
	NodeKanan interface{}
}

type NodeOperasiUner struct {
	utils.BasisPosisi
	Token lekser.JenisToken
	Node  interface{}
}

type NodeOperasiDanAtau struct {
	utils.BasisPosisi
	NodeKiri  interface{}
	Operasi   string
	NodeKanan interface{}
}

type NodeAturVariabel struct {
	utils.BasisPosisi
	NamaVariabel string
	Node         interface{}
}

type NodeAksesVariabel struct {
	utils.BasisPosisi
	NamaVariabel string
}

type NodeAturFungsi struct {
	utils.BasisPosisi
	NamaFungsi  string
	NamaArgumen []string
	NodeNode    *NodeAkar
}

type NodePanggilFungsi struct {
	utils.BasisPosisi
	NamaFungsi string
	Argumen    []interface{}
}

type NodeJika struct {
	utils.BasisPosisi
	Kondisi  interface{}
	NodeNode *NodeAkar
}

type NodeBalikan struct {
	utils.BasisPosisi
	Node interface{}
}
