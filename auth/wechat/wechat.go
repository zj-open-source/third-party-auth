package wechat

import (
	"context"
	"encoding/json"
	"github.com/go-courier/httptransport"
	"github.com/go-courier/httptransport/client"
	"github.com/go-courier/metax"
	"time"
)

type WechatOpts struct {
	Mp      WechatMpAppIDAndSecret      `env:""` // 小程序appid和密码
	Website WechatWebsiteAppIDAndSecret `env:""` // 微信网页端appid和密码
}

type WechatClient struct {
	*WechatOpts

	RequestTransformerMgr *httptransport.RequestTransformerMgr `env:"-"`

	metax.Ctx
}

func (c *WechatClient) LivenessCheck() map[string]string {
	m := map[string]string{
		"api.weixin.qq.com": "ok",
	}

	return m
}

func (c *WechatClient) WithContext(ctx context.Context) *WechatClient {
	return &WechatClient{
		WechatOpts:            c.WechatOpts,
		RequestTransformerMgr: c.RequestTransformerMgr,
		Ctx:                   c.Ctx.WithContext(ctx),
	}
}

func (WechatClient) BaseURL() string {
	return "https://api.weixin.qq.com"
}

func (c *WechatClient) SetDefaults() {
	if c.RequestTransformerMgr == nil {
		c.RequestTransformerMgr = httptransport.NewRequestTransformerMgr(nil, nil)
		c.SetDefaults()
	}
}

func (c *WechatClient) Available() bool {
	return !c.Mp.IsZero()
}

func (c *WechatClient) Do(id, method, url string, v interface{}, resp MayError, httpTransports ...client.HttpTransport) error {
	req, err := c.RequestTransformerMgr.NewRequest(method, c.BaseURL()+url, v)
	if err != nil {
		return err
	}

	req.Header.Add("X-Operation-Id", id)

	response, errForRequest := client.GetShortConnClientContext(c.Context(), 20*time.Second, httpTransports...).Do(req)
	if errForRequest != nil {
		return WeChatSendError
	}
	defer func() {
		_ = response.Body.Close()
	}()

	errForDecode := json.NewDecoder(response.Body).Decode(resp)
	if errForDecode != nil {
		return WeChatRespDecodeError
	}

	return resp.GetError()
}
