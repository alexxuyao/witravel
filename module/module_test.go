package module

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"text/template"
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
