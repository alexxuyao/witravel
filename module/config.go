package module

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type WechatConfig struct {
	AppId     string `json:"appID"`
	AppSecret string `json:"appsecret"`
}

type AppConfig struct {
	Wechat           WechatConfig `json:"wechat"`
	DomainAndSubPath string       `json:"domainAndSubPath"`
	ImgDir           string       `json:"imgDir"`
	IsDebug          bool         `json:"isDebug"`
}

type ConfigModule struct {
	config *AppConfig
	once   sync.Once
}

func (me *ConfigModule) readConfig() {
	filepath := "./config.json"
	r, err := os.Open(filepath)
	defer r.Close()

	if nil != err {
		log.Println("open config file error,", err)
		return
	}

	config := AppConfig{}

	bs, _ := ioutil.ReadAll(r)
	json.Unmarshal(bs, &config)

	me.config = &config

	ret, _ := json.Marshal(config)
	log.Println(string(ret))
}

func (me *ConfigModule) GetConfig() *AppConfig {
	me.once.Do(me.readConfig)

	return me.config
}
