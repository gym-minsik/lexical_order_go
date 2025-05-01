package lexical_order

func Between(
	keyToIndex map[rune]int,
	keys []string,
	prev string,
	next string,
) string {
	var p, n int
	g := 0
	result := ""

	// 공통 prefix 탐색
	for g == 0 || p == n {
		if g < len(prev) {
			p = keyToIndex[CharAt(prev, g)]
		} else {
			p = -1
		}

		if g < len(next) {
			n = keyToIndex[CharAt(next, g)]
		} else {
			n = len(keys)
		}
		g++
	}

	if g-1 > 0 {
		result += prev[:g-1]
	}

	// case: prev가 next의 prefix인 경우
	if p == -1 {
		for n == 0 {
			if g < len(next) {
				n = keyToIndex[CharAt(next, g)]
			} else {
				n = len(keys)
			}
			result += keys[0]
			g++
		}
		if n == 1 {
			result += keys[0]
			n = len(keys)
		}
	} else if p == n-1 {
		// 연속적인 경우 처리
		result += keys[p]
		n = len(keys)

		if g < len(prev) {
			p = keyToIndex[CharAt(prev, g)]
		} else {
			p = -1
		}

		for p == len(keys)-1 {
			result += keys[len(keys)-1]
			g++
			if g < len(prev) {
				p = keyToIndex[CharAt(prev, g)]
			} else {
				p = -1
			}
		}
	}

	// 중간값 문자 삽입
	median := (p + n + 1) / 2 // equivalent to ceil((p+n)/2)
	result += keys[median]

	return result
}
