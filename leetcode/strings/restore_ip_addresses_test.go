package strings

import "testing"

func TestRestoreIp(t *testing.T) {
	input := "25525511135"
	res := restoreIpAddresses(input)
	t.Logf("%v", res)
}
