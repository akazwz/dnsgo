package model

import "time"

type QueryLog struct {
	ID                string    `json:"id" gorm:"column:id;type:varchar(255);primary_key;not null"`
	ZoneID            string    `json:"zone_id" gorm:"column:zone_id;type:varchar(255);not null"`
	QueryType         string    `json:"query_type" gorm:"column:query_type;type:varchar(255);not null"`
	QueryName         string    `json:"query_name" gorm:"column:query_name;type:varchar(255);not null"`
	RemoteIP          string    `json:"remote_ip" gorm:"column:remote_ip;type:varchar(255);not null"`
	RemoteCountry     string    `json:"remote_country" gorm:"column:remote_country;type:varchar(255);not null"`
	RemoteCountryCode string    `json:"remote_country_code" gorm:"column:remote_country_code;type:varchar(255);not null"`
	ServerIP          string    `json:"server_ip" gorm:"column:server_ip;type:varchar(255);not null"`
	Rcode             string    `json:"rcode" gorm:"column:rcode;type:varchar(255);not null"`
	Content           string    `json:"content" gorm:"column:content;type:varchar(255);not null"`
	Cost              int64     `json:"cost" gorm:"column:cost;type:bigint;not null"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
}

func (QueryLog) TableName() string {
	return "query_logs"
}
