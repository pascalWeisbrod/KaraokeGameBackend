package model


type User struct {
	Name   string
	Points int
}

type UserList struct {
	Users []User
}

func NewUserList() UserList {
    users := make([]User, 0)
    return UserList{ Users: users }
}

func (list *UserList) Add(user User) {
	list.Users = append(list.Users, user)
}
