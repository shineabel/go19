package cacl

import "testing"

func TestAdd(t *testing.T) {
	r := Add(2,4)
	if r != 6{
		t.Fatal("test failed")
	}
	t.Logf("test success")
}
