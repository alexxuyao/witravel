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

// 行程详情
func TravelDetailHandler(c *iris.Context) {

	defer common.ErrorAjaxResp(c)

	travelId, err := strconv.Atoi(c.Param("travelId"))
	common.CheckError(err)

	travel, err := GetTravel(travelId)
	common.CheckError(err)

	participants, err := GetTravelParticipants(travelId)
	common.CheckError(err)

	result := make(map[string]interface{})
	result["travel"] = travel
	result["participants"] = participants

	common.AjaxRespSuccess(c, result)
}

func GetTravel(travelId int) (model.Travel, error) {

	var travel model.Travel

	err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select * from tb_travel where id = ?"

		rawSeter := o.Raw(sql, travelId)
		return rawSeter.QueryRow(&travel)
	})

	return travel, err
}

func GetTravelParticipants(travelId int) ([]model.TravelParticipants, error) {

	var participants []model.TravelParticipants

	err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select * from tb_travel_participants where travel_id = ?"

		rawSeter := o.Raw(sql, travelId)
		num, err := rawSeter.QueryRows(&participants)

		log.Infoln("query sql :", sql, ", return number:", num)

		return err
	})

	return participants, err
}
