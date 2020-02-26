package main

import (
	"UnityWebPlayer/infrastructure"
	"flag"
	"fmt"
)

func main() {
	p := flag.Int("p", 8080, "int flag")
	flag.Parse()

	port := fmt.Sprintf(":%d", *p)
	infrastructure.Router.Run(port)
}
