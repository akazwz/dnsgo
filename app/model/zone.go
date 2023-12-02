package model

import (
	"encoding/json"
	"time"
)

type Zone struct {
	ID        string    `json:"id" gorm:"column:id;type:varchar(255);primary_key;not null"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Comment   string    `json:"comment" gorm:"column:comment;type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null"`
}

func (*Zone) TableName() string {
	return "zones"
}

func (z *Zone) FromJSON(jsonStr string) error {
	err := json.Unmarshal([]byte(jsonStr), z)
	if err != nil {
		return err
	}
	return nil
}

func (z *Zone) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(z)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
