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

const BaseURL = "https://qyapi.weixin.qq.com/cgi-bin/"

type Agent struct {
	// 企业ID
	CorpID string
	// AgentID 应用ID
	AgentID int
	// Secret 应用秘钥
	Secret string
	// AccessToken 应用登录凭证
	AccessToken *AccountToken

	cachePath string

	client *http.Client
}

func NewAgentFromEnv() (*Agent, error) {
	cachePath := ".data/wework"
	if os.Getenv("CACHE_PATH") == "" {
		cachePath = os.Getenv("CACHE_PATH")
	}

	corpid := os.Getenv("WEWORK_CORP_ID")
	secret := os.Getenv("WEWORK_SECRET")
	agentid, _ := strconv.Atoi(os.Getenv("WEWORK_AGENT_ID"))
	if corpid == "" || secret == "" {
		return nil, errors.New("请检查 WEWORK_CORP_ID、WEWORK_SECRET、WEWORK_AGENT_ID 是否已全部设置成功")
	}

	return &Agent{
		CorpID:    corpid,
		AgentID:   agentid,
		Secret:    secret,
		cachePath: cachePath,
		client:    http.DefaultClient,
	}, nil
}

func NewAgent(corpid, secret string, agentid int, cachePath ...string) *Agent {
	cache := ".data/wework"
	if len(cachePath) > 0 {
		cache = cachePath[0]
	}

	return &Agent{
		CorpID:    corpid,
		AgentID:   agentid,
		Secret:    secret,
		cachePath: cache,
		client:    http.DefaultClient,
	}
}

type Caller interface {
	Success() bool
	Error() error
}

type baseCaller struct {
	ErrCode int    `json:"errcode,omitempty"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg,omitempty"`  // 返回码提示语
}

func (b *baseCaller) Success() bool {
	return b.ErrCode == 0
}

func (b *baseCaller) Error() error {
	return errors.New(b.ErrMsg)
}

func (a *Agent) Execute(method string, path string, body io.Reader, caller Caller) error {
	req, err := http.NewRequest(method, BaseURL+path, body)
	if err != nil {
		return err
	}

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

func (a *Agent) ExecuteWithToken(method string, path string, body io.Reader, caller Caller) error {

	accountToken, err := a.getAccountToken()
	if err != nil {
		return err
	}

	query := url.Values{}
	query.Set("account_token", accountToken)

	base, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}

	base.Path = path
	base.RawQuery = query.Encode()

	return a.Execute(method, base.String(), body, caller)
}
