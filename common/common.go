package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/alexxuyao/witravel/model"
	"github.com/kataras/iris"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicWechatError(resp []byte) {
	wresp := model.WechatResp{}
	err := json.Unmarshal(resp, &wresp)

	if nil != err {
		panic(err)
	}

	if wresp.ErrCode != 0 {
		panic(errors.New(wresp.ErrMsg))
	}
}

func GetWechatError(resp []byte) error {
	wresp := model.WechatResp{}
	err := json.Unmarshal(resp, &wresp)

	if nil != err {
		return err
	}

	if wresp.ErrCode != 0 {
		return errors.New(wresp.ErrMsg)
	}

	return nil
}

func DoGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func WechatGet(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	wresp := model.WechatResp{}
	err = json.Unmarshal(body, &wresp)

	if nil != err {
		return err
	}

	if wresp.ErrCode != 0 {
		return errors.New(wresp.ErrMsg)
	}

	return nil
}

func WechatPost(url string, data []byte) error {

	buf := bytes.NewBuffer(data)
	resp, err := http.Post(url, "application/json", buf)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	wresp := model.WechatResp{}
	err = json.Unmarshal(body, &wresp)

	if nil != err {
		return err
	}

	if wresp.ErrCode != 0 {
		return errors.New(wresp.ErrMsg)
	}

	return nil
}

// ajax response
type AjaxJsonResp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func AjaxRespSuccess(c *iris.Context, data interface{}) {
	resp := AjaxJsonResp{Success: true, Data: data}
	c.JSON(0, resp)
}

func AjaxRespFail(c *iris.Context, data interface{}) {
	resp := AjaxJsonResp{Success: false, Data: data}
	c.JSON(0, resp)
}

func ErrorAjaxResp(c *iris.Context) {
	if err := recover(); err != nil {
		log.Errorln(err)
		AjaxRespFail(c, nil)
	}
}

// 首字母小写
func LowerFirst(word string) string {
	bword := []byte(word)
	return strings.ToLower(string(bword[0:1])) + string(bword[1:])
}
