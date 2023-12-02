package service

import (
	"context"

	"dnsgo/app/dao"
	"dnsgo/app/model"
	"dnsgo/libs/mysql"
)

var Zone = &zoneSvc{}

type zoneSvc struct{}

func (s *zoneSvc) Find(ctx context.Context) ([]*model.Zone, error) {
	session := mysql.NewSession(ctx)
	return dao.Zone.Find(session)
}

func (s *zoneSvc) Get(ctx context.Context, id string) (*model.Zone, error) {
	session := mysql.NewSession(ctx)
	return dao.Zone.Get(session, id)
}

func (s *zoneSvc) Create(ctx context.Context, zone *model.Zone) error {
	session := mysql.NewSession(ctx)
	return dao.Zone.Create(session, zone)
}

func (s *zoneSvc) Update(ctx context.Context, zone *model.Zone) error {
	session := mysql.NewSession(ctx)
	return dao.Zone.Update(session, zone)
}

func (s *zoneSvc) Delete(ctx context.Context, id string) error {
	session := mysql.NewSession(ctx)
	return dao.Zone.Delete(session, id)
}
