package wechat

import (
	"fmt"
	"testing"
)

// 微信数据解密测试用例
func TestWXUserDataCrypt_Decrypt(t *testing.T) {
	t.Run("#wchat decrypt", func(t *testing.T) {
		sessionKey := `ehTXf9j8W/sUa/gmBAltNw==`
		encryptedData := "agMDx+X2QqRHtNy/fjHnRYlhq0/f74B0HuvGYKk7yp7Yps2Q/8jsDsiGJG2f1EYUiP1tM3awHmek0qjibBG2nb7oXEeaXyibS2/+R6TeiPrjyp06KLgqc4suqibC+hvaia9qX8G1PJmp4jyImZRu31DKSzAV2U04tOENnthRYZ2D2TgLIsTS8sCofbDSozU+8ifgrHgzXMr9UgBYaAExMZYHYCZvO/IKG4zeCOslJt0RySGVCssVQRuniAza/Lrk/h5xAGneXNUyxVcM0n8DjmNA+P7s0Ep+t1ASEqDb8D8J0PDODJwr7ebRey4RjlzqN2cZ3ad/fkgL8xiE/uYu8uyF6d3K1NGoSaR6SwkDinpsKmKWIOTq1Nb0E3L6XlYy83O5qFhDr1+1FB/sGdWeLhEuGtwprMY188iCCSTFg7koKGFc0NXcjs7Ad3DyZTR1SyHzfQJhKM5FbP4B5xcUNlDFL46WuDIHgUyg3AwJaHISuHTKWIR3wil9ULIziabSKG3Pl/5fK3/GIgiSgfTEVQ=="
		iv := "pGLBiY1oB/ZTgW/LbCgPQg=="
		pc := NewWXUserDataCrypt(appID, sessionKey)
		userInfo, err := pc.Decrypt(encryptedData, iv)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if userInfo.Watermark.AppID != appID {
			fmt.Println(ErrAppIDNotMatch)
			return
		}
		fmt.Printf("userData:%+v", userInfo)
	})
}
