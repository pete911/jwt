package cmd

import (
	"fmt"
	"github.com/pete911/jwt/internal/io"
	"github.com/pete911/jwt/internal/jwt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	encodeAlgFlag string
	encodeKeyFlag string

	encodeCmd = &cobra.Command{
		Use:   "encode",
		Short: "encode jwt token",
		RunE:  encodeCmdRunE,
	}
)

func init() {
	encodeCmd.Flags().StringVar(&encodeAlgFlag, "alg", "", "algorithm for signature validation, if it is empty, no validation is done")
	encodeCmd.Flags().StringVar(&encodeKeyFlag, "key", "", "key to validate token")
}

func encodeCmdRunE(_ *cobra.Command, args []string) error {

	input, err := io.LoadInput(args)
	if err != nil {
		return err
	}
	for _, in := range input {
		encode(in, encodeAlgFlag, encodeKeyFlag)
	}
	return nil
}

func encode(claims, alg, key string) {

	token, err := jwt.Encode(claims, alg, key)
	if err != nil {
		log.Error().Err(err).Msg("encode jwt token")
	}
	fmt.Println(token)
}
