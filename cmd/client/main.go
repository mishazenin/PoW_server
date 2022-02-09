package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"mishazenin/PoW_server/src/client"
	"os"
)

type config struct {
	ServerAddr string `required:"true"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(".env file not found")
	}
	addr, ok := os.LookupEnv("POW_TCP_SERVER_ADDR")
	if !ok {
		log.Fatal(errors.New("ENV is not set"))
	}
	cfg := &config{ServerAddr: addr}

	client := client.New(cfg.ServerAddr)
	client.GetQuote()
}
