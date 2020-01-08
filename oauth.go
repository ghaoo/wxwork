package wxwork

import "net/url"

// 获取访问用户身份
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/91023
// isInterior 是否为内部人员，true是 false不是
func (a *Agent) GetUserInfo(code string, isInterior ...bool) (id, deviceId string, err error) {
	var caller struct {
		baseCaller
		UserId   string `json:"UserId"`
		OpenId   string `json:"OpenId"`
		DeviceId string `json:"DeviceId"`
	}

	query := url.Values{}
	query.Set("code", code)

	err = a.ExecuteWithToken("GET", "user/getuserinfo", query, nil, &caller)

	if len(isInterior) > 0 && !isInterior[0] {
		return caller.OpenId, caller.DeviceId, err
	} else {
		return caller.UserId, caller.DeviceId, err
	}
}
