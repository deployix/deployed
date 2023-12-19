package main

import (
	"os"

	"github.com/deployix/deployed/internal/app/deployed"
)

func main() {
	if err := deployed.Execute(); err != nil {
		os.Exit(1)
	}
}
