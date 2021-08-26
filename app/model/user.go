package model

// User is the golang structure for table user.
type Users struct {
	Id       uint   `orm:"id,primary" json:"id"`       // 用户ID
	Password string `orm:"password"   json:"password"` // 用户密码
	Nickname string `orm:"nickname"   json:"nickname"` // 用户昵称
}
