package model

type ClientAdvisor struct {
	BaseModel
	ClientID	uint	`json:"-"`
	Client		*User	`gorm:"foreignkey:UserID;"`
	AdvisorID	uint	`json:"-"`
	Advisor		*User	`gorm:"foreignkey:UserID;"`
}
