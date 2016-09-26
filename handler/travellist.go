package handler

import (
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/alexxuyao/witravel/dao"
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris"
)

// 用于微信提交认证
func TravelListHandler(c *iris.Context) {

	rtype := c.FormValueString("type")
	date := c.FormValueString("date")
	destination := c.FormValueString("destination")
	lastId, err := strconv.ParseInt(c.FormValueString("lastId"), 10, 64)

	if nil != err {
		lastId = 0
	}

	if "prev" == rtype {
		PrevTravelList(date, destination, lastId)
	} else {
		NextTravelList(date, destination, lastId)
	}

	c.JSON(0, struct{}{})
}

// 下一页
func NextTravelList(date, destination string, lastId int64) []model.Travel {
	var travels []model.Travel

	if err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select id from tb_user where id > ? "
		args := []interface{}{lastId}

		if date != "" {
			sql += " and date = ?"
			args = append(args, date)
		}

		if destination != "" {
			sql += " and destination = ?"
			args = append(args, destination)
		}

		sql += " limit 10"

		log.Infoln(sql)
		log.Infoln(args)

		rawSeter := o.Raw(sql, args...)
		num, err := rawSeter.QueryRows(&travels)

		log.Infoln("query sql :", sql, ", return number:", num)

		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		log.Errorln("NextTravelList got an error, ", err, err.Error())
	}

	return travels
}

// 上一页(刷新数据)
func PrevTravelList(date, destination string, lastId int64) []model.User {

	return nil
}
