package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keitam913/reactive-dice/config"
	"github.com/keitam913/reactive-dice/dice"
	"github.com/keitam913/reactive-dice/rest"
)

type DI struct {
}

func (di DI) Router() *gin.Engine {
	return rest.NewRouter(di.DiceHandler())
}

func (di DI) DiceHandler() *rest.DiceHandler {
	return &rest.DiceHandler{
		Dice: di.Dice(),
	}
}

func (di DI) Dice() dice.Dice {
	return dice.New(di.Config().DiceName, di.Config().Faces)
}

func (DI) Config() config.Config {
	c, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	return c
}
