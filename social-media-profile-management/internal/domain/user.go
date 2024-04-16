package domain

type User struct {
	UserID   int
	UserName string
}

func (u User) ID() int {
	return u.UserID
}
