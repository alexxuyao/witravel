package module

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"text/template"

	"github.com/alexxuyao/witravel/dao"
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"
)

func Test_http(t *testing.T) {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET")
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Test_lock(t *testing.T) {
	tpl, err := template.New("menu").Parse("hello, {{.}}")
	if nil != err {
		fmt.Println(err)
	}

	buf := bytes.NewBuffer([]byte{})
	err = tpl.Execute(buf, "alex")
	fmt.Println(string(buf.Bytes()))
}

func Test_string(t *testing.T) {
	m := make(map[string]string)
	fmt.Println("get the key of alex is,", m["alex"])

	if m["alex"] == "" {
		fmt.Println("en. it's empty")
	}

}

func Test_read(t *testing.T) {
	if err := dao.DoTransaction(func(o orm.Ormer) error {
		u := model.User{OpenId: "12"}
		err := o.Read(&u, "OpenId")
		if nil != err {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println("user.id is ", u.Id)
		return nil
	}); nil != err {
		fmt.Println("error is ", err)
	}

	//	fmt.Println("the return is ," + defTest())
}

func defTest() string {
	defer func() string {
		if err := recover(); nil != err {
			fmt.Println(err.(error).Error())
		}
		fmt.Println("get the error .")
		return "return from defer....."
	}()

	panic(errors.New("this is the error"))

	return "return not from defer"
}
