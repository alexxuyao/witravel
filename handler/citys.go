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

// 国家列表
func CountrysHandler(c *iris.Context) {
	countrys, err := GetCountrys()

	if nil != err {
		log.Errorln("GetCountrys error:", err)
		common.AjaxRespFail(c, nil)
		return
	}

	common.AjaxRespSuccess(c, countrys)
}

// 省份列表
func ProvincesHandler(c *iris.Context) {
	countryId, _ := strconv.ParseInt(c.Param("countryId"), 10, 64)
	provinces, err := GetProvinces(countryId)

	if nil != err {
		log.Errorln("GetProvinces error:", err)
		common.AjaxRespFail(c, nil)
		return
	}

	common.AjaxRespSuccess(c, provinces)
}

// 城市列表
func CitysHandler(c *iris.Context) {
	provinceId, _ := strconv.ParseInt(c.Param("provinceId"), 10, 64)
	citys, err := GetCitys(provinceId)

	if nil != err {
		log.Errorln("GetCitys error:", err)
		common.AjaxRespFail(c, nil)
		return
	}

	common.AjaxRespSuccess(c, citys)
}

// 国家列表
func GetCountrys() ([]model.Country, error) {
	var countrys []model.Country

	err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select * from tb_country"
		rawSeter := o.Raw(sql)
		num, err := rawSeter.QueryRows(&countrys)

		log.Infoln("query sql :", sql, ", return number:", num)

		if nil != err {
			return err
		}

		return nil
	})

	return countrys, err
}

// 省份列表
func GetProvinces(countryId int64) ([]model.Province, error) {
	var provinces []model.Province

	err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select * from tb_province where country_id = ?"
		rawSeter := o.Raw(sql, countryId)
		num, err := rawSeter.QueryRows(&provinces)

		log.Infoln("query sql :", sql, ", return number:", num)

		if nil != err {
			return err
		}

		return nil
	})

	return provinces, err
}

// 城市列表
func GetCitys(provinceId int64) ([]model.City, error) {
	var citys []model.City

	err := dao.DoTransaction(func(o orm.Ormer) error {

		sql := "select * from tb_city where province_id = ?"
		rawSeter := o.Raw(sql, provinceId)
		num, err := rawSeter.QueryRows(&citys)

		log.Infoln("query sql :", sql, ", return number:", num)

		if nil != err {
			return err
		}

		return nil
	})

	return citys, err
}
