package main

import (
	"errors"
	"os"
	"strconv"

	"mishazenin/PoW_server/cmd/server/internal"
	"mishazenin/PoW_server/pkg"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Addr       string `required:"true"`
	TargetBits int64  `envconfig:"default=24"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}

	addr, ok := os.LookupEnv("POW_TCP_ADDR")
	targetBits, ok := os.LookupEnv("POW_TCP_TARGET_BITS")
	if !ok {
		log.Fatal(errors.New("ENV is not ser"))
	}
	bits, _ := strconv.Atoi(targetBits)

	cfg := &config{
		Addr:       addr,
		TargetBits: int64(bits),
	}

	hc := pkg.New(cfg.TargetBits)
	POWserver := internal.NewPOWServer(internal.ItQuotes, *hc)

	log.Printf("Listening on: %s", cfg.Addr)

	POWserver.Listen(cfg.Addr)
}
