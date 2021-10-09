package wxwork

import (
	"bytes"
	"encoding/json"
)

// CorpTag 企业标签
type CorpTag struct {
	// 标签id
	ID string `json:"id,omitempty"`

	// 标签名称
	Name string `json:"name,omitempty"`

	// 标签创建时间
	CreateTime int64 `json:"create_time,omitempty"`

	// 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order,omitempty"`

	// 标签是否已经被删除，只在指定tag_id/group_id进行查询时返回
	Deleted bool `json:"deleted,omitempty"`
}

// CorpTagGroup 企业标签组
type CorpTagGroup struct {
	// 标签组id
	GroupId string `json:"group_id,omitempty"`

	// 标签组名称
	GroupName string `json:"group_name,omitempty"`

	// 标签组创建时间
	CreateTime int64 `json:"create_time,omitempty"`

	// 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order,omitempty"`

	// 标签组是否已经被删除，只在指定tag_id进行查询时返回
	Deleted bool `json:"deleted,omitempty"`

	// 标签组内的标签列表
	Tag []CorpTag `json:"tag,omitempty"`

	// 规则组id
	StrategyId int `json:"strategy_id,omitempty"`
}

// GetCorpTagList 获取企业标签库
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/92117#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E6%A0%87%E7%AD%BE%E5%BA%93
// 若tag_id和group_id均为空，则返回所有标签。
// 同时传递tag_id和group_id时，忽略tag_id，仅以group_id作为过滤条件。
func (a *Agent) GetCorpTagList(tagId, groupId []string) ([]CorpTagGroup, error) {
	param := map[string][]string{
		"tag_id":   tagId,
		"group_id": groupId,
	}

	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		TagGroup []CorpTagGroup `json:"tag_group"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/get_corp_tag_list", nil, bytes.NewReader(body), &caller)

	return caller.TagGroup, err
}

// AddCorpTag 添加企业客户标签
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/92117#%E6%B7%BB%E5%8A%A0%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (a *Agent) AddCorpTag(group CorpTagGroup, tags []CorpTag) ([]CorpTagGroup, error) {
	param := map[string]interface{}{
		"group_id":   group.GroupId,
		"group_name": group.GroupName,
		"order":      group.Order,
		"tag":        tags,
	}

	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		TagGroup []CorpTagGroup `json:"tag_group"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/add_corp_tag", nil, bytes.NewReader(body), &caller)

	return caller.TagGroup, err
}

// EditCorpTag 编辑企业客户标签
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/92117#%E7%BC%96%E8%BE%91%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (a *Agent) EditCorpTag(id, name string, order int) error {
	param := map[string]interface{}{
		"id":    id,
		"name":  name,
		"order": order,
	}

	body, _ := json.Marshal(param)

	var caller baseCaller

	return a.ExecuteWithToken("POST", "externalcontact/edit_corp_tag", nil, bytes.NewReader(body), &caller)
}

// DelCorpTag 删除企业客户标签
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/92117#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (a *Agent) DelCorpTag(tagId, groupId []string) error {
	param := map[string][]string{
		"tag_id":   tagId,
		"group_id": groupId,
	}

	body, _ := json.Marshal(param)

	var caller baseCaller

	return a.ExecuteWithToken("POST", "externalcontact/del_corp_tag", nil, bytes.NewReader(body), &caller)
}

// GetStrategyTagList 获取指定规则组下的企业客户标签
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/94882#%E8%8E%B7%E5%8F%96%E6%8C%87%E5%AE%9A%E8%A7%84%E5%88%99%E7%BB%84%E4%B8%8B%E7%9A%84%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (a *Agent) GetStrategyTagList(strategy_id int, tagId, groupId []string) ([]CorpTagGroup, error) {
	param := map[string]interface{}{
		"strategy_id": strategy_id,
		"tag_id":      tagId,
		"group_id":    groupId,
	}

	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		TagGroup []CorpTagGroup `json:"tag_group"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/get_strategy_tag_list", nil, bytes.NewReader(body), &caller)

	return caller.TagGroup, err
}

// AddStrategyTag 为指定规则组创建企业客户标签
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/94882#%E4%B8%BA%E6%8C%87%E5%AE%9A%E8%A7%84%E5%88%99%E7%BB%84%E5%88%9B%E5%BB%BA%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (a *Agent) AddStrategyTag(group CorpTagGroup, tags []CorpTag) ([]CorpTagGroup, error) {
	param := map[string]interface{}{
		"group_id":   group.GroupId,
		"group_name": group.GroupName,
		"order":      group.Order,
		"tag":        tags,
	}

	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		TagGroup []CorpTagGroup `json:"tag_group"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/add_corp_tag", nil, bytes.NewReader(body), &caller)

	return caller.TagGroup, err
}
