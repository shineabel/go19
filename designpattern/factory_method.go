package main

type Operator interface {
	SetA(a int)
	SetB(b int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA( a int)  {
	o.a = a
}
func (o *OperatorBase) SetB(b int)  {
	o.b = b
}

type PlusOperatorFactory struct {}

type MinusOperatorFactory struct {}

type PlusOperator struct {
	*OperatorBase
}

type MinusOperator struct {
	*OperatorBase
}
func (po *PlusOperator) Result() int {

	return po.a + po.b
}

func (mo *MinusOperator)Result() int  {
	return  mo.a - mo.b
}

func (pof PlusOperatorFactory) Create() Operator {

	return &PlusOperator{
		&OperatorBase{},
	}
}

func (pof MinusOperatorFactory) Create() Operator  {
	return &MinusOperator{
		&OperatorBase{},
	}
}

