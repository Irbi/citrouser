package requests

type DefaultSearchRequest struct {
	Offset 	string	`form:"offset" binding:"required"`
	Limit 	string	`form:"limit" binding:"required"`
	Order 	string	`form:"order"`
	Sort 	string	`form:"sort"`
}
