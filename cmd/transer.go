package main

import (
	"fmt"
	"os"
	"github.com/tetor/transer/runner"
)

func main() {
	if err := runner.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
