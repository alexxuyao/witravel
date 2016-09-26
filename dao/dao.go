package dao

import (
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "witraveluser:witravelpasswd123@tcp(localhost:3306)/db_witravel?charset=utf8", 30)

	// register model
	orm.RegisterModelWithPrefix("tb_", new(model.User))

	// create table
	// orm.RunSyncdb("default", false, true)
}

func DoTransaction(fun func(ormer orm.Ormer) error) (err error) {
	o := orm.NewOrm()

	if err = o.Begin(); nil != err {
		return
	}

	defer func() {
		if rerr := recover(); nil != rerr {
			err = rerr.(error)

			if berr := o.Rollback(); nil != berr {
				// log the berr
			}
		} else {
			err = o.Commit()
		}
	}()

	if ferr := fun(o); nil != ferr {
		panic(ferr)
	}

	return
}
