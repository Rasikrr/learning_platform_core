package main

import (
	"fmt"
	"github.com/Rasikrr/learning_platform_core/configs"
)

const (
	appName = "core"
)

func main() {
	cfg, err := configs.Parse(appName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cfg.Env)
}
