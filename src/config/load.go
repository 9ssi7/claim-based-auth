package config

import (
	"go.deanishe.net/env"
)

func LoadConfig(c interface{}) {
	err := env.Bind(c)
	if err != nil {
		panic(err)
	}
}
