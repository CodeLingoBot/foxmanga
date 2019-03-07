package fox

import (
	"github.com/foxmanga/entity/chapter"
	"github.com/foxmanga/entity/site"
	"github.com/jackc/pgx/pgtype/ext/satori-uuid"
	"github.com/jinzhu/gorm"
)

type Fox struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255)"`
	Tag           string `gorm:"type:varchar(255)"`
	Rating        float32
	Completed     bool
	TotalChapter  int32
	Author        string
	OtherName     string
	PublishedYear string
	Banner        string
	Description   string
	MostPopular   bool
	Site          site.Site `gorm:"foreignkey:SiteID"`
	SiteID        uuid.UUID
	Chapter []chapter.Chapter `gorm:"foreignkey:FoxID"`
}
