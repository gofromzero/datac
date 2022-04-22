// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/4/22

package gormcli

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}
