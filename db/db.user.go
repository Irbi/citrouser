package db

import (
	"github.com/Irbi/citrouser/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

var defaultSort = "created_at"
var defaultOrder = "ASC"

type userModel struct {
	db *gorm.DB
}

func UserModel(tx *gorm.DB) *userModel {

	log.Debugf("Connection: %v", Connection)

	if tx == nil {
		tx = Connection
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

func (m *userModel) UpdateExcept(actorID uint, user *model.User, omitFields... string) (err error) {
	user.UpdatedBy = actorID
	err = m.db.Model(&model.User{}).Where("id = ?", user.ID).Omit(omitFields...).Updates(user).Error

	return
}

func (m *userModel) UpdateOnly(actorID uint, userID uint, updateFields... interface{}) (err error) {
	err = m.db.Model(&model.User{}).
		Where("id = ?", userID).
		Update("updated_by", actorID).
		Updates(updateFields).Error

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

func (m *userModel) Find(offset, limit, sortBy, sortOrder, role interface{}) (users []model.User, totalCount int64, err error) {

	v := m.db.Model(&model.User{}).Count(&totalCount)

	if sortBy == "" {
		sortBy = defaultSort
	}
	if sortOrder == "" {
		sortOrder = defaultOrder
	}

	intOffset, err := strconv.Atoi(offset.(string))
	if err != nil {
		return nil, 0, err
	}

	intLimit, err := strconv.Atoi(limit.(string))
	if err != nil {
		return nil, 0, err
	}

	v = v.Preload("Profile")

	v = v.Order(sortBy.(string) + " " + sortOrder.(string))
	v = v.Offset(intOffset)
	v = v.Limit(intLimit)
	if role != "" {
		v = v.Where("role = ?", role)
	}

	users = []model.User{}
	err = v.Find(&users).Error

	return
}