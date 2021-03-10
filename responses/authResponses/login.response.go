package authResponses

type LoginResponse struct {
	ID 			uint 	`json:"ID"`
	Token   	string	`json:"token"`
	Expire		string	`json:"expire"`
	Email		string	`json:"email"`
	UserStatus	string	`json:"userStatus"`
	UserRole 	string	`json:"userRole"`
}

type CheckToken struct {
	ID    		string 	`json:"id"`
	Email 		string 	`json:"email"`
}
