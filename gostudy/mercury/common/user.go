package common

const (
	UserSexMan = 1
	UserSexWomen = 2
)
type UserInfo struct {
	UserId int64 `json:"user_id" db:"user_id"`
	Username string	`json:"user" db:"username"`
	Nickname	string	`json:"nickname" db:"nickname"`
	Password string	`json:"password" "password"`
	Email string	`json:"email" db:"email"`
	Sex int	`json:"sex" db:"sex"`

}