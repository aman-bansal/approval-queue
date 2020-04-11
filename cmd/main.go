package main

import (
	"context"
	"github.com/aman-bansal/approval-queue/internal"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		log.Fatal("No app environment present")
	}

	if err := internal.Init(ctx, appEnv); err != nil {
		log.Fatal(err)
	}
}
