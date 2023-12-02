package dao

import (
	"gorm.io/gorm"

	"dnsgo/app/model"
)

var DnsRecord = &dao{}

type dao struct{}

func (d *dao) GetZoneDNSRecords(session *gorm.DB, id string) ([]*model.DNSRecord, error) {
	var records []*model.DNSRecord
	err := session.Where("zone_id = ?", id).Find(&records).Error
	return records, err
}

func (d *dao) Get(session *gorm.DB, id string) (*model.DNSRecord, error) {
	var record model.DNSRecord
	err := session.Where("id = ?", id).First(&record).Error
	return &record, err
}

func (d *dao) CreateDNSRecord(session *gorm.DB, record *model.DNSRecord) error {
	return session.Create(record).Error
}

func (d *dao) UpdateDNSRecord(session *gorm.DB, record *model.DNSRecord) error {
	return session.Save(record).Error
}

func (d *dao) DeleteDNSRecord(session *gorm.DB, id string) error {
	return session.Where("id = ?", id).Delete(&model.DNSRecord{}).Error
}

func (d *dao) DeleteDNSRecordByZoneId(session *gorm.DB, zoneID string) error {
	return session.Where("zone_id = ?", zoneID).Delete(&model.DNSRecord{}).Error
}

func (d *dao) FindByZoneName(session *gorm.DB, zoneName string) ([]*model.DNSRecord, error) {
	var records []*model.DNSRecord
	err := session.Where("zone_name = ?", zoneName).Find(&records).Error
	return records, err
}

func (d *dao) FindByZoneNameAndType(session *gorm.DB, zoneName string, recordType string) ([]*model.DNSRecord, error) {
	var records []*model.DNSRecord
	query := session.Where("zone_name = ?", zoneName)
	if recordType != "any" {
		query = query.Where("type = ?", recordType)
	}
	err := query.Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (d *dao) GetDNSRecordByZoneIdAndName(session *gorm.DB, zoneID, name string) (*model.DNSRecord, error) {
	var record model.DNSRecord
	err := session.Where("zone_id = ? AND name = ?", zoneID, name).First(&record).Error
	return &record, err
}

func (d *dao) GetDNSRecordByZoneIdAndNameAndType(session *gorm.DB, zoneID, name, recordType string) (*model.DNSRecord, error) {
	var record model.DNSRecord
	err := session.Where("zone_id = ? AND name = ? AND type = ?", zoneID, name, recordType).First(&record).Error
	return &record, err
}
