package main

import (
	"log"
	"os"
	"github.com/Place1/swagger-version-manager/cli"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
