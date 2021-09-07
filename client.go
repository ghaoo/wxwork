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
	corpID string
	// agentID 应用ID
	agentID int
	// secret 应用秘钥
	secret string
	// accessToken 应用登录凭证
	accessToken *AccessToken
	// 是否开启Debug
	debug bool

	cache Cache
	crypt *wxbizmsgcrypt.WXBizMsgCrypt

	client *http.Client
}

// SetMsgCrypt 设置消息加密认证
func (a *Agent) SetMsgCrypt(token, encodingAESKey string) *Agent {
	a.crypt = wxbizmsgcrypt.NewWXBizMsgCrypt(token, encodingAESKey, a.corpID, wxbizmsgcrypt.XmlType)

	return a
}

// NewAgent 新建一个应用
func NewAgent(corpid string, agentid int) *Agent {

	return &Agent{
		corpID:      corpid,
		agentID:     agentid,
		accessToken: new(AccessToken),
		debug:       false,
		client:      &http.Client{},
	}
}

// SetDebug 开启debug模式调用接口
// 注意: debug模式有使用频率限制，同一个api每分钟不能超过5次，所以在完成调试之后，请记得关掉debug。
func (a *Agent) SetDebug(debug bool) *Agent {
	a.debug = debug
	return a
}

// WithSecret 返回添加了secret的应用
func (a *Agent) WithSecret(secret string) *Agent {
	agent := NewAgent(a.corpID, a.agentID).SetDebug(a.debug)
	agent.secret = secret
	return agent
}

// SetCache 设置缓存处理器
func (a *Agent) SetCache(cache Cache) *Agent {
	a.cache = cache

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

	if a.debug {
		query.Set("debug", "1")
	}

	u, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}

	u.Path = path.Join(u.Path, uri)

	u.RawQuery = query.Encode()

	return a.Execute(method, u.String(), body, caller)
}
