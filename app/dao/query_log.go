package dao

import (
	"gorm.io/gorm"

	"dnsgo/app/model"
)

var QueryLog = &queryLogDao{}

type queryLogDao struct{}

func (d *queryLogDao) Create(session *gorm.DB, log *model.QueryLog) (*model.QueryLog, error) {
	err := session.Create(log).Error
	return log, err
}

func (d *queryLogDao) List(session *gorm.DB, page, pageSize int) ([]*model.QueryLog, error) {
	var logs []*model.QueryLog
	err := session.Offset((page - 1) * pageSize).Limit(pageSize).Order("created_at DESC").Find(&logs).Error
	return logs, err
}
