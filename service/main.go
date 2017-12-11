package main

import (
	"entities"
	"os"

	"service"

	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	dbPath := flag.StringP("db", "d", "", "sqlite database file")
	flag.Parse()

	if len(*pPort) != 0 {
		port = *pPort
	}
	if len(*dbPath) == 0 {
		os.Mkdir("data", 0755)
		*dbPath = "data/agenda.db"
	}

	entities.Init(*dbPath)
	server := service.NewServer()
	server.Run(":" + port)
}
