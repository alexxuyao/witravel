package handler

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/alexxuyao/witravel/common"
	"github.com/alexxuyao/witravel/dao"
	"github.com/alexxuyao/witravel/model"
	"github.com/alexxuyao/witravel/module"
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris"
)

type ValidateResp struct {
	Validate bool              `json:"validate"`
	Msg      map[string]string `json:"msg"`
}

// 验证数据
func TravelValidateHandler(c *iris.Context) {

	log.Println(string(c.PostBody()))

	travel := model.Travel{}
	err := json.Unmarshal(c.PostBody(), &travel)

	if nil != err {
		log.Errorln("TravelValidateHandler error:", err)
		common.AjaxRespFail(c, nil)
		return
	}

	result, msgs := TravelValidate(travel)

	common.AjaxRespSuccess(c, ValidateResp{Validate: result, Msg: msgs})
}

// 行程数据校验
func TravelValidate(travel model.Travel) (bool, map[string]string) {
	ret, keys := common.ValidateStruct(travel)

	if travel.SponsorMobile == "" && travel.SponsorWechat == "" {
		keys["sponsorMobile"] = "SponsorMobile or SponsorWechat cannot be empty"
		ret = false
	}

	return ret, keys
}

func SaveTravelHandler(c *iris.Context) {
	travel := model.Travel{}
	err := json.Unmarshal(c.PostBody(), &travel)

	if nil != err {
		log.Errorln("SaveTravelHandler error:", err)
		common.AjaxRespFail(c, nil)
		return
	}

	if travel.Id == 0 {
		AddTravel(c, travel)
	} else {

	}
}

// 新建行程
func AddTravel(c *iris.Context, travel model.Travel) {

	result, msgs := TravelValidate(travel)

	// 数据校验失败
	if !result {
		common.AjaxRespFail(c, ValidateResp{Validate: result, Msg: msgs})
		return
	}

	webuser := c.Get("webuser").(*module.WebUser)

	travel.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	travel.Sponsor = webuser.Get().Id
	travel.Status = model.TRAVEL_STATUS_SUCCESS
	travel.ParticipantsNumber = 1
	travel.Visitors = 0
	travel.ImgUrl = webuser.Get().HeadImgUrl

	err := dao.DoTransaction(func(o orm.Ormer) error {

		_, err := o.Insert(&travel)

		if nil != err {
			return err
		}

		return nil
	})

	if nil != err {
		common.AjaxRespFail(c, err.Error())
		return
	}

	common.AjaxRespSuccess(c, nil)
}
