package models

type UserSignIn struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserSignUp struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"firstName" json:"firstName"`
	LastName  string `form:"lastName" json:"lastName"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}
