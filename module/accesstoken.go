package module

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type TokenResp struct {
	AccessToken string `json:"access_token"` // 令牌
	ExpiresIn   int64  `json:"expires_in"`   // 过期时间, 秒
}

type AccessTokenModule struct {
	AppId       string
	AppSecret   string
	accessToken string // 令牌
	expireTime  int64  // 过期时间的unix时间戳, 秒
	lock        sync.Mutex
}

// 取得token
func (me *AccessTokenModule) GetAccessToken() string {

	me.reGetAccessToken()

	return me.accessToken
}

func (me *AccessTokenModule) reGetAccessToken() {

	if me.isExpired() {

		me.lock.Lock()
		defer me.lock.Unlock()

		// 再次检查，因为在等待锁的过程中，可能别人已经取了一次token
		if me.isExpired() {

			log.Println("access token expired, start go new")

			resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + me.AppId + "&secret=" + me.AppSecret)
			if err != nil {
				// handle error

			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			tokenResp := TokenResp{}
			err = json.Unmarshal(body, &tokenResp)

			if err != nil {

			}

			me.accessToken = tokenResp.AccessToken
			me.expireTime = time.Now().Unix() + tokenResp.ExpiresIn
		}
	}
}

// token是否已过期
func (me *AccessTokenModule) isExpired() bool {
	// 过期时间减去当前时间小于30秒
	return me.expireTime-time.Now().Unix() < 30
}
