package authRequests

type RegisterRequest struct {
	Email		string	`json:"email" binding:"required"`
	Password	string	`json:"password" binding:"required"`
	Phone		string	`json:"phone" binding:"required"`
	Role		string 	`json:"role" binding:"required"`
}
