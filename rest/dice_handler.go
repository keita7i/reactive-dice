package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keitam913/reactive-dice/dice"
)

type DiceHandler struct {
	Dice dice.Dice
}

func (h DiceHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, h.Dice.Roll())
}
