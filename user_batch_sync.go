package wxwork

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// BatchSync 异步批量处理请求结构体
type BatchSync struct {
	// 上传的csv文件的media_id
	MediaId string `json:"media_id" xml:"media_id"`
	// 是否邀请新建的成员使用企业微信（将通过微信服务通知或短信或邮件下发邀请，每天自动下发一次，最多持续3个工作日），默认值为true。
	ToInvite bool `json:"to_invite,omitempty" xml:"to_invite,omitempty"`
	// 回调信息。如填写该项则任务完成后，通过callback推送事件给企业。具体请参考应用回调模式中的相应选项
	Callback Callback `json:"callback" xml:"callback"`
}

// SyncUserBatch 增量更新成员
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90980
func (a *Agent) SyncUserBatch(bs *BatchSync) (string, error) {
	body, _ := json.Marshal(bs)

	var caller struct {
		baseCaller
		JobId string `json:"jobid" xml:"jobid"`
	}

	err := a.ExecuteWithToken("POST", "batch/syncuser", nil, bytes.NewReader(body), &caller)
	if err != nil {
		return "", err
	}

	return caller.JobId, nil
}

// ReplaceUserBatch 全量覆盖成员
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90981
func (a *Agent) ReplaceUserBatch(bs *BatchSync) (string, error) {
	body, _ := json.Marshal(bs)

	var caller struct {
		baseCaller
		JobId string `json:"jobid" xml:"jobid"`
	}

	err := a.ExecuteWithToken("POST", "batch/replaceuser", nil, bytes.NewReader(body), &caller)
	if err != nil {
		return "", err
	}

	return caller.JobId, nil
}

// ReplacePartyBatch 全量覆盖部门
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90982
func (a *Agent) ReplacePartyBatch(bs *BatchSync) (string, error) {
	body, _ := json.Marshal(bs)

	var caller struct {
		baseCaller
		JobId string `json:"jobid" xml:"jobid"`
	}

	err := a.ExecuteWithToken("POST", "batch/replaceparty", nil, bytes.NewReader(body), &caller)
	if err != nil {
		return "", err
	}

	return caller.JobId, nil
}

// BatchResult 异步任务结果
type BatchResult struct {
	baseCaller
	// 任务状态，整型，1表示任务开始，2表示任务进行中，3表示任务已完成
	Status int `json:"status" xml:"status"`
	// 操作类型，字节串，目前分别有：1. sync_user(增量更新成员) 2. replace_user(全量覆盖成员)3. replace_party(全量覆盖部门)
	Type string `json:"type" xml:"type"`
	// 任务运行总条数
	Total int `json:"total" xml:"total"`
	// 目前运行百分比，当任务完成时为100
	Percentage int `json:"percentage" xml:"percentage"`
	// 详细的处理结果
	Result BatchResultDetail `json:"result" xml:"result"`
}

// BatchResultDetail 异步任务结果内容
type BatchResultDetail struct {
	// 成员UserID。对应管理端的帐号
	UserID string `json:"userid,omitempty" xml:"userid,omitempty"`
	// 操作类型（按位或）：1 新建部门 ，2 更改部门名称， 4 移动部门， 8 修改部门排序
	Action int `json:"action,omitempty" xml:"action,omitempty"`
	// 部门ID
	PartyID int `json:"partyid,omitempty" xml:"partyid,omitempty"`
	baseCaller
}

// GetResultBatch 获取异步任务结果
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90983
func (a *Agent) GetResultBatch(jobid string) (BatchResult, error) {
	var caller BatchResult

	query := url.Values{}
	query.Set("jobid", jobid)
	err := a.ExecuteWithToken("GET", "batch/getresult", query, nil, &caller)

	return caller, err
}
