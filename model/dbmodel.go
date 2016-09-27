package model

// 用户
type User struct {
	Id         int64 `orm:"auto"` // 用户ID
	OpenId     string
	Nickname   string
	Sex        int
	Province   string
	City       string
	Country    string
	HeadImgUrl string
	Privilege  string
	UnionId    string
	Mobile     string
}

// 行程
type Travel struct {
	Id                 int64  `orm:"auto" ` // 主键
	Destination        string // 目的地
	MeetingTime        string // 集合时间(北京时间)
	MeetingCountry     string // 集合点国家
	MeetingProvince    string // 集合点省份
	MeetingCity        string // 集合点城市
	MeetingPlace       string // 集合地点详细地址
	ReturnDate         string // 返程日期
	Sponsor            int64  // 发起人Id
	CreateTime         string // 发起时间(北京时间)
	ParticipantsNumber int    // 已参加人数
	Status             int    // 状态
	Title              string // 标题
	Description        string // 详细说明
	SponsorMobile      string // 发起人手机号(手机号及微信号必须填写一个,只有参与人能看到手机及微信号)
	SponsorWechat      string // 发起人微信号
	Visitors           int    // 浏览人数
	Budget             int    // 预算(人民币)
	ImgUrl             string // 头像
}

// 行程参与者
type TravelParticipants struct {
	Id           int64  `orm:"auto" ` // 主键
	TravelId     int64  // 行程ID
	Participants int64  // 参与人 Id
	createTime   string // 参与时间
	Province     string // 参与人所在省份
	City         string // 参与人所在城市
	Mobile       string // 手机号
	Wechat       string // 微信号
}

// 行程浏览记录
type TravelViewRecord struct {
	Id       int64  `orm:"auto" ` // 主键
	UserId   string // 查看人Id
	ViewTime string // 查看时间
	TravelId int64  // 行程ID
	// 查看时所在地理位置
}

// 行程留言
type TravelComments struct {
	Id int64 `orm:"auto" ` // 主键
}

// 导游表
type TourGuide struct {
	Id          int64  `orm:"auto" ` // 主键
	UserId      int64  // 用户Id， 可为0
	Realname    string // 姓名
	Sex         int    // 性别
	Birth       string // 出生年月
	HeadImgUrl  string // 头像
	Nationality string // 国籍
	Country     string // 所在国家
	Province    string // 所在省份
	City        string // 所在城市
	Mobile      string // 手机
	Wechat      string // 微信号
	Description string // 说明
	Status      int    // 状态
}

// 药店
type Pharmacy struct {
	Id int64 `orm:"auto" ` // 主键
}

// 国家
type Country struct {
	Id   int64  `orm:"auto" json:"id"` // 主键
	Name string `json:"name"`          // 名称
}

// 省份
type Province struct {
	Id        int64  `orm:"auto" json:"id"` // 主键
	CountryId int64  `json:"countryId"`     // 国家id
	Name      string `json:"name"`          // 名称
}

// 城市
type City struct {
	Id         int64  `orm:"auto" json:"id"` // 主键
	CountryId  int64  `json:"countryId"`     // 国家id
	ProvinceId int64  `json:"provinceId"`    // 省份id
	Name       string `json:"name"`          // 名称
}

// 区
type County struct {
	Id         int64  `orm:"auto" json:"id"` // 主键
	CountryId  int64  `json:"countryId"`     // 国家id
	ProvinceId int64  `json:"provinceId"`    // 省份id
	CityId     int64  `json:"cityId"`        // 城市id
	Name       string `json:"name"`          // 名称
}
