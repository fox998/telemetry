package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fox998/telemetry/internal/node"
	"github.com/fox998/telemetry/internal/sink"
)

func main() {
	isNode := flag.Bool("node", false, "type of application: node or sink")
	isSink := flag.Bool("sink", false, "type of application: node or sink")
	nodeArgs := node.SetNodeArgs()
	sinkArgs := sink.SetSinkArgs()

	flag.Parse()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Application is running. Press Ctrl+C to stop.")

	if *isNode == *isSink {
		log.Fatal("Expected onlly one of [ node, sink ].")
	}

	if *isNode {
		node.ValidateNodeArgs(nodeArgs)
		node.RunNode(sigChan, nodeArgs)
	} else {
		sink.ValidateSinkArgs(sinkArgs)
		sink.RunSink(sigChan, sinkArgs)
	}
}
