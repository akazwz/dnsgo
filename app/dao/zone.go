package dao

import (
	"gorm.io/gorm"

	"dnsgo/app/model"
)

var Zone = &zoneDao{}

type zoneDao struct{}

func (d *zoneDao) Find(session *gorm.DB) ([]*model.Zone, error) {
	var zones = make([]*model.Zone, 0)
	err := session.Find(&zones).Error
	return zones, err
}

func (d *zoneDao) Get(session *gorm.DB, id string) (*model.Zone, error) {
	var zone = &model.Zone{}
	err := session.Where("id = ?", id).First(&zone).Error
	return zone, err
}

func (d *zoneDao) Create(session *gorm.DB, zone *model.Zone) error {
	return session.Create(zone).Error
}

func (d *zoneDao) Update(session *gorm.DB, zone *model.Zone) error {
	return session.Save(zone).Error
}

func (d *zoneDao) Delete(session *gorm.DB, id string) error {
	return session.Where("id = ?", id).Delete(&model.Zone{}).Error
}
