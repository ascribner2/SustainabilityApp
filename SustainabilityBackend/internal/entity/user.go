package entity

type UserEntity interface {
	GetEmail() string
	GetPass() string
}

// Model for new users
type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (nu *NewUser) GetEmail() string {
	return nu.Email
}

func (nu *NewUser) GetPass() string {
	return nu.Password
}

// Model for user
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPass() string {
	return u.Password
}
