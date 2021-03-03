package userRequests

import "github.com/Irbi/citrouser/model"

type UserCreateRequest struct  {
	User 	model.User 		`json:"user"`
	Profile model.Profile 	`json:"profile"`
}
