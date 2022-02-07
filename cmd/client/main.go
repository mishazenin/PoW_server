package main

import (
	"log"

	"github.com/vrischmann/envconfig"

	"mishazenin/PoW_server/src/client"
)

type config struct {
	ServerAddr string `required:"true"`
}

func main() {
	cfg := &config{}
	if err := envconfig.InitWithPrefix(cfg, "pow_tcp"); err != nil {
		log.Fatal(err)
	}

	cli := client.New(cfg.ServerAddr)
	cli.Quote()
}
