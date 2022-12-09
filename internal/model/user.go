package model

type User struct {
	baseModel
	Username      string `json:"username" gorm:"username"`               // 昵称
	Mobile        string `json:"mobile" gorm:"mobile"`                   // 手机号
	Email         string `json:"email" gorm:"email"`                     // 邮箱
	Password      string `json:"password" gorm:"password"`               // 密码
	PublicKey     string `json:"public_key" gorm:"public_key"`           // 公钥
	PersonalSign  string `json:"personal_sign" gorm:"personal_sign"`     // 个性签名
	TotalUsedTime int64  `json:"total_used_time" gorm:"total_used_time"` // 在线时长
	LastIp        string `json:"last_ip" gorm:"last_ip"`                 // 最后登录ip
	LastLoginTime int64  `json:"last_login_time" gorm:"last_login_time"` // 最后登录时间
	Status        int8   `json:"status" gorm:"status"`                   // 账号状态: 0:初始化, 1:待激活, 2:已激活, 3:禁用, 4:已注销
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}
