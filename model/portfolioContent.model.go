package model

type PortfolioContent struct {
	ID 					uint						`json:"ID" gorm:"primary_key"`
	SourceID			uint						`json:"-"`
	Source 				*PortfolioSource			`gorm:"foreignkey:PortfolioSourceID;"`
	AssetCategoryID		uint						`json:"-"`
	AssetCategory		*FinancialAssetCategory		`gorm:"foreignkey:FinancialAssetCategoryID;"`
	AssetID				uint						`json:"-"`
	Asset				*FinancialAsset				`gorm:"foreignkey:FinancialAssetID;"`
	InstrumentID		uint						`json:"-"`
	Instrument			*FinancialInstrument		`gorm:"foreignkey:FinancialInstrumentID;"`
	ISIN 				string						`gorm:"varchar(255);"`
	UnitsNumber			string						`gorm:"varchar(255);"`
	Value				string						`gorm:"varchar(255);"`
	Fee					uint						`gorm:"int;"`
}
