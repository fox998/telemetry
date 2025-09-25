package main

import (
	"flag"

	"github.com/fox998/telemetry/internal/common"
	"github.com/fox998/telemetry/internal/sink"
)

func main() {

	sinkArgs := sink.SetSinkArgs()

	flag.Parse()

	sigChan := common.GetStop()

	sink.ValidateSinkArgs(sinkArgs)
	sink.RunSink(sigChan, sinkArgs)
}
