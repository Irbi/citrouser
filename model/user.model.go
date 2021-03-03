package model

type User struct {
	BaseModel
	Email 			string 		`json:"email" gorm:"type:varchar(255);unique_index;not null;" binding:"required"`
	Password 		string  	`json:"-" gorm:"type:varchar(255);"`
	Role 			string 		`json:"role" gorm:"type:varchar(255);" binding:"required"`
	Status 			string 		`json:"status" gorm:"type:varchar(255);"`
	ProfileID 		uint		`json:"-"`
	Profile 		*Profile 	`gorm:"foreignkey:ProfileID;"`
	PermissionID	uint		`json:"-"`
	Permission		*Permission	`gorm:"foreignkey:PermissionID;"`
}
