package cache

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"dnsgo/app/dao"
	"dnsgo/app/model"
	"dnsgo/libs/mysql"
	"dnsgo/libs/redis"
)

var DNSRecord = &dnsRecordCache{}

type dnsRecordCache struct{}

func (c *dnsRecordCache) Get(ctx context.Context, id string) (*model.DNSRecord, error) {
	client := redis.GetClient()
	data, err := client.Get(ctx, "dns_record:"+id).Result()
	if err == nil {
		record := &model.DNSRecord{}
		err = record.FromJSON(data)
		if err == nil {
			return record, nil
		}
	}
	session := mysql.NewSession(ctx)
	record, err := dao.DnsRecord.Get(session, id)
	if err != nil {
		return nil, err
	}
	jsonStr, err := record.ToJSON()
	if err == nil {
		client.Set(ctx, "dns_record:"+id, jsonStr, time.Minute*60)
	}
	return record, nil
}

func (c *dnsRecordCache) FindByZoneName(ctx context.Context, zoneName string) ([]*model.DNSRecord, error) {
	client := redis.GetClient()
	data, err := client.Get(ctx, "dns_record:zone_name:"+zoneName).Result()
	if err == nil {
		records := make([]*model.DNSRecord, 0)
		if data == "[]" {
			return records, nil
		}
		err = json.Unmarshal([]byte(data), &records)
		if err == nil {
			return records, nil
		} else {
			log.Println("err: ", err)
		}
	}

	session := mysql.NewSession(ctx)
	records, err := dao.DnsRecord.FindByZoneName(session, zoneName)
	if err == nil {
		jsonStr, err := json.Marshal(records)
		if err == nil {
			client.Set(ctx, "dns_record:zone_name:"+zoneName, jsonStr, time.Minute*60)
		}
	}

	return records, err
}

func (c *dnsRecordCache) FindByZoneNameAndType(ctx context.Context, zoneName string, recordType string) ([]*model.DNSRecord, error) {
	client := redis.GetClient()
	data, err := client.Get(ctx, "dns_record:zone_name:"+zoneName+":type:"+recordType).Result()
	if err == nil {
		records := make([]*model.DNSRecord, 0)
		if data == "[]" {
			return records, nil
		}
		err = json.Unmarshal([]byte(data), &records)
		if err == nil {
			return records, nil
		} else {
			log.Println("err: ", err)
		}
	}

	session := mysql.NewSession(ctx)
	records, err := dao.DnsRecord.FindByZoneNameAndType(session, zoneName, recordType)
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			records = []*model.DNSRecord{}
			client.Set(ctx, "dns_record:zone_name:"+zoneName+":type:"+recordType, "[]", time.Minute*60)
			return records, nil
		}
		jsonStr, err := json.Marshal(records)
		if err == nil {
			client.Set(ctx, "dns_record:zone_name:"+zoneName+":type:"+recordType, jsonStr, time.Minute*60)
		}
	}

	return records, err
}
