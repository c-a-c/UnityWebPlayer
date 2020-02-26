package main

import (
	"UnityWebPlayer/di"
	"UnityWebPlayer/infrastructure"
)

func main() {
	di.Init()
	infrastructure.Router.Run(":8080")
}
