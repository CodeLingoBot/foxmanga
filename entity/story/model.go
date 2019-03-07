package story

import (
	"github.com/jinzhu/gorm"
)

type Story struct {
	gorm.Model
	Title       string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	StoryFox    string `gorm:"type:text"`
}
