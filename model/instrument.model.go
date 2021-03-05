package model

type FinancialInstrument struct {
	ID 					uint 				`json:"ID" gorm:"primary_key"`
	Name 				string  			`json:"name" gorm:"type:varchar(255);not null;"`
	SystemName 			string  			`json:"sysName" gorm:"type:varchar(255);uniqueIndex;not null;"`
	FinancialAssetID	uint 				`json:"-"`
	FinancialAsset		*FinancialAsset 	`gorm:"foreignkey:FinancialAssetID;"`
	Description			string				`json:"description" gorm:"type:varchar(255);"`
}
