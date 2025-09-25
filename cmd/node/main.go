package main

import (
	"flag"

	"github.com/fox998/telemetry/internal/common"
	"github.com/fox998/telemetry/internal/node"
)

func main() {
	nodeArgs := node.SetNodeArgs()

	flag.Parse()

	sigChan := common.GetStop()

	node.ValidateNodeArgs(nodeArgs)
	node.RunNode(sigChan, nodeArgs)
}
