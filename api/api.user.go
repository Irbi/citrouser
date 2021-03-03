package api

import (
	"errors"
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	"github.com/Irbi/citrouser/requests/userRequests"
	"github.com/Irbi/citrouser/responses"
	"github.com/Irbi/citrouser/responses/userResponses"
	"github.com/Irbi/citrouser/tools"
	"github.com/Irbi/citrouser/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userApi struct {}

func (a *userApi) Routes(r gin.IRoutes) {
	r.POST("", a.create)
	r.GET("", a.getAll)
	r.GET("/:id", a.get)
	r.POST("/:id", a.update)
}

func (a *userApi) get(ctx *gin.Context) {
	userID, err := tools.GetIntFromParams(ctx, "id")
	if err != nil {
		HandleError(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := db.UserModel(nil).Get(userID, true)
	if err != nil {
		HandleError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (a *userApi) getAll(ctx *gin.Context) {}

func (a *userApi) create(ctx *gin.Context) {
	actorID := uint(1)
	req := &userRequests.UserCreateRequest{}

	err := ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, http.StatusBadRequest)
		return
	}

	isPasswordValid := validators.ValidatePassword(req.User.Password)
	if !isPasswordValid {
		err = errors.New(constants.ERR_INVALID_PASSWORD_INSECURE)
		HandleError(ctx, err, http.StatusBadRequest)
		return
	}

	req.User.Profile = &req.Profile
	req.User.Status = constants.USER_STATUS_ACTIVE
	req.User.Password, err = tools.HashPassword(req.User.Password)
	if err != nil {
		HandleError(ctx, err, http.StatusInternalServerError)
		return
	}

	err = db.UserModel(nil).Create(actorID, &req.User)
	if err != nil {
		HandleError(ctx, err, http.StatusInternalServerError)
		return
	}

	user, err := db.UserModel(nil).GetByEmail(req.User.Email, false)
	if err != nil {
		HandleError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, userResponses.UserCreateResponse{ID: user.ID})
}

func (a *userApi) update(ctx *gin.Context) {
	actorID := uint(1)

	userID, err := tools.GetIntFromParams(ctx, "id")
	if err != nil {
		HandleError(ctx, err, http.StatusBadRequest)
		return
	}

	req := &userRequests.UserCreateRequest{}
	err = ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, http.StatusBadRequest)
		return
	}

	req.User.ID = userID

	userExisted, err := db.UserModel(nil).Get(userID, true)
	if err != nil {
		HandleError(ctx, err, http.StatusNotFound)
		return
	}

	req.User.Profile.ID = userExisted.ProfileID

	err = db.UserModel(nil).Update(actorID, &req.User)
	if err != nil {
		HandleError(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse{Success: true})
}
