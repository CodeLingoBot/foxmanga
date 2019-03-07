package genre

import (
	"github.com/jinzhu/gorm"
)

type Story struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)"`
}
