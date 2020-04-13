package rest

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "/usr/share/reactive-dice/assets/static")
	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("/usr/share/reactive-dice/assets/index.html")
	})

	return r
}
