package collections

import "testing"

func TestMapWithFuncVal(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int{return op}
	m[2] = func(op int) int{return op*op}
	m[3] = func(op int) int{return op*op*op}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func verifyExisted(t *testing.T, mySet map[int]bool, n int) {
	if mySet[n] {
		t.Logf("%d is existed", n)
	} else {
		t.Logf("%d is not existed", n)
	}
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	verifyExisted(t, mySet, n)
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	verifyExisted(t, mySet, n)
}