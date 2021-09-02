package cmd

import (
	"errors"
	"fmt"
	"github.com/pete911/jwt/internal/io"
	"github.com/pete911/jwt/internal/jwt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	indentFlag bool

	decodeCmd = &cobra.Command{
		Use:   "decode",
		Short: "decode jwt token",
		Args:  cobra.MinimumNArgs(1),
		RunE:  decodeCmdRunE,
	}
)

func init() {
	decodeCmd.Flags().BoolVar(&indentFlag, "indent", false, "indent output")
}

func decodeCmdRunE(_ *cobra.Command, args []string) error {

	input, err := io.LoadInput(args)
	if err != nil {
		return errors.New("no input provided")
	}
	for _, in := range input {
		decode(in, indentFlag)
	}
	return nil
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
