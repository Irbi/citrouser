package db

import (
	"github.com/Irbi/citrouser/model"
	"gorm.io/gorm"
)

type profileModel struct {
	db *gorm.DB
}

func ProfileModel(tx *gorm.DB) *profileModel {
	if tx == nil {
		tx = Connection
	}
	return &profileModel{db: tx}
}

func (m *profileModel) Update(actorID uint, profile *model.Profile) (err error) {
	profile.UpdatedBy = actorID
	err = m.db.Model(&model.Profile{}).Where("id = ?", profile.ID).Updates(profile).Error

	return
}
