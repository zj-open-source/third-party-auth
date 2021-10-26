package wechat

import (
	"bytes"
	"net/http"
	"net/url"
)

const (
	AuthURL = "https://api.weixin.qq.com/sns/jscode2session"
)

func (c *WechatClient) AuthCodeURLByMP(state string) *url.URL {
	var buf bytes.Buffer
	buf.WriteString(AuthURL)
	v := url.Values{
		"appid":         {c.Mp.AppID},
		"response_type": {"code"},
		"scope":         {"snsapi_userinfo"},
		"state":         {state},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	buf.WriteString("#wechat_redirect")
	u, _ := url.Parse(buf.String())
	return u
}

// ExchangeSessionByCode code置换session
func (c *WechatClient) ExchangeSessionByCode(code string) (*Code2SessionResp, error) {
	req := Code2SessionRequest{
		AppID:     c.Mp.AppID,
		Secret:    c.Mp.AppSecret,
		GrantType: "authorization_code",
		Code:      code,
	}

	resp := struct {
		WechatError
		Code2SessionResp
	}{}

	if err := c.Do("ExchangeSessionByCode", http.MethodGet, WechatMiniProgramCode2Session, req, &resp); err != nil {
		return nil, err
	}

	return &resp.Code2SessionResp, resp.GetError()
}

func (c *WechatClient) getUserInfoByMP(accessToken string, openid string) (*WechatUser, error) {
	req := struct {
		AccessToken string `name:"access_token" in:"query"`
		OpenID      string `name:"openid" in:"query"`
		Lang        string `name:"lang" in:"query"`
	}{}

	req.AccessToken = accessToken
	req.OpenID = openid
	req.Lang = "zh-cn"

	resp := &struct {
		WechatError
		WechatUser
	}{}

	if err := c.Do("UserInfo", "GET", "/sns/userinfo", req, resp); err != nil {
		return nil, err
	}

	return &resp.WechatUser, nil
}
