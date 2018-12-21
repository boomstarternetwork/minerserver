package dbstore

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type DBStore struct {
	gdb *gorm.DB
}

func New(connStr string, runMode string) (*DBStore, error) {
	gdb, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("failed to open gorm DB: " + err.Error())
	}

	switch runMode {
	case "production", "testing":
		gdb.LogMode(false)
	case "development":
		gdb.LogMode(true)
	default:
		return nil, errors.New("invalid mode")
	}

	return &DBStore{gdb: gdb}, nil
}
