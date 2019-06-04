package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	client *gorm.DB
}

func NewDbConnection(dbUri string) (*Database, error) {
	c, err := gorm.Open("postgres")
	if err != nil {
		return nil, err
	}
	return &Database{
		client: c,
	}, nil
}
