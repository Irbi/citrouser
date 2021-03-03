package db

import (
	"github.com/Irbi/citrouser/model"
	"gorm.io/gorm"
)

type userModel struct {
	db *gorm.DB
}

func UserModel(tx *gorm.DB) *userModel {
	if tx == nil {
		tx = &gorm.DB{}
	}
	return &userModel{db: tx}
}

func (m *userModel) Create(actorID uint, user *model.User) (err error) {
	user.CreatedBy = actorID
	user.UpdatedBy = actorID

	err = m.db.Create(user).Error

	return
}

func (m *userModel) Update(actorID uint, user *model.User) (err error) {
	user.UpdatedBy = actorID

	err = m.db.Model(&model.User{}).Updates(user).Error

	return
}

func (m *userModel) Get(id uint, preloadProfile bool) (user *model.User, err error) {
	user = &model.User{}
	v := m.db.Model(&model.User{})

	if preloadProfile {
		v = v.Preload("Profile")
	}

	err = v.Where("ID = ?", id).Find(user).Error

	return
}

func (m *userModel) GetByEmail(email string, preloadProfile bool) (user *model.User, err error) {
	user = &model.User{}
	v := m.db.Model(&model.User{})

	if preloadProfile {
		v = v.Preload("Profile")
	}

	err = v.Where("Email = ?", email).Find(user).Error

	return
}