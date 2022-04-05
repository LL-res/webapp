package models

type SignUpUser struct {
	Name      string `json:"name" binding:"required" form:"name"`
	Password  string `json:"password" binding:"required" form:"password"`
	RPassword string `json:"r_password" binding:"required,eqfield=Password" form:"r_password"`
}
type User struct {
	Name     string
	PassWord string
	ID       string
}
type LogInUser struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"name" binding:"required"`
}
