package wechat

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWechatClient_Do(t *testing.T) {
	t.Run("#code2session", func(t *testing.T) {
		c := WechatClient{}
		c.SetDefaults()
		req := Code2SessionRequest{
			AppID:     appID,
			Secret:    secret,
			GrantType: "authorization_code",
			Code:      code,
		}

		resp := struct {
			WechatError
			Code2SessionResp
		}{}

		if err := c.Do("Code2Session", http.MethodGet, WechatMiniProgramCode2Session, req, &resp); err != nil {
			fmt.Printf("get session by code failure:[err:%v]\n", err)
			return
		}
		fmt.Printf("wechat response:%+v", resp.Code2SessionResp)
	})

	t.Run("#ExchangeSessionByCode", func(t *testing.T) {
		c := WechatClient{}
		c.WechatOpts = &WechatOpts{
			Mp: WechatMpAppIDAndSecret{
				AppID:     appID,
				AppSecret: secret,
			},
		}
		c.SetDefaults()

		resp, err := c.ExchangeSessionByCode(code)
		fmt.Printf("wechat response:%v,err:%v", resp, err)
	})

}

const (
	appID  = "appid"
	secret = "secret"
	code   = "091jAFGa1Z0QYB0wTAFa17l1sE4jAFGL"
)
