package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"mishazenin/PoW_server/src/hashcash"
	"mishazenin/PoW_server/src/library"
	"mishazenin/PoW_server/src/server"
	"os"
	"strconv"
)

// init is invoked before main()
//func init() {
//	// loads values from .env into the system
//	if err := godotenv.Load(); err != nil {
//		log.Print("No .env file found")
//	}
//}

type config struct {
	Addr       string `required:"true"`
	TargetBits int64  `envconfig:"default=24"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
	// Get the GITHUB_USERNAME environment variable
	POW_TCP_ADDR, _ := os.LookupEnv("POW_TCP_ADDR")
	POW_TCP_TARGET_BITS, _ := os.LookupEnv("POW_TCP_TARGET_BITS")
	bits, _ := strconv.Atoi(POW_TCP_TARGET_BITS)

	cfg := &config{Addr: POW_TCP_ADDR, TargetBits: int64(bits)}

	//if err := config.InitWithPrefix(cfg, "pow_tcp"); err != nil {
	//	log.WithField("server", "config").Errorf("Couldn't read env config")
	//}

	hc := hashcash.New(cfg.TargetBits)
	server := server.NewPOWServer(library.BookDostoevsky, hc)

	log.Println("Listening on", cfg.Addr)
	server.Listen(cfg.Addr)
}
