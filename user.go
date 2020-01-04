package wxwork

/**
 * 成员信息:
 * 	 - 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/90194
 */
type Users struct {
	// 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母。
	UserID string `json:"userid"`

	// 成员名称。长度为1~64个utf8字符
	Name string `json:"name"`

	// 成员别名。长度1~32个utf8字符
	Alias string `json:"alias,omitempty"`

	// 手机号码。企业内必须唯一，mobile/email二者不能同时为空
	Mobile string `json:"mobile,omitempty"`

	// 成员所属部门id列表,不超过20个
	Department []string `json:"department,omitempty"`

	// 部门内的排序值，默认为0，成员次序以创建时间从小到大排列。数量必须和department一致，数值越大排序越前面。有效的值范围是[0, 2^32)
	Order int `json:"order,omitempty"`

	// 职务信息。长度为0~128个字符
	Position string `json:"position,omitempty"`

	// 性别。1表示男性，2表示女性
	Gender uint8 `json:"gender,omitempty"`

	// 邮箱。长度6~64个字节，且为有效的email格式。企业内必须唯一，mobile/email二者不能同时为空
	Email string `json:"email,omitempty"`

	// 座机。32字节以内，由纯数字或’-‘号组成。
	Telephone string `json:"telephone,omitempty"`

	// 个数必须和department一致，表示在所在的部门内是否为上级。1表示为上级，0表示非上级。在审批等应用里可以用来标识上级审批人
	IsLeaderInDept []int `json:"is_leader_in_dept,omitempty"`

	// 成员头像的mediaid，通过素材管理接口上传图片获得的mediaid
	AvatarMediaID int `json:"avatar_mediaid,omitempty"`

	// 启用/禁用成员。1表示启用成员，0表示禁用成员
	Enable bool `json:"enable,omitempty"`

	// 自定义字段。自定义字段需要先在WEB管理端添加，见扩展属性添加方法，否则忽略未知属性的赋值。与对外属性一致，不过只支持type=0的文本和type=1的网页类型，详细描述查看对外属性
	Extattr interface{} `json:"extattr,omitempty"`

	// 是否邀请该成员使用企业微信（将通过微信服务通知或短信或邮件下发邀请，每天自动下发一次，最多持续3个工作日），默认值为true。
	ToInvite bool `json:"to_invite,omitempty"`

	// 成员对外属性
	ExternalProfile UserExternal `json:"external_profile,omitempty"`

	// 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。长度12个汉字内
	ExternalPosition string `json:"external_position,omitempty"`

	Address     string `json:"address,omitempty"`      // 地址。
	Avatar      string `json:"avatar,omitempty"`       // 头像url。 第三方仅通讯录应用可获取
	ThumbAvatar string `json:"thumb_avatar,omitempty"` // 头像缩略图url。第三方仅通讯录应用可获取
	QrCode      string `json:"qr_code,omitempty"`      // 员工二维码
	Status      int8   `json:"status,omitempty"`       // 激活状态: 1=已激活，2=已禁用，4=未激活。
	baseCaller
}
