package db

type LogoutQuery struct {
	Auth string `db:"auth"`
}

type SignInQuery struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type SignUpQuery struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Username string `db:"username" json:"username"`
}
