package dao

import (
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("witravel", "mysql", "witraveluser:witravelpasswd123@tcp(localhost:3306)/my_db?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(model.User))

	// create table
	// orm.RunSyncdb("default", false, true)
}
