package entity

type User struct {
	name string
	id   int64
}

func NewUser(name string, id int64) User {
	return User{
		name: name,
		id:   id,
	}
}
