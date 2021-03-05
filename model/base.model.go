package model

import "time"

type BaseModel struct {
	ID 			uint 		`json:"ID" gorm:"primary_key"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	CreatedBy 	uint		`json:"createdBy"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	UpdatedBy 	uint 		`json:"updatedBy"`
	DeletedAt 	*time.Time 	`json:"deletedAt" gorm:"default:null"`
	DeletedBy 	*uint		`json:"deletedBy" gorm:"default:null"`
}
