package main

import (
	"pocket-serv/bootstrap"
)

func main() {
	bootstrap.InitializeConfig()

	bootstrap.RunServer()
}
