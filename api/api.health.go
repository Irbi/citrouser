package api

import (
	"github.com/Irbi/citrouser/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthApi struct {}

func (a *healthApi) Routes(r gin.IRoutes) {
	r.GET("", a.health)
}

func (a * healthApi) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responses.SuccessResponse{Success: true})
}