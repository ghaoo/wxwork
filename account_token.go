package workwx

import (
	"fmt"
	"sync"
	"time"
)

type AccountToken struct {
	AccountToken string    `json:"account_token"`        // 获取到的凭证，最长为512字节
	ExpiresIn    int64     `json:"expires_in,omitempty"` // 凭证的有效时间（秒），通常为2小时（7200秒）
	ExpireAt     time.Time `json:"expire_at,omitempty"`  // 过期时间，超过时重新获取
	baseCaller

	mu sync.Mutex
}

// IsExpire 验证 account_token 是否过期
func (token *AccountToken) IsExpire() bool {
	return token.ExpireAt.Before(time.Now())
}

func (a *Agent) getAccountToken() (string, error) {
	// account_token 为空或者已经过期，重新获取
	if a.AccessToken == nil || a.AccessToken.IsExpire() {
		var token AccountToken
		path := fmt.Sprintf("/gettoken?corpid=%s&corpsecret=%s", a.CorpID, a.Secret)
		err := a.Execute("GET", path, nil, &token)
		if err != nil {
			return "", err
		}
		a.AccessToken.mu.Lock()
		a.AccessToken.ExpireAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
		a.AccessToken.mu.Unlock()
	}
	return a.AccessToken.AccountToken, nil
}
