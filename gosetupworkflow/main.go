package main

import (
	"flag"
	"fmt"

	"github.com/chakra/gosetupworkflow/workflow"
)

/* Options This is the structure of our command line options*/
type Options struct {
	bpmnDirectory *string
}

func cmdParse() Options {
	var options Options
	options.bpmnDirectory = flag.String("dir", "", "Directory location for the BPMN files ")

	flag.Parse()
	return options

}

func mainProcess(opt Options) {
	workflow.Setup(*opt.bpmnDirectory)
	fmt.Println("Finished setup of Workflow")
}

func main() {
	options := cmdParse()
	fmt.Println(*options.bpmnDirectory)
	mainProcess(options)

}
