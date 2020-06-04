package strings

/**
415
https://leetcode.com/problems/add-strings/

easier solution: https://leetcode.com/problems/add-strings/discuss/523226/golang-0ms-100-solution
 */
func addStrings(num1 string, num2 string) string {
	l := 0
	if len(num1) > len(num2) {
		l = len(num1)
	} else {
		l = len(num2)
	}
	var res []byte = make([]byte, l+1)
	i := len(num1)-1
	j := len(num2)-1
	for i >= 0 && j >= 0 {
		tmp := (num1[i]-'0')+(num2[j]-'0')
		res[l] += tmp
		if res[l] >= 10 {
			res[l-1] += res[l]/10
			res[l] = res[l] % 10
		}
		i--
		j--
		l--
	}
	for i >= 0 && l >= 0 {
		res[l] += num1[i]-'0'
		if res[l] >= 10 {
			res[l-1] += res[l]/10
			res[l] = res[l] % 10
		}
		l--;i--
	}
	for j >= 0 && l >= 0 {
		res[l] += num2[j]-'0'
		if res[l] >= 10 {
			res[l-1] += res[l]/10
			res[l] = res[l] % 10
		}
		l--;j--
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for x := range res {
		res[x] = res[x]+'0'
	}
	return string(res)
}

func add(s1 string, s2 string) string {
	if s1 == "0" {
		return s2
	}
	if s2 == "0" {
		return s1
	}
	l := 0
	if len(s1) > len(s2) {
		l = len(s1)
	} else {
		l = len(s2)
	}
	var res []byte = make([]byte, l+1)
	i := len(s1)-1
	j := len(s2)-1
	for i >= 0 && j >= 0 {
		tmp := (s1[i]-'0')+(s2[j]-'0')
		res[l] += tmp
		if res[l] >= 10 {
			res[l] = res[l] % 10
			res[l-1] += res[l]/10
		}
		i--
		j--
		l--
	}
	if i >= 0 && l >= 0 {
		res[l] = s1[i]-'0'
		l--
	}
	if j >= 0 && l >= 0 {
		res[l] = s2[j]-'0'
		l--
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for x := range res {
		res[x] = res[x]+'0'
	}
	return string(res)
}
