package main

import (
	"log"
	"os"
	"strconv"
)

func hash() (string, error) {
	sorStr := string([]rune(os.Args[2]))

	tgtStr := 0
	for i := 0; i < len(sorStr); i++ {
		tgtStr = 31*tgtStr + int(sorStr[i])
	}
	log.Print(tgtStr)
	return strconv.Itoa(tgtStr), nil
}
