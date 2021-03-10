package api

import (
	"errors"
	"fmt"
	"github.com/Irbi/citrouser/constants"
	"github.com/Irbi/citrouser/db"
	"github.com/Irbi/citrouser/mail"
	"github.com/Irbi/citrouser/middleware"
	"github.com/Irbi/citrouser/model"
	"github.com/Irbi/citrouser/requests/authRequests"
	"github.com/Irbi/citrouser/responses"
	"github.com/Irbi/citrouser/responses/authResponses"
	"github.com/Irbi/citrouser/tools"
	"github.com/Irbi/citrouser/validators"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type authApi struct {
	Auth *middleware.Auth
}

const IDENTITY_KEY = "UserID"
const IDENTITY_EMAIL = "Email"

func (a *authApi) Routes(r gin.IRoutes) {
	a.Auth.SetAuthenticator(a.Login)
	a.Auth.SetAuthorizator(a.IsAccessAllowed)
	a.Auth.SetLoginResponse(a.LoginResponse)

	loginHandler := a.Auth.GetLoginHandler()
	refreshHandler := a.Auth.GetRefreshHandler()

	r.POST("/register", a.register)

	r.POST("/login", loginHandler)
	r.GET("/refresh_token", refreshHandler)
	r.POST("/reset-password", a.resetPassword)
	r.PUT("/reset-password-verify-code", a.resetPasswordVerifyCode)

	rr := r.Use(a.Auth.Jwt.MiddlewareFunc())
	rr.GET("check-token", a.checkToken)
}

func (a *authApi) Login(ctx *gin.Context) (interface{}, error) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "auth",
		"event": "login",
	})

	req := &authRequests.LoginRequest{}
	err := ctx.Bind(req)
	if err != nil {
		return nil, err
	}

	user, err := db.UserModel(nil).GetByEmail(req.Email, false)
	if err != nil || !tools.CheckPasswordFromHash(req.Password, user.Password){
		return nil, errors.New(constants.ERR_USER_INVALID_CREDENTIALS)
	}

	contextLogger.Debugf("UserLogin. User: %s", user.Email)

	if user.Status != constants.USER_STATUS_ACTIVE {
		return nil, errors.New(constants.ERR_USER_INACTIVE)
	}

	ctx.Set("userRole", user.Role)
	ctx.Set("userStatus", user.Status)

	ctx.Set(IDENTITY_KEY, strconv.Itoa(int(user.ID)))
	ctx.Set(IDENTITY_EMAIL, user.Email)

	return &model.User{Email: user.Email, Status: user.Status}, nil
}

func (a *authApi) IsAccessAllowed(data interface{}, ctx *gin.Context) bool {
	if userData, ok := data.(*model.User); ok && userData.Email != "" {

		userExists, err := db.UserModel(nil).GetByEmail(userData.Email, false)
		if err == nil && userExists.Status == constants.USER_STATUS_ACTIVE {
			return true
		}
	}

	return false
}

func (a *authApi) LoginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	if code == http.StatusOK {
		userRole, _ := ctx.Get("userRole")
		userStatus, _ := ctx.Get("userStatus")
		id, _ := strconv.Atoi(ctx.GetString(IDENTITY_KEY))

		ctx.JSON(http.StatusOK, authResponses.LoginResponse{
			ID:              uint(id),
			Email:           ctx.GetString(IDENTITY_EMAIL),
			Token:           token,
			Expire:          expire.Format(time.RFC3339),
			UserRole:        userRole.(string),
			UserStatus:      userStatus.(string),
		})

		return

	}  else {
		HandleError(ctx, errors.New(constants.ERR_AUTH_CANNOT_AUTHORIZE), nil, http.StatusInternalServerError)
		return
	}
}

func (a *authApi) checkToken(ctx *gin.Context) {
	user, err := GetUser(ctx, false)
	if HandleError(ctx, err, nil, http.StatusBadRequest) {
		return
	}

	ctx.JSON(http.StatusOK, authResponses.CheckToken{
		Email: user.Email,
		ID:    strconv.Itoa(int(user.ID)),
	})
}


