package domain

// User struct defines a user
type User struct {
	UserID   int    `json:"userID"`
	UserName string `json:"userName"`
}

func (u User) ID() int {
	return u.UserID
}
