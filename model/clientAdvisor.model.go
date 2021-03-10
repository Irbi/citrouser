package model

type ClientAdvisor struct {
	BaseModel
	ClientUserID	uint	`json:"-"`
	ClientUser		*User	`gorm:"foreignkey:ClientUserID;"`
	AdvisorUserID	uint	`json:"-"`
	AdvisorUser		*User	`gorm:"foreignkey:AdvisorUserID;"`
}
