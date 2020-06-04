package strings

import "testing"

func TestForStringAdd(t *testing.T) {
	s1 := "123"
	s2 := "489"
	s3 := addStrings(s1, s2)
	if len(s3) != 0 {
		t.Logf("get %v", s3)
	}

	s1 = "9133"
	s2 = "0"
	s3 = addStrings(s1, s2)
	if s3 != "9133" {
		t.Errorf("get %v", s3)
	}

	s1 = "9"
	s2 = "1"
	s3 = addStrings(s1, s2)
	if s3 != "10" {
		t.Errorf("get %v", s3)
	}
}
