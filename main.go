package main

import (
	"os"

	"github.com/kaust-rc/licensebeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
