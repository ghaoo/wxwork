package wxwork

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// AccessToken 定义了获取 access_token 时的响应
type AccessToken struct {
	AccessToken string    `json:"access_token"`         // 获取到的凭证，最长为512字节
	ExpiresIn   int64     `json:"expires_in,omitempty"` // 凭证的有效时间（秒），通常为2小时（7200秒）
	ExpireAt    time.Time `json:"expire_at,omitempty"`  // 过期时间，超过时重新获取
	baseCaller

	mu sync.Mutex
}

// IsExpire 验证 access_token 是否过期
func (token *AccessToken) IsExpire() bool {
	return token.ExpireAt.Before(time.Now())
}

// RefreshAccessToken 用于刷新 access_token
func (c *Client) RefreshAccessToken() error {
	c.AccessToken.mu.Lock()
	defer c.AccessToken.mu.Unlock()

	var token AccessToken
	path := fmt.Sprintf("%sgettoken?corpid=%s&corpsecret=%s", BaseURL, c.CorpID, c.Secret)

	err := c.Execute("GET", path, nil, &token)
	if err != nil {
		return err
	}

	token.ExpireAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	c.AccessToken = &token

	if c.Cache != nil {
		bt, _ := json.Marshal(token)
		c.Cache.Set("access_token", bt)
	}

	return nil
}

// getAccessTokenFromCache 从缓存中获取 access_token
func (c *Client) getAccessTokenFromCache() (string, error) {
	if c.Cache == nil {
		return "", fmt.Errorf("client cache processor not found")
	}

	accessToken := c.Cache.Get("access_token")
	err := json.Unmarshal(accessToken, &c.AccessToken)

	if c.AccessToken.IsExpire() || c.AccessToken.AccessToken == "" {
		err = c.RefreshAccessToken()
	}

	return c.AccessToken.AccessToken, err

}

// GetAccessToken 获取access_token
func (c *Client) GetAccessToken() (string, error) {
	// 如果设置了 缓存器，从缓存器中获取 token，防止频繁刷新
	if c.Cache != nil {
		return c.getAccessTokenFromCache()
	}

	var err error
	if c.AccessToken.IsExpire() {
		err = c.RefreshAccessToken()
	}
	return c.AccessToken.AccessToken, err
}
