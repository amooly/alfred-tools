package main

import "strings"

type TrimFunc struct {
}

func (TrimFunc) execute(str string) (string, error) {
	return strings.TrimSpace(str), nil
}

func (TrimFunc) tip() string {
	return "清除格式："
}
