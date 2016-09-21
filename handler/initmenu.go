package handler

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/alexxuyao/witravel/common"
	"github.com/alexxuyao/witravel/module"
	"github.com/kataras/iris"
)

// 初始化菜单
func InitMenuHandler(c *iris.Context) {
	container := c.Get("container").(*module.WebContainer)
	accessToken := container.Token.GetAccessToken()

	// 读取菜单
	menu, err := readMenu()
	if nil != err {
		log.Println("readMenu error,", err)
	}

	tpl, err := template.New("menu").Parse(string(menu))
	if nil != err {
		log.Println("tpl Parse error,", err)
	}

	buf := bytes.NewBuffer([]byte{})
	err = tpl.Execute(buf, container.Config.DomainAndSubPath)
	if nil != err {
		log.Println("tpl Execute error,", err)
	}

	// 先删除原菜单
	err = deleteMenu(accessToken)
	if nil != err {
		log.Println("deleteMenu error,", err)
	}

	// 创建新菜单
	err = createMenu(buf.Bytes(), accessToken)
	if nil != err {
		log.Println("createMenu error,", err)
	}
}

func readMenu() ([]byte, error) {
	filepath := "./menu.json"
	r, err := os.Open(filepath)
	defer r.Close()

	if nil != err {
		log.Println("open menu file error,", err)
		return nil, err
	}

	return ioutil.ReadAll(r)
}

func deleteMenu(accessToken string) error {
	return common.WechatGet("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=" + accessToken)
}

func createMenu(menuJson []byte, accessToken string) error {
	return common.WechatPost("https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+accessToken, menuJson)
}
