package service

import (
	"context"

	"dnsgo/app/dao"
	"dnsgo/app/model"
	"dnsgo/libs/mysql"
)

var DNSRecord = &dnsRecordSvc{}

type dnsRecordSvc struct{}

func (s *dnsRecordSvc) GetZoneDNSRecords(ctx context.Context, id string) ([]*model.DNSRecord, error) {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.GetZoneDNSRecords(session, id)
}

func (s *dnsRecordSvc) Get(ctx context.Context, id string) (*model.DNSRecord, error) {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.Get(session, id)
}

func (s *dnsRecordSvc) CreateDNSRecord(ctx context.Context, record *model.DNSRecord) error {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.CreateDNSRecord(session, record)
}

func (s *dnsRecordSvc) UpdateDNSRecord(ctx context.Context, record *model.DNSRecord) error {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.UpdateDNSRecord(session, record)
}

func (s *dnsRecordSvc) DeleteDNSRecord(ctx context.Context, id string) error {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.DeleteDNSRecord(session, id)
}

func (s *dnsRecordSvc) DeleteDNSRecordByZoneId(ctx context.Context, zoneID string) error {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.DeleteDNSRecordByZoneId(session, zoneID)
}

func (s *dnsRecordSvc) GetDNSRecordByZoneIdAndName(ctx context.Context, zoneID, name string) (*model.DNSRecord, error) {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.GetDNSRecordByZoneIdAndName(session, zoneID, name)
}

func (s *dnsRecordSvc) GetDNSRecordByZoneIdAndNameAndType(ctx context.Context, zoneID, name, recordType string) (*model.DNSRecord, error) {
	session := mysql.NewSession(ctx)
	return dao.DnsRecord.GetDNSRecordByZoneIdAndNameAndType(session, zoneID, name, recordType)
}