func (a *authApi) register(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "auth",
		"event": "registerUser",
	})

	req := &authRequests.RegisterRequest{}
	err := ctx.Bind(req)

	if HandleError(ctx, err, contextLogger, http.StatusBadRequest) {
		return
	}

	isPhoneValid := validators.ValidatePhoneNumber(req.Phone)
	if !isPhoneValid {
		err = errors.New(constants.ERR_INVALID_PHONE_NUMBER)
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}
	isPasswordValid := validators.ValidatePassword(req.Password)
	if !isPasswordValid {
		err = errors.New(constants.ERR_INVALID_PASSWORD_INSECURE)
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}
	req.Password, err = tools.HashPassword(req.Password)
	if HandleError(ctx, err, contextLogger, http.StatusInternalServerError) {
		return
	}

	userProfile := &model.Profile{
		Phone: req.Phone,
		PasswordLastUpdate: time.Now(),
	}
	user := &model.User{
		Email: req.Email,
		Password: req.Password,
		Role: req.Role,
		Status: constants.USER_STATUS_AWAIT_APPROVE,
		Profile: userProfile,
	}

	log.Printf("Register user: %v", user)

	err = db.UserModel(nil).Create(0, user)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse{Success:true})
}

func (a *authApi) resetPassword(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "auth",
		"event": "resetPassword",
	})

	req := &authRequests.AuthResetPasswordRequest{}
	err := ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
	}

	user, err := db.UserModel(nil).GetByEmail(req.Email, false)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusInternalServerError)
		return
	}

	if user.Status != constants.USER_STATUS_ACTIVE {
		HandleError(ctx, errors.New(constants.ERR_USER_INACTIVE), contextLogger, http.StatusInternalServerError)
	}

	code := tools.GenerateUUID()
	tempPassID := tools.GetMD5Hash(code)
	tempPass := &model.TempPassword {
		ID:             tempPassID,
		UserID:         user.ID,
		TimeExpiration: time.Now().UTC().Add(time.Hour * constants.RESET_PASSWORD_LINK_EXPIRATION),
	}

	contextLogger.Debugf("Verification code: %s\nRequestID: %s\nExpiration: %s", code, tempPassID, tempPass.TimeExpiration)

	err = db.TempPassModel(nil).Create(tempPass)
	if err != nil {
		HandleError(ctx, errors.New(constants.ERR_AUTH_CANNOT_CREATE_TEMPORARY_PASSWORD), contextLogger, http.StatusInternalServerError)
		return
	}

	resetLink := fmt.Sprintf("%s%s?code=%s", os.Getenv("WEB_URL"), "/reset-password-verify-code", code)
	contextLogger.Debugf("Reset link: %s", resetLink)
	err = mail.Mail.ResetPassword(user.Email, user.Email, resetLink)
	if err != nil {
		HandleError(ctx, errors.New(constants.ERR_FAILED_SEND_MAIL), contextLogger, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, &responses.SuccessResponse{Success: true})
}

func (a *authApi) resetPasswordVerifyCode(ctx *gin.Context) {
	contextLogger := log.WithFields(log.Fields{
		"api":   "auth",
		"event": "resetPasswordVerifyCode",
	})

	req := &authRequests.AuthVerifyResetPasswordCode{}
	err := ctx.Bind(req)
	if err != nil {
		HandleError(ctx, err, contextLogger, http.StatusBadRequest)
		return
	}

	if req.ResetPasswordCode == "" {
		HandleError(ctx, errors.New(constants.ERR_RESET_PASSWORD_INVALID_CODE), contextLogger, http.StatusBadRequest)
		return
	}
	tempPass := tools.GetMD5Hash(req.ResetPasswordCode)
	changePasswordRequest, err := db.TempPassModel(nil).Get(tempPass, true)
	if err != nil {
		HandleError(ctx, errors.New(constants.ERR_RESET_CANNOT_RECREATE_PASSWORD), contextLogger, http.StatusInternalServerError)
		return
	}
	if changePasswordRequest.User.Status != constants.USER_STATUS_ACTIVE {
		HandleError(ctx, errors.New(constants.ERR_USER_INACTIVE), contextLogger, http.StatusBadRequest)
		return
	}

	if changePasswordRequest.TimeExpiration.Before(time.Now().UTC()) {
		HandleError(ctx, errors.New(constants.ERR_RESET_PASSWORD_EXPIRED_CODE), contextLogger, http.StatusGone)
		return
	}

	status, err := a.changePassword(&changePasswordRequest.User, req.Password)
	if err != nil {
		HandleError(ctx, err, contextLogger, status)
		return
	}
	_ = db.TempPassModel(nil).DeleteByUser(changePasswordRequest.User.ID)

	ctx.JSON(http.StatusNoContent, nil)
}

func (a *authApi) changePassword(user *model.User, pwd string) (status int, err error) {
	password, err := tools.HashPassword(pwd)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = db.UserModel(nil).UpdatePassword(user.ID, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = db.UserModel(nil).UpdateStatus(user.ID, user.ID, constants.USER_STATUS_ACTIVE)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return
}
