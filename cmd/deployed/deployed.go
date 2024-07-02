package main

import (
	"context"
	"fmt"
	"os"

	deployedV1 "github.com/deployix/deployed/internal/app/deployed/v1"
)

func main() {
	ctx := context.Background()

	if err := deployedV1.Execute(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
