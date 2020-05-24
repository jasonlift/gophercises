package linked

import "testing"

func TestTwoNumbersUseCases(t *testing.T) {
	l1Table := [][]int{
		[]int{3,7},
		[]int{2,4,3},
	}
	l2Table := [][]int{
		[]int{9, 2},
		[]int{5,6,4},
	}
	resTable := [][]int {
		[]int{2,0,1},
		[]int{7,0,8},
	}

	for idx, _ := range l1Table {
		ll1 := array2LinkedList(l1Table[idx])
		ll2 := array2LinkedList(l2Table[idx])
		res := addTwoNumbers(ll1, ll2)
		ok := verifyCorrect(
			res,
			resTable[idx])
		if !ok {
			t.Errorf("Round%d: want %v", idx, resTable[idx])
		}
	}
}