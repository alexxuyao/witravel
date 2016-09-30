package common

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetImg(url, dir string) (string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(resp.Body)
	fname := dir + "/" + RandomFilename("jpg")

	f, err := os.Create(fname)
	defer f.Close()

	if nil != err {
		return "", err
	}

	_, err = f.Write(content)

	if nil != err {
		return "", err
	}

	return fname, err
}

func RandomFilename(postfix string) string {
	//	now := time.Now().Unix()
	//	random := rand.New(rand.NewSource(now))

	return strconv.Itoa(int(time.Now().UnixNano())) + "." + postfix
}
