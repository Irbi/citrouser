package api

import (
	"errors"
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	"github.com/Irbi/citrouser/middleware"
	"github.com/Irbi/citrouser/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/appleboy/gin-jwt/v2"
	"os"
)

type IRoutingApi interface {
	Routes(r gin.IRoutes)
}

type api struct {
	Health		*healthApi
	Auth 		*authApi
	Users 		*userApi
	UserProfile *profileApi
	Portfolio 	*portfolioApi
	Clients		*clientApi
}

func New(r *gin.Engine) (router *api) {

	log.Debug("JWT_KEY: ", os.Getenv("JWT_KEY"))
	log.Debug("WEB_URL: ", os.Getenv("WEB_URL"))

	userHandler := middleware.InitJWT(os.Getenv("WEB_URL"), os.Getenv("JWT_KEY"))
	userAuth := &middleware.Auth{
		Jwt: userHandler,
	}

	authApi := &authApi{}
	authApi.Auth = userAuth

	router = &api{
		Auth: authApi,
	}

	router.routes(r.Group("v1/api"))
	return router
}

func (a *api) routes(r *gin.RouterGroup) {

	jwtMiddleware := a.Auth.Auth.Jwt.MiddlewareFunc()

	a.Health.Routes(r.Group("health"))
	a.Auth.Routes(r.Group("auth"))
	a.Users.Routes(r.Group("users").Use(jwtMiddleware))
	a.Clients.Routes(r.Group("clients").Use(jwtMiddleware))
	a.UserProfile.Routes(r.Group("profile").Use(jwtMiddleware))
	a.Portfolio.Routes(r.Group("portfolio").Use(jwtMiddleware))
}

func HandleError(ctx *gin.Context, err error,  entry *log.Entry, code int) bool {
	if err != nil {
		if entry != nil {

			endpoint := "api error: "

			if len(entry.Data) != 0 {
				if _, ok := entry.Data["api"]; ok {
					endpoint = entry.Data["api"].(string)
				}
				if _, ok := entry.Data["event"]; ok {
					endpoint = endpoint + " :: " + entry.Data["event"].(string)
				}
			}
			entry.Warnf("%s failed. Error: %s", endpoint, err.Error())

		} else {
			log.Warn("api error: ", err)
		}

		log.Error(err)
		ctx.AbortWithStatusJSON(code, err.Error())
		return true
	}

	return false
}

func GetUser(ctx *gin.Context, preloadProfile bool) (user *model.User, err error) {
	claims := jwt.ExtractClaims(ctx)
	userEmail := claims[middleware.IdentityKey]
	if userEmail == nil {
		return nil, errors.New(constants.ERR_AUTH_EMPTY_JWT_CREDENTIALS)
	}
	user, err = db.UserModel(nil).GetByEmail(userEmail.(string), preloadProfile)
	return
}


