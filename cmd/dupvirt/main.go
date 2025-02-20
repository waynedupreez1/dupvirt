package main

import (
	"dupvirt/internal/flags"
	"dupvirt/internal/logger"
)

func main() {
    log := logger.New(logger.Info)

    flags := flags.Get(log)

    println(flags)
}
