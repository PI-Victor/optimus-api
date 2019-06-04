package database

import (
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDbConnection(dbUri string) (*gorm.DB, error) {
	c, err := gorm.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}
	logrus.Info("Successfully connected to the database!")
	return c, nil
}
