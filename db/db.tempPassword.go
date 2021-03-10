package db

import (
	"github.com/Irbi/citrouser/model"
	"gorm.io/gorm"
)

type tempPassModel struct {
	db *gorm.DB
}

func TempPassModel(tx *gorm.DB) *tempPassModel {
	if tx == nil {
		tx = Connection
	}
	return &tempPassModel{db: tx}
}

func (m *tempPassModel) Create(tempPass *model.TempPassword) (err error) {
	err = m.db.Create(tempPass).Error
	return
}

func (m *tempPassModel) DeleteByUser(userId uint) (err error) {
	err = m.db.Where("user_id = ?", userId).Delete(&model.TempPassword{}).Error
	return
}

func (m *tempPassModel) Get(tmpId string, preloadUser bool) (tempPass *model.TempPassword, err error) {
	tempPass = &model.TempPassword{}
	v := m.db.Model(&model.TempPassword{})

	if preloadUser {
		v = v.Preload("User")
	}
	err = v.Where("ID = ?", tmpId).Find(tempPass).Error

	return
}


