package middleware

import (
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	"github.com/Irbi/citrouser/model"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

var IdentityKey = "Email"

type AuthenticatorHandlerDef func(c *gin.Context) (interface{}, error)

type AuthorizatorHandlerDef func(data interface{}, c *gin.Context) bool

type AuthLoginResponseDef func(*gin.Context, int, string, time.Time)

type Auth struct {
	Jwt *jwt.GinJWTMiddleware
}

type TokenPayload struct {
	Email string
}

func InitJWT(realm, signKey string) *jwt.GinJWTMiddleware {
	jwtHandler, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       realm,
		Key:         []byte(signKey),
		Timeout:     time.Hour * constants.AUTH_TOKEN_EXRIPATION,
		MaxRefresh:  time.Hour * constants.AUTH_TOKEN_EXRIPATION,
		IdentityKey: IdentityKey,

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Email: claims[IdentityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if userData, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: userData.Email,
				}
			}
			return jwt.MapClaims{}
		},
	})

	return jwtHandler
}

func (a *Auth) SetAuthenticator(fn AuthenticatorHandlerDef) {
	a.Jwt.Authenticator = fn
}

func (a *Auth) SetAuthorizator(fn AuthorizatorHandlerDef) {
	a.Jwt.Authorizator = fn
}

func (a *Auth) SetLoginResponse(fn AuthLoginResponseDef) {
	a.Jwt.LoginResponse = fn
}

func (a *Auth) GetLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a.Jwt.LoginHandler(ctx)
	}
}

func (a *Auth) GetRefreshHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		a.Jwt.RefreshHandler(c)
	}
}

func (a *Auth) GenerateToken(email string) (token string, err error) {

	u, err := db.UserModel(nil).GetByEmail(email, false)
	if err != nil {
		log.Errorf("Problem while token generation: %v", err)
		return
	}

	token, exp, err := a.Jwt.TokenGenerator(u)

	log.Debugf("Token %s generated for %s :: Exp: %v", token, email, exp)

	return
}
