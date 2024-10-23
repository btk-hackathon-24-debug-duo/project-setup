package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type GetUserModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
