package model

import "time"

type TempPassword struct {
	ID             string 	`gorm:"primary_key" json:"requestID"`
	UserID         uint
	User           User      `gorm:"foreignkey:UserID;"`
	TimeExpiration time.Time `json:"timeCreated"`
}
