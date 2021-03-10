package model

type PortfolioHistoryModel struct {
	ID 					uint				`json:"ID" gorm:"primary_key"`
	SourceID			uint				`json:"-"`
	Source 				*PortfolioSource	`gorm:"foreignkey:PortfolioSourceID;"`
	ContentInitialID	uint				`json:"-"`
	ContentInitial		*PortfolioContent	`gorm:"foreignkey:ContentInitialID;"`
	ContentFinalID		uint				`json:"-"`
	ContentFinal		*PortfolioContent	`gorm:"foreignkey:ContentFinalID;"`
	Diff				string				`gorm:"text;"`
}
