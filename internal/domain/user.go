package domain

var USER = make(map[int]*User)

type User struct {
	Id       int
	Username string
	Password string
	Role     string
}

type UserRepository interface {
	Create(user *User) error
}
