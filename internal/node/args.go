package node

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type NodeArgs struct {
	Rate *uint
	Name *string
	Addr *string
}

func SetNodeArgs() NodeArgs {
	return NodeArgs{
		Rate: flag.Uint("rate", 1, "number of messages per second"),
		Name: flag.String("sensor_name", fmt.Sprintf("sensor_%d", rand.Int()%100), "the name of the sensor to use"),
		Addr: flag.String("node_addr", "localhost:8080", "Address of the telemetry sink"),
	}
}

func ValidateNodeArgs(args NodeArgs) {
	if *args.Rate <= 0 {
		log.Fatal("Invalid rate")
	}

	if *args.Name == "" {
		log.Fatal("Invalid sensor name")
	}

	if *args.Addr == "" {
		log.Fatal("Invalid server address")
	}

	if !strings.Contains(*args.Addr, ":") {
		log.Fatal("Invalid server address format, expected host:port")
	}
}
