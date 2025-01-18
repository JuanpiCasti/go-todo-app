package model

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	RoleID   int    `db:"role_id"`
	Active   bool   `db:"active"`
}

type Role struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type UserWithRole struct {
	User
	Role Role `db:"role"`
}
