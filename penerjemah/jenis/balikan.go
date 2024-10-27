package jenis

type Balikan struct {
	Isi interface{}
}

func (balikan *Balikan) Unwrap() interface{} {
	return balikan.Isi
}
