package bootstrap

import (
	"fmt"
	"github.com/foxmanga/config"
	"github.com/foxmanga/entity/chapter"
	"github.com/foxmanga/entity/fox"
	"github.com/foxmanga/entity/genre"
	"github.com/foxmanga/entity/site"
	"github.com/foxmanga/entity/story"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

func init() {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%d@/%s?charset=utf8&parseTime=True&loc=Local",
			config.Configuration.Database.User,
			config.Configuration.Database.Password,
			config.Configuration.Database.Name))
	if err != nil {

	}
	defer db.Close()

	db.AutoMigrate(&site.Site{}, &genre.Genre{}, &fox.Fox{}, &chapter.Chapter{}, &story.Story{})
}
