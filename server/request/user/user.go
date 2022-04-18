package user

type UserReq struct {
	Login
}

// Register 用户注册
type Register struct {
	Username    string `json:"userName"`                                                               // 用户名
	Password    string `json:"passWord"`                                                               // 密码
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`                                   // 昵称
	Avatar      string `json:"avatar" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"` // 头像
	AuthorityId string `json:"authorityId" gorm:"default:888"`                                         // 用户权限id
	//AuthorityIds []string `json:"authorityIds"`                                                           // 用户权限组
}

// Login 用户登录
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
