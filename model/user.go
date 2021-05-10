package model

type User struct {
	ID int
	Name string
	Status string
}

var Users []User = []User{
	{1, "a", "ADMIN"},
	{2, "b", "B"},
	{3, "c", "C"},
}

//ユーザーが ADMIN かどうかを判断するドメインロジック
func (u *User) IsAdmin() bool {
	return u.Status == "ADMIN"
}

func UpdateUser(u *User) error {
	index := u.ID - 1
	updatedUser := User{
		u.ID, u.Name, u.Status,
	}
	Users[index] = updatedUser
	return nil
}