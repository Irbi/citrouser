package model

type PortfolioSource struct {
	Source string `json:"source" gorm:"varchar(255);"`
}
