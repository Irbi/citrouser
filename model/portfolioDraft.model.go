package model

type PortfolioDraft struct {
	BaseModel
	SourceID 		uint 				`json:"-"`
	Source 			*PortfolioSource 	`gorm:"foreignkey:SourceID;"`
	UserID 			uint 				`json:"-"`
	User			*User				`gorm:"foreignkey:UserID;"`
	PortfolioID 	uint 				`json:"-"`
	Portfolio		*Portfolio			`gorm:"foreignkey:PortfolioID;"`
}
