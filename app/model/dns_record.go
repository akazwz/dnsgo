package model

import (
	"encoding/json"
	"time"
)

type DNSRecord struct {
	ID        string    `json:"id" gorm:"column:id;type:varchar(255);primary_key;not null"`
	ZoneID    string    `json:"zone_id" gorm:"column:zone_id;type:varchar(255);not null"`
	ZoneName  string    `json:"zone_name" gorm:"column:zone_name;type:varchar(255);not null"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Comment   string    `json:"comment" gorm:"column:comment;type:varchar(255);not null"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled;type:tinyint(1);not null"`
	Type      string    `json:"type" gorm:"column:type;type:varchar(255);not null"`
	Content   string    `json:"content" gorm:"column:content;type:varchar(255);not null"`
	TTL       int       `json:"ttl" gorm:"column:ttl;type:int(11);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null"`
}

func (*DNSRecord) TableName() string {
	return "dns_records"
}

func (z *DNSRecord) FromJSON(jsonStr string) error {
	err := json.Unmarshal([]byte(jsonStr), z)
	if err != nil {
		return err
	}
	return nil
}

func (z *DNSRecord) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(z)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
