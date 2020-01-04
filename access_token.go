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
func (a *Agent) RefreshAccessToken() error {
	a.AccessToken.mu.Lock()
	defer a.AccessToken.mu.Unlock()

	var token AccessToken
	path := fmt.Sprintf("%sgettoken?corpid=%s&corpsecret=%s", BaseURL, a.CorpID, a.Secret)

	err := a.Execute("GET", path, nil, &token)
	if err != nil {
		return err
	}

	token.ExpireAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	a.AccessToken = &token

	if a.Cache != nil {
		bt, _ := json.Marshal(token)
		a.Cache.Set("access_token", bt)
	}

	return nil
}

// getAccessTokenFromCache 从缓存中获取 access_token
func (a *Agent) getAccessTokenFromCache() (string, error) {
	if a.Cache == nil {
		return "", fmt.Errorf("client cache processor not found")
	}

	accessToken := a.Cache.Get("access_token")
	err := json.Unmarshal(accessToken, &a.AccessToken)

	if a.AccessToken.IsExpire() || a.AccessToken.AccessToken == "" {
		err = a.RefreshAccessToken()
	}

	return a.AccessToken.AccessToken, err

}

// GetAccessToken 获取access_token
func (a *Agent) GetAccessToken() (string, error) {
	// 如果设置了 缓存器，从缓存器中获取 token，防止频繁刷新
	if a.Cache != nil {
		return a.getAccessTokenFromCache()
	}

	var err error
	if a.AccessToken.IsExpire() {
		err = a.RefreshAccessToken()
	}
	return a.AccessToken.AccessToken, err
}
