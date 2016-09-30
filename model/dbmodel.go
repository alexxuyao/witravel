package model

const (
	TRAVEL_STATUS_COMMIT  int = 0 // 提交，未审核
	TRAVEL_STATUS_SUCCESS int = 1 // 审核成功
	TRAVEL_STATUS_FAIL    int = 2 // 审核失败
)

// 用户
type User struct {
	Id         int64  `orm:"auto" json:"id"` // 用户ID
	OpenId     string `json:"openId"`        // 用户OpenID
	Nickname   string `json:"nickname"`      // 用户名
	Sex        int    `json:"sex"`           // 性别
	Province   string `json:"province"`      // 省份
	City       string `json:"city"`          // 城市
	Country    string `json:"country"`       // 国家
	HeadImgUrl string `json:"headImgUrl"`    // 头像链接
	Privilege  string `json:"privilege"`     // 权限
	UnionId    string `json:"unionId"`       // 用户UnionID
	Mobile     string `json:"mobile"`        // 手机号
}

// 行程
type Travel struct {
	Id                 int64  `orm:"auto" json:"id" `                              // 主键
	Destination        string `json:"destination" validate:"strlen=255, nonzero"`  // 目的地
	MeetingTime        string `json:"meetingTime" validate:"nonzero"`              // 集合时间(北京时间)
	MeetingCountry     int64  `json:"meetingCountry" validate:"nonzero"`           // 集合点国家
	MeetingProvince    int64  `json:"meetingProvince" validate:"nonzero"`          // 集合点省份
	MeetingCity        int64  `json:"meetingCity" validate:"nonzero"`              // 集合点城市
	MeetingPlace       string `json:"meetingPlace" validate:"strlen=255, nonzero"` // 集合地点详细地址
	ReturnDate         string `json:"returnDate" validate:"nonzero"`               // 返程日期
	Sponsor            int64  `json:"sponsor" `                                    // 发起人Id
	CreateTime         string `json:"createTime" `                                 // 发起时间(北京时间)
	ParticipantsNumber int    `json:"participantsNumber" `                         // 已参加人数
	Status             int    `json:"status" `                                     // 状态
	Description        string `json:"description" validate:"strlen=255"`           // 详细说明
	SponsorMobile      string `json:"sponsorMobile" validate:"strlen=11"`          // 发起人手机号(手机号及微信号必须填写一个,只有参与人能看到手机及微信号)
	SponsorWechat      string `json:"sponsorWechat" validate:"strlen=255"`         // 发起人微信号
	Visitors           int    `json:"visitors" `                                   // 浏览人数
	Budget             int    `json:"budget" validate:"nonzero"`                   // 预算(人民币)
	ImgUrl             string `json:"imgUrl" validate:"strlen=512"`                // 头像
}

// 行程参与者
type TravelParticipants struct {
	Id           int64  `orm:"auto" json:"id" ` // 主键
	TravelId     int64  `json:"travelId"`       // 行程ID
	Participants int64  `json:"participants"`   // 参与人 Id
	CreateTime   string `json:"createTime"`     // 参与时间
	CountryId    int64  `json:"countryId"`      // 国家ID
	ProvinceId   int64  `json:"provinceId"`     // 参与人所在省份ID
	CityId       int64  `json:"cityId"`         // 参与人所在城市ID
	Country      string `json:"country"`        // 国家
	Province     string `json:"province"`       // 参与人所在省份
	City         string `json:"city"`           // 参与人所在城市
	Mobile       string `json:"mobile"`         // 手机号
	Wechat       string `json:"wechat"`         // 微信号
	Nickname     string `json:"nickname"`       // 用户名
	Sex          int    `json:"sex"`            // 性别
	HeadImgUrl   string `json:"headImgUrl"`     // 头像链接
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
