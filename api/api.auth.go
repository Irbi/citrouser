package api

import (
	"errors"
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	authReq "github.com/Irbi/citrouser/requests/auth"
	authResp "github.com/Irbi/citrouser/responses/auth"
	"github.com/Irbi/citrouser/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiAuth struct {}

func (a *apiAuth) Routes(r gin.IRoutes) {
	r.POST("login", a.login)
}

func (a *apiAuth) login(ctx *gin.Context) {
	req := &authReq.LoginRequest{}
	err := ctx.Bind(req)

	if err != nil {
		HandleError(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := db.UserModel(nil).GetByEmail(req.Email, false)
	if err != nil || !tools.CheckPasswordFromHash(req.Password, user.Password){
		HandleError(ctx, errors.New(constants.ERR_USER_INVALID_CREDENTIALS), http.StatusUnauthorized)
		return
	}

	if user.Status != constants.USER_STATUS_ACTIVE {
		HandleError(ctx, errors.New(constants.ERR_USER_INACTIVE), http.StatusForbidden)
		return
	}

	ctx.JSON(http.StatusOK, authResp.LoginResponse{
		ID: user.ID,
		Email: user.Email,
		Status: user.Status,
		Role: user.Role,
	})
}
