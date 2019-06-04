package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model
	Project    Project
	StartTime  time.Time
	FinishTime time.Time
	Result     string
}
