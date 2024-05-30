package main

import (
	"os"

	deployedV1 "github.com/deployix/deployed/internal/app/deployed/v1"
)

func main() {
	if err := deployedV1.Execute(); err != nil {
		os.Exit(1)
	}
}
