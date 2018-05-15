package main

type d interface {
	D()
}

type e interface {
	E()
}

type f interface {
	d
	e
}
type g struct{}

func (g *g) D() {}
func (g *g) E() {}

func main() {
	var foo f = &g{}
	foo.D()
	foo.E()
}
