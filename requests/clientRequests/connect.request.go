package clientRequests

type ConnectClientRequest struct {
	AdvisorID 	uint `json:"advisorID" binding:"required"`
}
