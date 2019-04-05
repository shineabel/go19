package main

import "testing"

var pm *PrototypeManager

func TestPrototypeManager_Get(t *testing.T) {

}


type Type1 struct {

	name string
}

func (t *Type1) Clone() Cloneable  {
	tc := *t
	return &tc
}



type Type2 struct {

	name string
}

func (t *Type2)  Clone() Cloneable  {
	tc := *t
	return &tc
}

func TestClone( t *testing.T)  {

	t1 := pm.Get("type1")
	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("error t1 == t2")
	}
}

func init()  {
 pm = NewPrototypeManager()
  type1 := &Type1{

  	name:"type1",
  }

  pm.Set("type1",type1)
}