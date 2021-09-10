package wxwork

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

// Department 成员部门信息
// 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/90204
type Department struct {
	// 部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id
	ID int `json:"id,omitempty" xml:"Id"`
	// 部门名称。长度限制为1~32个字符，字符不能包括\:?”<>｜
	Name string `json:"name" xml:"Name"`
	// 英文名称，需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:?”<>｜
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 父部门id，32位整型
	ParentID int `json:"parentid" xml:"ParentID"`
	// 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order,omitempty" xml:"order,omitempty"`
}

// CreateDepartment 创建部门
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90205
func (a *Agent) CreateDepartment(dept Department) (int, error) {
	body, _ := json.Marshal(dept)

	var caller struct {
		baseCaller
		ID int `json:"id"`
	}
	err := a.ExecuteWithToken("POST", "department/create", nil, bytes.NewReader(body), &caller)

	return caller.ID, err
}

// UpdateDepartment 更新部门
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90206
func (a *Agent) UpdateDepartment(dept Department) error {
	body, _ := json.Marshal(dept)

	var caller baseCaller
	return a.ExecuteWithToken("POST", "department/update", nil, bytes.NewReader(body), &caller)
}

// DeleteDepartment 删除部门
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90207
func (a *Agent) DeleteDepartment(id int) error {
	var caller baseCaller
	query := url.Values{}
	query.Set("id", strconv.Itoa(id))
	return a.ExecuteWithToken("GET", "department/delete", query, nil, &caller)
}

// ListDepartment 获取部门列表
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90208
func (a *Agent) ListDepartment(id ...int) ([]Department, error) {
	query := url.Values{}
	if len(id) > 0 {
		query.Set("id", strconv.Itoa(id[0]))
	}

	var caller struct {
		baseCaller
		Department []Department `json:"department"`
	}

	var err error
	err = a.ExecuteWithToken("GET", "department/list", query, nil, &caller)

	return caller.Department, err
}
