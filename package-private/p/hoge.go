package p

import "fmt"

type Hoge struct{}

type hogeConst struct {
	name string
}

func NewHoge() Hoge {
	return Hoge{}
}

func (h *Hoge) c() hogeConst {
	return hogeConst{
		name: "hoge",
	}
}

func (h *Hoge) Print() error {
	_, err := fmt.Println(h.c().name)
	if err != nil {
		return err
	}
	return nil
}
