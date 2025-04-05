package encoding

func MostCommonPair(s string) (string, int) {
	if len(s) < 2 {
		return "", 0
	}
	l, r := 0, 1
	maxOccurence := 0
	mostCommon := ""
	counts := map[string]int{}
	for r < len(s) {
		left := s[l]
		right := s[r]
		pair := string(left) + string(right)
		_, ok := counts[pair]
		if ok {
			counts[pair]++
			count := counts[pair]
			if count > maxOccurence {
				mostCommon = pair
				maxOccurence = count
			}
		} else {
			counts[pair] = 1
		}
		l++
		r++
	}

	return mostCommon, maxOccurence
}
