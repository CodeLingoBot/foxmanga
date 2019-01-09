package bootstrap


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "hieppp:123456@tcp(localhost:3306)/slocker?charset=utf8")
	if err != nil{
		panic(err)
	}
}