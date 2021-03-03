package model

type PortfolioHistory struct {
	BaseModel
	SourceID 		uint 				`json:"-"`
	Source 			*PortfolioSource 	`gorm:"foreignkey:SourceID;"`
	PortfolioID 	uint 				`json:"-"`
	Portfolio		*Portfolio			`gorm:"foreignkey:PortfolioID;"`
}
