package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam913/reactive-dice/config"
	"github.com/keitam913/reactive-dice/rest"
)

type DI struct {
}

func (DI) Router() *gin.Engine {
	return rest.NewRouter()
}

func (DI) Config() config.Config {
	c, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
