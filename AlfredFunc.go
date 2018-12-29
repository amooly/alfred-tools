package main

type AlfredFunc interface {
	// 执行函数
	execute(str string) (string, error)
	// 提示信息
	tip() string
}
