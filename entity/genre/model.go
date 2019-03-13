package genre

import (
	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)"`
}
