// Package wechat 微信返回的数据进行解密
package wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

var (
	ErrAppIDNotMatch       = errors.New("app id not match")
	ErrInvalidBlockSize    = errors.New("invalid block size")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

type WXDataCrypt struct {
	appId      string
	sessionKey string
}

func NewWXDataCrypt(appid, sessionKey string) *WXDataCrypt {
	return &WXDataCrypt{
		appId:      appid,
		sessionKey: sessionKey,
	}
}

// Decrypt WeChat data decrypt
// The appid is judged by the external
// business layer, and the decryption layer
// is only responsible for decoding,
// and does not judge the accuracy of the data.
func (w *WXDataCrypt) Decrypt(encryptedData, iv string, res interface{}) error {
	aesKey, err := base64.StdEncoding.DecodeString(w.sessionKey)
	if err != nil {
		return err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return err
	}
	err = json.Unmarshal(cipherText, res)
	if err != nil {
		return err
	}
	return nil
}

// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}
