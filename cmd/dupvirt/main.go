package main

import (
	"github.com/waynedupreez1/duptfy/internal/cli"
	"github.com/waynedupreez1/duptfy/internal/flags"
	"github.com/waynedupreez1/duptfy/internal/logger"
)

func main() {
    log := logger.New(logger.Info)

    flags := flags.Get(log)

    cli := cli.New(log, flags)

    cli.Main()
}
