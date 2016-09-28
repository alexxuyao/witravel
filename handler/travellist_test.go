package handler

import (
	"testing"

	"github.com/alexxuyao/witravel/model"
)

func Test_travellist(t *testing.T) {
	//	l := NextTravelList("ss", "ss", 22)
	//	fmt.Println(l)
}

func Test_validate(t *testing.T) {

	//	govalidator.SetFieldsRequiredByDefault(true)

	//	travel := model.Travel{Destination: "新德里", Budget: 1000, MeetingCity: 123, MeetingPlace: "dddccc"}

	//	result, err := govalidator.ValidateStruct(travel)
	//	keys := make(map[string]string)
	//	if err != nil {
	//		println("error: " + err.Error())
	//		for _, v := range strings.Split(err.Error(), ";") {
	//			println(v)
	//			if v != "" {
	//				pair := strings.Split(v, ":")

	//				keys[pair[0]] = pair[1]
	//			}
	//		}
	//	}
	//	println(result)
	//	println(keys)
	travel := model.Travel{Destination: "新德里", Budget: 1000, MeetingCity: 123, MeetingPlace: "dddccc", SponsorMobile: "100000"}
	i, k := TravelValidate(travel)

	println(i)
	for key, val := range k {
		println(key)
		println(val)

	}

}
