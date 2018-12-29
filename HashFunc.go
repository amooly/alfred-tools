package main

import (
	"fmt"
)

type HashFunc struct {
}

func (HashFunc) execute(str string) (string, error) {
	sorStr := string([]rune(str))

	target := 0
	for i := 0; i < len(sorStr); i++ {
		target = 31*target + int(sorStr[i])
		target = target & 0xFFFFFFFF
	}
	if target&0x80000000 != 0 {
		target = (target ^ 0xFFFFFFFF) + 1
	}

	return fmt.Sprintf("%.4d", target%1000), nil
}

func (HashFunc) tip() string {
	return "哈希分表位："
}
