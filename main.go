package main

import (
	"github.com/pete911/jwt/cmd"
	"log"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
