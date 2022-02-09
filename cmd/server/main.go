package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"mishazenin/PoW_server/src/hashcash"
	"mishazenin/PoW_server/src/library"
	"mishazenin/PoW_server/src/server"
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
	POW_TCP_TARGET_BITS, ok := os.LookupEnv("POW_TCP_TARGET_BITS")
	if !ok {
		log.Fatal(errors.New("ENV is not ser"))
	}
	bits, _ := strconv.Atoi(POW_TCP_TARGET_BITS)

	cfg := &config{
		Addr:       addr,
		TargetBits: int64(bits),
	}

	hc := hashcash.New(cfg.TargetBits)
	POWserver := server.NewPOWServer(library.Quotes, *hc)

	log.Println("Listening on", cfg.Addr)

	POWserver.Listen(cfg.Addr)
}
