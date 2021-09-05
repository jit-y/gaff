package main

import (
	"fmt"
	"os"

	"github.com/jit-y/gaff"
)

type exitCode int

const (
	exitOk  exitCode = 0
	exitErr exitCode = 1
)

func main() {
	code := run()

	os.Exit(int(code))
}

func run() exitCode {
	cmd := gaff.NewRootCmd()

	_, err := cmd.ExecuteC()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitErr
	}

	return exitOk
}
