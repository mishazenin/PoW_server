package main

import (
	"log"

	"github.com/vrischmann/envconfig"

	"github.com/LarsFox/pow-tcp/src/client"
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
