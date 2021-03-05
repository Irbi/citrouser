package userResponses

import "github.com/Irbi/citrouser/model"

type UsersListResponse struct {
	Users 		[]model.User	`json:"list"`
	TotalCount	uint			`json:"totalCount"`
}
