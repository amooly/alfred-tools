package main

import (
	"github.com/deanishe/awgo"
	"log"
	"os"
)

var (
	wf *aw.Workflow

	functionList = []AlfredFunc{HashFunc{}, RevertFunc{}, TrimFunc{}}
)

func init() {
	wf = aw.New()
}

func run() {
	sourceStr := os.Args[1]

	for _, function := range functionList {
		result, err := function.execute(sourceStr)

		if err != nil {
			log.Printf("input:%s", os.Args[1:])
			log.Printf("error:%s", err)
			wf.NewItem("运行异常，请进入调试模式排查")
			break
		} else {
			wf.NewItem(function.tip() + result).Arg(result).Valid(true)
		}
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
