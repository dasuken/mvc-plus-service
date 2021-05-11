package service

import "github.com/noguchidaisuke/mvc-plus-service/model"

// modelに対する処理を切り出す。
func UpdateUserStatus(id int, name string) error {
	// NewDbConnでconnection作成
	// userオブジェクトは本来GetByIdから取得
	user := model.User{
		ID:   id,
		Name: name,
	}

	// 第一引数にconnectonを渡し、第２引数に取得したuserを渡す。
	err := model.UpdateUser(&user)

	return err
}
