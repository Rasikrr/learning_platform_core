package main

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/application"
)

const (
	appName = "core"
)

func main() {
	ctx := context.Background()
	app := application.NewApp(ctx, appName)
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
}
