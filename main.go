package main

import (
	"fmt"
	"github.com/pete911/jwt/pkg/io"
	"github.com/pete911/jwt/pkg/jwt"
	"github.com/rs/zerolog/log"
	"os"
)

var Version = "dev"

func main() {

	flags := ParseFlags()
	initLog(flags)

	if flags.Version {
		fmt.Println(Version)
		os.Exit(0)
	}

	input, err := io.LoadInput(flags.Args)
	if err != nil {
		flags.Usage("no input provided")
		os.Exit(1)
	}

	if flags.Decode {
		for _, in := range input {
			decode(in, flags.Indent)
		}
	}
}

func decode(input string, indent bool) {

	token, err := jwt.Decode(input)
	if err != nil {
		log.Error().Err(err).Msg("decode jwt token")
	}

	if indent {
		fmt.Println(token.StringIndent())
		return
	}
	fmt.Println(token)
}
