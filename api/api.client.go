package api

import (
	"errors"
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	"github.com/Irbi/citrouser/model"
	"github.com/Irbi/citrouser/requests/clientRequests"
	"github.com/Irbi/citrouser/responses"
	"github.com/Irbi/citrouser/tools"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type clientApi struct {}

func (a *clientApi) Routes(r gin.IRoutes) {
	r.POST("/:id/connect", a.connect)
}

func (a *clientApi) connect(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api": "connectApi",
		"event": "connectClientToAdvisor",
	})

	actorID, err := GetUser(ctx, false)
	if HandleError(ctx, err, contextLogger, http.StatusUnauthorized) {
		return
	}

	req := &clientRequests.ConnectClientRequest{}
	err = ctx.Bind(req)
	if HandleError(ctx, err, contextLogger, http.StatusBadRequest) {
		return
	}

	clientID, err := tools.GetIntFromParams(ctx, "id")
	if HandleError(ctx, err, contextLogger, http.StatusBadRequest) {
		return
	}

	client, err := db.UserModel(nil).Get(clientID, false)
	if HandleError(ctx, err, contextLogger, http.StatusInternalServerError) {
		return
	}

	clientAdvisor, err := db.ClientAdvisorModel(nil).GetAdvisorByClient(clientID)
	if HandleError(ctx, err, contextLogger, http.StatusInternalServerError) {
		return
	}
	if clientAdvisor != nil {
		HandleError(ctx, errors.New(constants.ERR_CLIENT_ADVISOR_EXISTS), contextLogger, http.StatusConflict)
		return
	}

	advisor, err := db.UserModel(nil).Get(req.AdvisorID, false)
	if HandleError(ctx, err, contextLogger, http.StatusInternalServerError) {
		return
	}

	connection := &model.ClientAdvisor{
		ClientUser: client,
		AdvisorUser: advisor,
	}

	err = db.ClientAdvisorModel(nil).ConnectClientToAdvisor(actorID.ID, connection)
	if HandleError(ctx, err, contextLogger, http.StatusInternalServerError) {
		return
	}

	ctx.JSON(http.StatusNoContent, responses.SuccessResponse{Success:true})
}

func (a *clientApi) checkActorPermissions(actor *model.User, clientID uint) (status int, err error) {
	return http.StatusOK, nil
}