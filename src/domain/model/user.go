package model

type User struct {
	Id int64
	Name  string
	Email string
	Password string
}

func NewUser(id int64, name string, email string, password string) (*User, error) {

	user := &User{
		Id:   id,
		Name: name,
		Email: email,
		Password: password,
	}

	return user, nil
}