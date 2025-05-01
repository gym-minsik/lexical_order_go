package lexical_order

type AlphabetOrderService struct {
	keySpace *KeySpace
}

func NewAlphabetKeySpace() *KeySpace {
	Keys := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	KeyToIndex := make(map[rune]int)
	for i, k := range Keys {
		KeyToIndex[CharAt(k, 0)] = i
	}

	return &KeySpace{
		keys:       Keys,
		keyToIndex: KeyToIndex,
	}
}

func NewAlphabetOrderService() *AlphabetOrderService {
	return &AlphabetOrderService{
		keySpace: NewAlphabetKeySpace(),
	}
}

func (a *AlphabetOrderService) Generate(count int) []string {
	return GenerateOrderKeys(a.keySpace.keys, count)
}

func (a *AlphabetOrderService) Between(prev string, next string) string {
	return Between(NewAlphabetKeySpace().keyToIndex, NewAlphabetKeySpace().keys, prev, next)
}
