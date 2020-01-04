package workwx

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// 企业微信API接口基础网址
const BaseURL = "https://qyapi.weixin.qq.com/cgi-bin/"

type Client struct {
	// 企业ID
	CorpID string
	// AgentID 应用ID
	AgentID int
	// Secret 应用秘钥
	Secret string
	// AccessToken 应用登录凭证
	AccessToken *AccessToken

	Cache  Cache
	client *http.Client
}

func NewClientFromEnv() (*Client, error) {
	corpid := os.Getenv("WEWORK_CORP_ID")
	secret := os.Getenv("WEWORK_SECRET")
	agentid, _ := strconv.Atoi(os.Getenv("WEWORK_AGENT_ID"))
	if corpid == "" || secret == "" {
		return nil, errors.New("请检查 WEWORK_CORP_ID、WEWORK_SECRET、WEWORK_AGENT_ID 是否已全部设置成功")
	}

	return &Client{
		CorpID:      corpid,
		AgentID:     agentid,
		Secret:      secret,
		AccessToken: new(AccessToken),
		Cache:       Bolt(),
		client:      &http.Client{},
	}, nil
}

func NewClient(corpid, secret string, agentid int) *Client {

	return &Client{
		CorpID:      corpid,
		AgentID:     agentid,
		Secret:      secret,
		AccessToken: new(AccessToken),
		Cache:       Bolt(),
		client:      &http.Client{},
	}
}

// SetCache 设置缓存处理器
func (c *Client) SetCache(cache Cache) {
	c.Cache = cache
}

// SetHttpClient 设置一个可用的 http client
func (c *Client) SetHttpClient(client *http.Client) {
	c.client = client
}

type Caller interface {
	Success() bool
	Error() error
}

type baseCaller struct {
	ErrCode int    `json:"errcode,omitempty"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg,omitempty"`  // 返回码提示语
}

func (b baseCaller) Success() bool {
	return b.ErrCode == 0
}

func (b baseCaller) Error() error {
	return errors.New(b.ErrMsg)
}

// Execute 在默认的http客户端执行一个http请求
func (c *Client) Execute(method string, url string, body io.Reader, caller Caller) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
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
func (c *Client) ExecuteWithToken(method string, path string, body io.Reader, caller Caller) error {

	accessToken, err := c.GetAccessToken()
	if err != nil {
		return err
	}

	query := url.Values{}
	query.Set("access_token", accessToken)

	u, err := url.Parse(BaseURL + path)
	if err != nil {
		panic(err)
	}

	u.RawQuery = query.Encode()

	return c.Execute(method, u.String(), body, caller)
}
