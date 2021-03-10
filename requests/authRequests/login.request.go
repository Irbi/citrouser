package authRequests

type LoginRequest struct {
	Email		string	`json:"email" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

type AuthResetPasswordRequest struct {
	Email 		string	`json:"email" binding:"required"`
}

type AuthVerifyResetPasswordCode struct {
	ResetPasswordCode string `json:"resetPasswordCode" binding:"required"`
	Password          string `json:"password" binding:"required"`
}
