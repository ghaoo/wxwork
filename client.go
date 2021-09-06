package wxwork

import (
	"encoding/json"
	"errors"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"io"
	"net/http"
	"net/url"
	"path"
)

// BaseURL 企业微信API接口基础网址
const BaseURL = "https://qyapi.weixin.qq.com/cgi-bin/"

// Agent 应用结构
type Agent struct {
	// 企业ID
	CorpID string
	// AgentID 应用ID
	AgentID int
	// Secret 应用秘钥
	Secret string
	// AccessToken 应用登录凭证
	AccessToken *AccessToken

	Cache    Cache
	callback *Callback

	client *http.Client
}

// Callback 应用回调配置，需加密
type Callback struct {
	// 企业应用接收企业微信推送请求的访问协议和地址，支持http或https协议
	URL string `json:"url" xml:"url"`
	// 用于生成签名
	Token string `json:"token" xml:"token"`
	// 用于消息体的加密，是AES密钥的Base64编码
	EncodingAESKey string `json:"encodingaeskey" xml:"encodingaeskey"`

	crypt *wxbizmsgcrypt.WXBizMsgCrypt
}

// SetCallback 设置回调参数
func (a *Agent) SetCallback(token, encodingAESKey string) *Agent {
	callback := &Callback{
		Token:          token,
		EncodingAESKey: encodingAESKey,
	}

	callback.crypt = wxbizmsgcrypt.NewWXBizMsgCrypt(token, encodingAESKey, a.CorpID, wxbizmsgcrypt.XmlType)

	a.callback = callback

	return a
}

// NewAgent 新建一个应用
func NewAgent(corpid string, agentid int) *Agent {

	return &Agent{
		CorpID:      corpid,
		AgentID:     agentid,
		AccessToken: new(AccessToken),
		client:      &http.Client{},
	}
}

// WithSecret 返回添加了secret的应用
func (a *Agent) WithSecret(secret string) *Agent {
	agent := NewAgent(a.CorpID, a.AgentID)
	agent.Secret = secret
	return agent
}

// SetCache 设置缓存处理器
func (a *Agent) SetCache(cache Cache) *Agent {
	a.Cache = cache

	return a
}

// SetHttpClient 设置一个可用的 http client
func (a *Agent) SetHttpClient(client *http.Client) *Agent {
	a.client = client

	return a
}

// Caller 执行 http 访问时响应成功接口
type Caller interface {
	Success() bool
	Error() error
}

// baseCaller 基础响应
type baseCaller struct {
	ErrCode int    `json:"errcode,omitempty" xml:"ErrCode"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg,omitempty" xml:"ErrMsg"`   // 返回码提示语
}

// Success 返回是否调用成功
func (b baseCaller) Success() bool {
	return b.ErrCode == 0
}

// Error 返回失败信息
func (b baseCaller) Error() error {
	return errors.New(b.ErrMsg)
}

// Execute 在默认的http客户端执行一个http请求
func (a *Agent) Execute(method string, url string, body io.Reader, caller Caller) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(caller); err != nil {
		return err
	}

	if !caller.Success() {
		return caller.Error()
	}

	return nil
}

// ExecuteWithToken 在默认的http客户端执行一个http请求，并在请求中附带 AccessToken
func (a *Agent) ExecuteWithToken(method string, uri string, query url.Values, body io.Reader, caller Caller) error {

	accessToken, err := a.GetAccessToken()
	if err != nil {
		return err
	}

	if query == nil {
		query = url.Values{}
	}

	query.Set("access_token", accessToken)

	u, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}

	u.Path = path.Join(u.Path, uri)

	u.RawQuery = query.Encode()

	return a.Execute(method, u.String(), body, caller)
}
