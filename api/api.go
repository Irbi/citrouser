package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IRoutingApi interface {
	Routes(r gin.IRoutes)
}

type api struct {
	Users 		*userApi
	UserProfile *profileApi
}

func New(r *gin.Engine) (router *api) {
	router = &api{}
	router.routes(r.Group("v1/api"))
	return router
}

func (a *api) routes(r *gin.RouterGroup) {
	a.Users.Routes(r.Group("users"))
	a.UserProfile.Routes(r.Group("profile"))
}

func HandleError(ctx *gin.Context, err error, code int) bool {
	if err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(code, err.Error())
		return true
	}

	return false
}


