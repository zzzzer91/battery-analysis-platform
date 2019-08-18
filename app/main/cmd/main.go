package main

import (
	"battery-anlysis-platform/app/main/server"
	"log"
)

func main() {
	if err := server.Start(); err != nil {
		log.Println(err)
	}
}
