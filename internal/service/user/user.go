package user

import (
	"fmt"
	"studysystem/internal/repository"
	"studysystem/vo"
)

// 获取用户信息
func GetUserInfo(id int64) *vo.UserInfo {
	u := repository.Search_user_by_id(id)
	uinfo := &vo.UserInfo{
		ID:               u.ID,
		Email:            u.Email,
		Avatar:           u.Avatar,
		NickName:         u.NickName,
		IndividualResume: u.IndividualResume,
	}
	return uinfo
}

// 更新用户信息
func UpdateUserInfo(id int64, action int, data string) (int, string) {
	switch action {
	case 1:
		repository.Update_user(id, action, data)
	case 2:
		err := repository.Save_local_avatar_path(id, data)
		if err != nil {
			fmt.Println(err)
			return 500, "保存头像失败"
		}
		repository.Update_user(id, action, "/identified/avatar/")
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
