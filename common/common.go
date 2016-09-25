package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/alexxuyao/witravel/model"
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
