package main

import (
	"os"
	"strconv"
)

// java hashCode function implemented in golang
func hash() (string, error) {
	sorStr := string([]rune(os.Args[2]))

	target := 0
	for i := 0; i < len(sorStr); i++ {
		target = 31*target + int(sorStr[i])
		target = target & 0xFFFFFFFF
	}
	if target&0x80000000 != 0 {
		target = (target ^ 0xFFFFFFFF) + 1
	}

	return strconv.Itoa(target), nil
}
