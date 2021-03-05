package auth

type LoginResponse struct {
	ID 		uint 	`json:"ID"`
	Email	string	`json:"email"`
	Status	string	`json:"status"`
	Role 	string	`json:"role"`
}
