package db

import (
	"github.com/Irbi/citrouser/model"
	"gorm.io/gorm"
)

type portfolioModel struct {
	db *gorm.DB
}

func PortfolioModel(tx *gorm.DB) *portfolioModel {
	if tx == nil {
		tx = Connection
	}
	return &portfolioModel{db: tx}
}

func (m *portfolioModel) Create(actorID uint, portfolio *model.Portfolio) (err error) {
	portfolio.CreatedBy = actorID
	portfolio.UpdatedBy = actorID

	err = m.db.Create(portfolio).Error

	return
}


