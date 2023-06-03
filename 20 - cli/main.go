package main

import (
	"example.com/cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("ponto de partida")

	application := app.Gerar()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
