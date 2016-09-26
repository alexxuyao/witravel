package handler

import (
	"fmt"
	"testing"
)

func Test_travellist(t *testing.T) {
	l := NextTravelList("ss", "ss", 22)
	fmt.Println(l)
}
