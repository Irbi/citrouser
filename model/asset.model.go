package model

type FinancialAsset struct {
	ID 							uint 						`json:"ID" gorm:"primary_key"`
	Name 						string  					`json:"name" gorm:"type:varchar(255);not null;"`
	SystemName 					string  					`json:"sysName" gorm:"type:varchar(255);uniqueIndex;not null;"`
	Description					string						`json:"description" gorm:"type:varchar(255);"`
	FinancialAssetCategoryID	uint						`json:"-"`
	FinancialAssetCategory		*FinancialAssetCategory		`gorm:"foreignkey:FinancialAssetCategoryID;"`
}
