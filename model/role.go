package model

type Role int

const (
	Member Role = iota
	ADMIN
)

func (r Role) String() string {
	return [...]string{"Member", "ADMIN"}[r]
}
