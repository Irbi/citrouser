package userRequests

import (
	"github.com/Irbi/citrouser/requests"
)

type GetPortfoliusByUser struct {
	requests.DefaultSearchRequest
	UserID	uint `json:"userId"`
}