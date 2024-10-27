package parser

import (
	"indoscript/lexer"
	"indoscript/utils"
)

type NodeAkar struct {
	NodeNode []interface{}
}

type NodeBilangan struct {
	utils.BasePosisi
	Token lexer.Token
}

type NodeTeks struct {
	utils.BasePosisi
	Teks string
}

type NodeTidak struct {
	utils.BasePosisi
	Node interface{}
}

type NodeOperasi struct {
	utils.BasePosisi
	NodeKiri  interface{}
	Operasi   lexer.JenisToken
	NodeKanan interface{}
}

type NodeOperasiUner struct {
	utils.BasePosisi
	Token lexer.JenisToken
	Node  interface{}
}

type NodeAturVariabel struct {
	utils.BasePosisi
	NamaVariabel string
	Node         interface{}
}

type NodeAksesVariabel struct {
	utils.BasePosisi
	NamaVariabel string
}

type NodeAturFungsi struct {
	utils.BasePosisi
	NamaFungsi  string
	NamaArgumen []string
	NodeNode    *NodeAkar
}

type NodePanggilFungsi struct {
	utils.BasePosisi
	NamaFungsi string
	Argumen    []interface{}
}

type NodeJika struct {
	utils.BasePosisi
	Kondisi  interface{}
	NodeNode *NodeAkar
}

type NodeBalikan struct {
	utils.BasePosisi
	Node interface{}
}
