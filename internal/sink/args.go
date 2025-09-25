package sink

import (
	"flag"
	"log"
	"strings"
)

type SinkArgs struct {
	BufferSize    *int
	WriteInterval *uint
	LogFile       *string
	Addr          *string
}

func SetSinkArgs() SinkArgs {
	return SinkArgs{
		BufferSize:    flag.Int("buffer_size", 100, "Size of log buffer"),
		WriteInterval: flag.Uint("write_interval", 100, "Delta time between flush triggers"),
		LogFile:       flag.String("log_file", "sink.log", "Log file name"),
		Addr:          flag.String("sink_addr", "localhost:8080", "Address to listen on"),
	}
}

func ValidateSinkArgs(args SinkArgs) {
	if *args.BufferSize <= 0 {
		log.Fatal("Invalid buffer size")
	}
	if *args.WriteInterval <= 0 {
		log.Fatal("Invalid write interval")
	}
	if *args.LogFile == "" {
		log.Fatal("Invalid log file")
	}
	if *args.Addr == "" {
		log.Fatal("Invalid server address")
	}

	if !strings.Contains(*args.Addr, ":") {
		log.Fatal("Invalid server address format, expected host:port")
	}
}
