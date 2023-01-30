package main

import (
	"fmt"
	"os"
)

func main() {
	exitOnError(rootCmd.Execute())
}

func exitOnError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
