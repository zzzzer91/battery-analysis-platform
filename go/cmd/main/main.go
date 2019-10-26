package main

import (
	"battery-analysis-platform/app/main/server"
)

func main() {
	if err := server.Start(); err != nil {
		panic(err)
	}
}
