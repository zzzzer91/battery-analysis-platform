package main

import (
	"battery-anlysis-platform/app/main/server"
	"battery-anlysis-platform/pkg/conf"
)

func main() {
	if err := server.Start(&conf.App.Main.Gin); err != nil {
		panic(err)
	}
}
