package query

type LogoutQuery struct {
	Auth string `db:"auth"`
}

type SignInQuery struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Out      struct {
		Account     Account `db:"account" json:"account"`
		AccessToken string  `json:"access_token"`
	}
}

type GetAccountQuery struct {
	Email string `json:"email" db:"email"`
	Out   struct {
		Account *Account `json:"account" db:"account"`
	}
}

type Account struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}

type SignUpQuery struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Username string `db:"username" json:"username"`
}
