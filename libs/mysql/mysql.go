package mysql

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db gorm.DB

func InitWithDsn(dns string) {
	open, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = *open
}

func NewSession(ctx context.Context) *gorm.DB {
	if ctx == nil {
		ctx = context.Background()
	}
	return db.Session(&gorm.Session{Context: ctx, NewDB: true})
}
