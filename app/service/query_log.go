package service

import (
	"context"

	"dnsgo/app/dao"
	"dnsgo/app/model"
	"dnsgo/libs/mysql"
)

var QueryLog = &queryLogSvc{}

type queryLogSvc struct{}

func (s *queryLogSvc) Create(ctx context.Context, log *model.QueryLog) (*model.QueryLog, error) {
	session := mysql.NewSession(ctx)
	return dao.QueryLog.Create(session, log)
}

func (s *queryLogSvc) List(ctx context.Context, page, pageSize int) ([]*model.QueryLog, error) {
	session := mysql.NewSession(ctx)
	return dao.QueryLog.List(session, page, pageSize)
}
