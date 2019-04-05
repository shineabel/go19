package main

type Cloneable interface {
	Clone() Cloneable
}


type PrototypeManager struct {
	prototype map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager  {
	return &PrototypeManager{
		prototype:make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Get(name string) Cloneable  {

	return p.prototype[name]
}

func (p *PrototypeManager) Set( name string, c Cloneable)  {

	p.prototype[name] = c
}