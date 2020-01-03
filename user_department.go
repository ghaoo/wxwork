package workwx

/**
 * 成员部门信息
 * - 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/90204
 */
type Department struct {
	// 部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id
	ID int
	// 部门名称。长度限制为1~32个字符，字符不能包括\:?”<>｜
	Name string
	// 英文名称，需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:?”<>｜
	NameEn string
	// 父部门id，32位整型
	ParentID int
	// 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int
}
