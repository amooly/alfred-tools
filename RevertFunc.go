package main

type RevertFunc struct {
}

func (RevertFunc) execute(str string) (string, error) {
	runes := []rune(str)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes), nil
}

func (RevertFunc) tip() string {
	return "字符串反转："
}
