package main

import (
	"dupvirt/internal/args"
	"dupvirt/internal/logger"
	"dupvirt/internal/entrypoint"
)

func main() {
    log := logger.New(logger.Info)

    args := args.Get(log)

	entrypoint.Main(log, args)
}
