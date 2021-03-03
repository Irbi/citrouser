package api

import "github.com/gin-gonic/gin"

type profileApi struct {}

func (a *profileApi) Routes(r gin.IRoutes) {
	r.POST("", a.create)
	r.GET("", a.get)
}

func (a *profileApi) create(ctx *gin.Context) {}

func (a *profileApi) get(ctx *gin.Context) {}
