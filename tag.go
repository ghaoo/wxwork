package wxwork

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

// Tag 标签结构
type Tag struct {
	TagID     int    `json:"tagid,omitempty"`     // 标签id
	TagName   string `json:"tagname,omitempty"`   // 标签名称
	UserList  []User `json:"userlist,omitempty"`  // 标签成员ID列表
	PartyList []int  `json:"partylist,omitempty"` // 标签部门ID列表
}

// CreateTag 创建标签
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90210
func (a *Agent) CreateTag(tagName string, tagId ...int) (int, error) {
	tag := &Tag{
		TagName: tagName,
	}
	if len(tagId) > 0 {
		tag.TagID = tagId[0]
	}
	body, _ := json.Marshal(tag)

	var caller struct {
		baseCaller
		TagId int `json:"tagid"`
	}
	err := a.ExecuteWithToken("POST", "tag/create", nil, bytes.NewReader(body), &caller)
	if err != nil {
		return 0, err
	}

	return caller.TagId, nil
}

// UpdateTag 更新标签名称
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90211
func (a *Agent) UpdateTag(tagName string, tagId int) error {
	tag := &Tag{
		TagID:   tagId,
		TagName: tagName,
	}

	body, _ := json.Marshal(tag)

	var caller baseCaller
	err := a.ExecuteWithToken("POST", "tag/update", nil, bytes.NewReader(body), &caller)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTag 删除标签，必须为标签创建者才可删除
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90212
func (a *Agent) DeleteTag(id int) error {
	var caller baseCaller
	query := url.Values{}
	query.Set("tagid", strconv.Itoa(id))
	err := a.ExecuteWithToken("GET", "tag/delete", query, nil, &caller)
	if err != nil {
		return err
	}

	return nil
}

// GetTag 获取标签成员
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90213
func (a *Agent) GetTag(id int) (tag Tag, err error) {
	var caller struct {
		baseCaller
		TagName   string `json:"tagname"`
		UserList  []User `json:"userlist"`
		PartyList []int
	}

	query := url.Values{}
	query.Set("tagid", strconv.Itoa(id))
	err = a.ExecuteWithToken("GET", "tag/get", query, nil, &caller)
	if err != nil {
		return tag, err
	}

	tag = Tag{
		TagID:     id,
		TagName:   caller.TagName,
		PartyList: caller.PartyList,
	}

	for _, users := range caller.UserList {
		tag.UserList = append(tag.UserList, users)
	}

	return tag, nil
}

func (a *Agent) addOrDelTagUsers(path string, tag Tag) (invalidlist string, invalidparty []int, err error) {
	userList := make([]string, 0)
	for _, user := range tag.UserList {
		userList = append(userList, user.UserID)
	}

	tags := map[string]interface{}{
		"tagid":     tag.TagID,
		"userlist":  userList,
		"partylist": tag.PartyList,
	}

	body, _ := json.Marshal(tags)

	var caller struct {
		baseCaller
		InvalidList  string `json:"invalidlist,omitempty"`
		InvalidParty []int  `json:"invalidparty,omitempty"`
	}

	err = a.ExecuteWithToken("POST", path, nil, bytes.NewReader(body), &caller)

	return caller.InvalidList, caller.InvalidParty, err
}

// AddTagUsers 增加标签成员
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90214
func (a *Agent) AddTagUsers(id int, users []User, parties []int) (invalidlist string, invalidparty []int, err error) {
	tag := Tag{
		TagID:     id,
		UserList:  users,
		PartyList: parties,
	}

	return a.addOrDelTagUsers("tag/addtagusers", tag)
}

// DelTagUsers 删除标签成员
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90215
func (a *Agent) DelTagUsers(id int, users []User, parties []int) (invalidlist string, invalidparty []int, err error) {
	tag := Tag{
		TagID:     id,
		UserList:  users,
		PartyList: parties,
	}
	return a.addOrDelTagUsers("tag/deltagusers", tag)
}

// ListTags 获取标签列表
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90216
func (a *Agent) ListTags() ([]Tag, error) {

	var caller struct {
		baseCaller
		Taglist []Tag
	}

	err := a.ExecuteWithToken("GET", "tag/list", nil, nil, &caller)
	return caller.Taglist, err
}
