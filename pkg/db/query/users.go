package query

type GetUserQuery struct {
	UserId string `json:"user_id"`
	Out    struct {
		Account Account `db:"account" json:"account"`
	}
}

type UpdateUserQuery struct {
	UserId   string `db:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type DeleteUserQuery struct {
	UserId string `db:"id"`
}
