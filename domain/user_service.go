package domain

type UserService interface {
	Create(user *User) error
}
