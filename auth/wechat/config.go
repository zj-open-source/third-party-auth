package wechat

import "errors"

// openapi:strfmt wechat-mp-appid-secret
type WechatMpAppIDAndSecret struct {
	AppID     string
	AppSecret string
}

func (t WechatMpAppIDAndSecret) IsZero() bool {
	return t.AppID == "" || t.AppSecret == ""
}

// openapi:strfmt wechatwebsite-appid-secret
type WechatWebsiteAppIDAndSecret struct {
	AppID     string
	AppSecret string
}

const (
	WechatMiniProgramCode2Session = "/sns/jscode2session" // 登录凭证校验
)

var (
	WeChatSendError       = errors.New("访问微信接口网络错误")
	WeChatRespDecodeError = errors.New("微信返回参数解析失败")
)
