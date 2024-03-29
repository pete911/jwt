package io

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// LoadInput checks stdin, if there is nothing in stdin then it returns default input
func LoadInput(defaultInput []string) ([]string, error) {

	stdin, err := isStdin()
	if err != nil {
		return nil, err
	}

	if stdin {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("read stdin: %w", err)
		}
		return []string{string(content)}, nil
	}

	// no stdin and not args
	if len(defaultInput) == 0 {
		return nil, errors.New("no input received")
	}
	return defaultInput, nil
}

func isStdin() (bool, error) {

	info, err := os.Stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("stdin stat: %w", err)
	}

	if (info.Mode() & os.ModeCharDevice) == 0 {
		return true, nil
	}
	return false, nil
}
