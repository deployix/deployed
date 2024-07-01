package main

import (
	"fmt"
	"os"

	deployedV1 "github.com/deployix/deployed/internal/app/deployed/v1"
)

func main() {
	if err := deployedV1.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
