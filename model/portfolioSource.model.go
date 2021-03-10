package model

type PortfolioSource struct {
	BaseModel
	PortfolioID		uint 		`json:"-"`
	Portfolio		*Portfolio	`gorm:"foreignkey:PortfolioID;"`
}
