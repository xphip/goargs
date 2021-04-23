package goargs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	UsageComplement = "\nRun '%s' for usage."
	MissingParameter = "error: missing parameter." + UsageComplement
	UnknownCommand = "error: unknown command." + UsageComplement
)

var BasePathUsage = []string{filepath.Base(os.Args[0])}

func ErrorFormatter(error string, arr []string) error {
	return errors.New(
		fmt.Sprintf(error,
			strings.Join(arr, " ")))
}