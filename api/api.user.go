package api

import (
	"errors"
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	"github.com/Irbi/citrouser/requests"
	"github.com/Irbi/citrouser/requests/userRequests"
	"github.com/Irbi/citrouser/responses"
	"github.com/Irbi/citrouser/responses/userResponses"
	"github.com/Irbi/citrouser/tools"
	"github.com/Irbi/citrouser/validators"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type userApi struct {}

func (a *userApi) Routes(r gin.IRoutes) {
	r.POST("", a.create)
	r.GET("", a.getAll)
	r.GET("/:id", a.get)
	r.PUT("/:id", a.update)
	r.PUT("/:id/activate", a.activate)
}

func (a *userApi) get(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "user",
		"event": "getUserByID",
	})

	userID, err := tools.GetIntFromParams(ctx, "id")
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	user, err := db.UserModel(nil).Get(userID, true)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
}

func (a *userApi) getAll(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "user",
		"event": "getUsers",
	})

	req := &requests.DefaultSearchRequest{}
	err := ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	usersList, totalCount, err := db.UserModel(nil).Find(req.Offset, req.Limit, req.Sort, req.Order, "")
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(usersList); i++ {
		usersList[i].Password = ""
	}

	ctx.JSON(http.StatusOK, userResponses.UsersListResponse{
		Users:      usersList,
		TotalCount: uint(totalCount),
	})
}

func (a *userApi) create(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "user",
		"event": "createUser",
	})

	actorID, err := GetUser(ctx, false)
	if HandleError(ctx, err, contextLogger, http.StatusUnauthorized) {
		return
	}

	req := &userRequests.UserCreateRequest{}
	err = ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	isPasswordValid := validators.ValidatePassword(req.User.Password)
	if !isPasswordValid {
		err = errors.New(constants.ERR_INVALID_PASSWORD_INSECURE)
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	req.User.Profile = &req.Profile
	req.User.Profile.PasswordLastUpdate = time.Now()
	req.User.Status = constants.USER_STATUS_ACTIVE
	req.User.Password, err = tools.HashPassword(req.User.Password)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	err = db.UserModel(nil).Create(actorID.ID, &req.User)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	user, err := db.UserModel(nil).GetByEmail(req.User.Email, false)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, userResponses.UserCreateResponse{ID: user.ID})
}

func (a *userApi) update(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "user",
		"event": "updateUser",
	})
	actorID, err := GetUser(ctx, false)
	if HandleError(ctx, err, contextLogger, http.StatusUnauthorized) {
		return
	}

	userID, err := tools.GetIntFromParams(ctx, "id")
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	req := &userRequests.UserCreateRequest{}
	err = ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	req.User.ID = userID

	userExisted, err := db.UserModel(nil).Get(userID, true)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusNotFound)
		return
	}

	req.User.Profile = &req.Profile
	req.User.Profile.ID = userExisted.ProfileID

	err = db.UserModel(nil).UpdateExcept(actorID.ID, &req.User, "password")
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	err = db.ProfileModel(nil).Update(actorID.ID, req.User.Profile)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse{Success: true})
}

func (a *userApi) activate(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "user",
		"event": "activateUser",
	})
	actorID, err := GetUser(ctx, false)
	if HandleError(ctx, err, contextLogger, http.StatusUnauthorized) {
		return
	}

	userID, err := tools.GetIntFromParams(ctx, "id")
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	err = db.UserModel(nil).UpdateStatus(actorID.ID, userID, constants.USER_STATUS_ACTIVE)
	if HandleError(ctx, err, contextLogger, http.StatusInternalServerError) {
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse{Success: true})
}