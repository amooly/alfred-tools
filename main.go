package main

import (
	"errors"
	"github.com/deanishe/awgo"
	"log"
	"os"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func run() {
	funcName := os.Args[1]

	var (
		result string
		err    error
	)
	if "revert" == funcName {
		result, err = revert()
	} else if "hash" == funcName {
		result, err = hash()
	} else {
		err = errors.New("不支持该操作")
	}

	if err != nil {
		log.Printf("input:%s", os.Args[1:])
		log.Printf("error:%s", err)
		wf.NewItem("运行异常，请进入调试模式排查")
	} else {
		wf.NewItem(result).Arg(result).Valid(true)
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
