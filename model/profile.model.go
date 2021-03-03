package model

import "time"

type Profile struct {
	BaseModel
	FirstName 			string 		`json:"firstName" gorm:"type:varchar(255);" binding:"required"`
	LastName 			string 		`json:"lastName" gorm:"type:varchar(255);" binding:"required"`
	Phone 				string 		`json:"phone" gorm:"type:varchar(255);"`
	PasswordLastUpdate 	*time.Time 	`json:"passwordLastUpdate"`
}
