package main

import (
	"log"

	"github.com/mitubaEX/gitrefzf/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
