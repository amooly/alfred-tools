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
			wf.NewItem("Error!!! Please open the debug mode.")
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
