package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam913/reactive-dice/rest"
)

type DI struct {
}

func (DI) Router() *gin.Engine {
	return rest.NewRouter()
}
