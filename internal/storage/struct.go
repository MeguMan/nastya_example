package storage

type Test struct {
	ID   int    `db:"id"`
	Text string `db:"text"`
}

type Tokens struct {
	ID    int
	Token string
	Role  role
}

type role string

const (
	Admin   role = "admin"
	Default role = "default"
)
