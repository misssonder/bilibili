package main

import (
	"fmt"
	"os"

	bilibili "github.com/misssonder/bilibili/pkg/client"
)

var (
	client *bilibili.Client
)

func init() {
	client = &bilibili.Client{}
}

func main() {
	exitOnError(rootCmd.Execute())
}

func exitOnError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
