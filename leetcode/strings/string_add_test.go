package strings

import "testing"

func TestForStringAdd(t *testing.T) {
	s1 := "123"
	s2 := "489"
	s3 := add(s1, s2)
	if len(s3) != 0 {
		t.Logf("get %v", s3)
	}
}
