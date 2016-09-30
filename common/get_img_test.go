package common

import (
	"fmt"
	"testing"
)

func TestRandomFilename(t *testing.T) {
	f := RandomFilename("jpg")
	f1 := RandomFilename("jpg")
	fmt.Println(f)
	fmt.Println(f1)
}

func TestGetImg(t *testing.T) {
	f, err := GetImg("http://pic26.nipic.com/20130123/10558908_130511870000_2.jpg", ".")
	fmt.Println("get img finish")
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(f)
}
