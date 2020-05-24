package strings

/**
76
https://leetcode.com/problems/minimum-window-substring/description/
 */

func minWindow(s string, t string) string {
	requirements := make(map[byte]int)
	for i := range t {
		if _, ok := requirements[t[i]]; ok {
			requirements[t[i]]++
		} else {
			requirements[t[i]] = 1
		}
	}

	appearances := make(map[byte]int)
	minLen := len(s)+1
	counter := len(t)
	resLeft, resRight := 0, 0
	for left, right := 0, 0; right < len(s); {
		rch, lch := s[right], s[left]
		if _, ok := requirements[rch]; ok {
			appearances[rch]++
			if appearances[rch] <= requirements[rch] && counter != 0 {
				counter-- // meant to find the trigger of the first time full
			}
		}

		if counter == 0 {
			for {
				_, ok := requirements[lch]
				if !ok || appearances[lch] > requirements[lch] {
					// if not find lch in requirements map, then enter
					// otherwise, if counter that `appearances` map contains is
					// more than `requirements` has
					// which means right index traverse and add new character in `target`
					if appearances[lch] > requirements[lch] {
						appearances[lch]--
					}
					left++
					lch = s[left]
				} else {
					break
				}
			}
			if right-left+1 < minLen {
				minLen = right-left+1
				resLeft = left
				resRight = right+1 // because of the slice index
			}
		}
		right++
	}

	return s[resLeft:resRight]
}

/**
similar question in niuke
https://www.nowcoder.com/questionTerminal/58569ba19c05436e9eb492244b0902d8?orderByHotValue=1&mutiTagIds=645&page=1&onlyReference=false
 */

