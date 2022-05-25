package main

import (
	"URLShortner/model"
	"URLShortner/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
