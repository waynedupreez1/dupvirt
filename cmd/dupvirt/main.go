package main

import (
	"dupvirt/internal/args"
	"dupvirt/internal/logger"
)

func main() {
    log := logger.New(logger.Info)

    args := args.Get(log)
}
