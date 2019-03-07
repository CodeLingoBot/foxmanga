package bootstrap

import (
	"fmt"
	"github.com/foxmanga/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func init() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		config.Configuration.Database.User,
		config.Configuration.Database.Password,
		config.Configuration.Database.Name))
	if err != nil {

	}
	defer db.Close()
}
