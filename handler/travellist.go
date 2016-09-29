package handler

import (
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/alexxuyao/witravel/common"
	"github.com/alexxuyao/witravel/dao"
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris"
)

// 行程列表
func TravelListHandler(c *iris.Context) {

	rtype := c.FormValueString("type")
	date := c.FormValueString("date")
	destination := c.FormValueString("destination")
	meetingTime := c.FormValueString("meetingTime")
	lastId, err := strconv.ParseInt(c.FormValueString("lastId"), 10, 64)

	log.Println("lastId is :", lastId)

	if nil != err {
		lastId = 0
	}

	var list []model.Travel

	if "prev" == rtype {
		list = PrevTravelList(date, destination, meetingTime, lastId)
	} else {
		list = NextTravelList(date, destination, meetingTime, lastId)
	}

	common.AjaxRespSuccess(c, list)
}

// 下一页
func NextTravelList(date, destination, meetingTime string, lastId int64) []model.Travel {
	var travels []model.Travel

	if err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select * from tb_travel where id < ? and meeting_time < ?"
		args := []interface{}{lastId, meetingTime}

		if date != "" {
			sql += " and date = ?"
			args = append(args, date)
		}

		if destination != "" {
			sql += " and destination = ?"
			args = append(args, destination)
		}

		sql += " order by meeting_time desc, id desc limit 10"

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
func PrevTravelList(date, destination, meetingTime string, lastId int64) []model.Travel {

	return nil
}
