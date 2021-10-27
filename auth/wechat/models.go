package wechat

import (
	"errors"
	"fmt"
)

type MayError interface {
	GetError() error
}

type WechatError struct {
	ErrCode int    `json:"errcode,omitempty"` // 错误码
	ErrMsg  string `json:"errmsg,omitempty"`  // 错误信息
}

func (err WechatError) GetError() error {
	if err.ErrCode != 0 {
		return errors.New(err.Error())
	}
	return nil
}

func (err WechatError) Error() string {
	return fmt.Sprintf("code: %d msg: %s", err.ErrCode, err.ErrMsg)
}

type WechatUser struct {
	// 普通用户的标识，对当前开发者帐号唯一
	OpenID string `json:"openid"`
	// 普通用户昵称
	Nickname string `json:"nickname"`
	// 普通用户性别，1 为男性，2 为女性
	Sex int `json:"sex"`
	// 普通用户个人资料填写的省份
	Province string `json:"province"`
	// 普通用户个人资料填写的城市
	City string `json:"city"`
	// 国家，如中国为CN
	Country string `json:"country"`
	// 	用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
	HeadImgUrl string `json:"headimgurl"`
	// 用户特权信息，json数组，如微信沃卡用户为（chinaunicom）
	Privilege []string `json:"privilege"`
	// 用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的 unionid 是唯一
	UnionID string `json:"unionid,omitempty"`
}

// region WeChat Mini Program:code2session

type Code2SessionRequest struct {
	AppID     string `name:"appid" in:"query"`
	Secret    string `name:"secret" in:"query"`
	GrantType string `name:"grant_type" in:"query"`
	Code      string `name:"js_code" in:"query"`
}

type Code2SessionResp struct {
	// 用户唯一标识
	OpenID string `json:"openid"`
	// 会话密钥
	SessionKey string `json:"session_key"`
	// 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回
	UnionID string `json:"unionid,omitempty"`
}

// endregion WeChat Mini Program:code2session

// region WXUser

// WxUserInfo 微信用户信息
type WxUserInfo struct {
	OpenID    string `json:"openId"`
	UnionID   string `json:"unionId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	Language  string `json:"language"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

// endregion WXUser

// WxUserPhoneInfo 微信用户信息
type WxUserPhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string `json:"countryCode"`     // 区号
	Watermark       struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

// endregion WxUserPhoneInfo
