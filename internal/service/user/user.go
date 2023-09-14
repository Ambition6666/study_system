package user

import (
	"studysystem/internal/repository"
	"studysystem/models"
)

// 获取用户信息
func GetUserInfo(id int64) *models.User {
	u := repository.Search_user_by_id(id)
	return u
}

// 更新用户信息
func UpdateUserInfo(id int64, action int, data string) (int, string) {
	switch action {
	case 1:
		repository.Update_user(id, action, data)
	case 2:
		repository.Update_user(id, action, data)
	case 3:
		repository.Update_user(id, action, data)
	}
	return 200, "修改成功"
}

// 删除用户信息
func DeleteUser(id int64) (int, string) {
	repository.DeleteUser(id)
	return 200, "删除成功"
}
