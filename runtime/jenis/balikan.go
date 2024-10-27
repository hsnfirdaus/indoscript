package jenis

type Balikan struct {
	Isi interface{}
}

func UnwrapBalikan[T interface{}](balikan *Balikan) T {
	return balikan.Isi.(T)
}
