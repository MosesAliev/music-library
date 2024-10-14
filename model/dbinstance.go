package model

import "gorm.io/gorm"

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance
