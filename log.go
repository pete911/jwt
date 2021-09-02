package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
)

func initLog(flags Flags) {

	if flags.Silent {
		log.Logger = log.Output(ioutil.Discard)
		return
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
}
