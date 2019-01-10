package bootstrap


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var database *xorm.Engine

func init() {
	var err error
	database, err = xorm.NewEngine("mysql", "hieppp:123456@tcp(localhost:3306)/foxmanga?charset=utf8")
	if err != nil{
		panic(err)
	}
}