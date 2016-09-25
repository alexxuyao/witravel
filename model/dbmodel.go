package model

type User struct {
	Id         int64 `orm:"auto"`
	OpenId     string
	Nickname   string
	Sex        string
	Province   string
	City       string
	Country    string
	HeadImgUrl string
	Privilege  string
	UnionId    string
	Mobile     string
}
