package cmd

import (
	"fmt"
	"github.com/pete911/jwt/internal/io"
	"github.com/pete911/jwt/internal/jwt"
	"github.com/spf13/cobra"
	"log/slog"
)

var (
	decodeIndentFlag bool
	decodeAlgFlag    string
	decodeKeyFlag    string

	decodeCmd = &cobra.Command{
		Use:   "decode",
		Short: "decode jwt token",
		RunE:  decodeCmdRunE,
	}
)

func init() {
	decodeCmd.Flags().BoolVar(&decodeIndentFlag, "indent", false, "indent output")
	decodeCmd.Flags().StringVar(&decodeAlgFlag, "alg", "", "algorithm for signature validation, if it is empty, no validation is done")
	decodeCmd.Flags().StringVar(&decodeKeyFlag, "key", "", "key to validate token")
}

func decodeCmdRunE(_ *cobra.Command, args []string) error {

	input, err := io.LoadInput(args)
	if err != nil {
		return err
	}
	for _, in := range input {
		decode(in, decodeAlgFlag, decodeKeyFlag, decodeIndentFlag)
	}
	return nil
}

func decode(input, alg, key string, indent bool) {

	token, err := jwt.Decode(input, alg, key)
	if err != nil {
		slog.Error(fmt.Sprintf("decode jwt token: %v", err))
	}

	if indent {
		fmt.Println(token.StringIndent())
		return
	}
	fmt.Println(token)
}
