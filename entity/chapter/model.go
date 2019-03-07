package chapter

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/pgtype/ext/satori-uuid"
	"github.com/jinzhu/gorm"
)

type Chapter struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255)"`
	FoxID     uuid.UUID
	Page      string
	View      int32
	Date      timestamp.Timestamp
	IsPublish bool
}
