package main

import (
	"github.com/Place1/swagger-version-manager/cli"
	"log"
	"os"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
