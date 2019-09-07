package main

import (
	"battery-anlysis-platform/app/main/server"
)

func main() {
	if err := server.Start(); err != nil {
		panic(err)
	}
}
