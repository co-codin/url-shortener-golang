package main

import (
	"github.com/co-codin/model"
	"github.com/co-codin/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}