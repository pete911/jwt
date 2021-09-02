package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Flags struct {
	Decode  bool
	Indent  bool
	Silent  bool
	Version bool
	Args    []string
	usage   func()
	output  io.Writer
}

func ParseFlags() Flags {

	var flags Flags
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.BoolVar(&flags.Decode, "decode", getBoolEnv("JWT_CLI_DECODE", true),
		"decode jwt token")
	flagSet.BoolVar(&flags.Indent, "indent", getBoolEnv("JWT_CLI_INDENT", false),
		"print indented output")
	flagSet.BoolVar(&flags.Silent, "silent", getBoolEnv("JWT_CLI_SILENT", false),
		"suppress logs (warn, debug, ...)")
	flagSet.BoolVar(&flags.Version, "version", getBoolEnv("JWT_CLI_VERSION", false),
		"jwt cli version")

	flagSet.Usage = func() {
		fmt.Fprint(flagSet.Output(), "Usage: jwt [flags] [input]\n")
		flagSet.PrintDefaults()
	}
	flagSet.Parse(os.Args[1:])

	flags.Args = flagSet.Args()
	flags.usage = flagSet.Usage
	flags.output = flagSet.Output()
	return flags
}

// Usage print usage prefixed with msg, useful to provide additional error message
func (f Flags) Usage(msg string) {
	fmt.Fprintln(f.output, msg)
	f.usage()
}

func getBoolEnv(envName string, defaultValue bool) bool {

	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}

	if intValue, err := strconv.ParseBool(env); err == nil {
		return intValue
	}
	return defaultValue
}
