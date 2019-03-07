package site

import (
	"github.com/jinzhu/gorm"
)

type Site struct {
	gorm.Model
	URL  string `gorm:"type:varchar(100);"`
	Name string `gorm:"type:varchar(100);"`
}
