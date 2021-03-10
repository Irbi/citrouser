package db

import (
	"github.com/Irbi/citrouser/model"
	"gorm.io/gorm"
)

type clientAdvisorModel struct {
	db *gorm.DB
}

func ClientAdvisorModel(tx *gorm.DB) *clientAdvisorModel {
	if tx == nil {
		tx = Connection
	}
	return &clientAdvisorModel{db: tx}
}

func (m *clientAdvisorModel) ConnectClientToAdvisor(actorID uint, connection *model.ClientAdvisor) (err error) {
	connection.CreatedBy = actorID
	connection.UpdatedBy = actorID
	err = m.db.Create(connection).Error

	return
}

func (m *clientAdvisorModel) GetAdvisorByClient(clientID uint) (advisor *model.User, err error) {
	return nil, nil
}

