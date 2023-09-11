package repository

import (
	"context"
	"strconv"
	"studysystem/models"
	"studysystem/sql"
)

// 创建用户
func Create_user(a *models.User) {
	db := sql.GetMysqlDB()
	db.Create(a)
}

// 查询用户

// 根据email查询用户,通常在登录时用
func Search_user(email string) *models.User {
	db := sql.GetMysqlDB()
	user := new(models.User)
	db.Where("email = ?", email).Find(user)
	return user
}

// 根据id查询用户
func Search_user_by_id(id int64) *models.User {
	db := sql.GetMysqlDB()
	user := new(models.User)
	db.Where("id = ?", id).Find(user)
	return user
}

// 修改用户
// action:
// 1-->nickname,2-->avatar,3-->individualresume
func Update_user(id int64, action int, data string) {
	db := sql.GetMysqlDB()
	u := Search_user_by_id(id)
	switch action {
	case 1:
		u.NickName = data
	case 2:
		u.Avatar = data
	case 3:
		u.IndividualResume = data
	}
	db.Save(u)
}

// 在redis保存头像本地地址
func Save_local_avatar_path(id int64, data string) error {
	rdb := sql.GetRedis()
	return rdb.Set(context.Background(), "avatar"+strconv.FormatInt(id, 10), data, -1).Err()
}

// 获取redis保存的头像地址
func Get_local_avatar_path(id int64) (string, error) {
	rdb := sql.GetRedis()
	return rdb.Get(context.Background(), "avatar"+strconv.FormatInt(id, 10)).Result()
}

// 删除用户
func DeleteUser(id int64) {
	db := sql.GetMysqlDB()
	db.Where("id = ?", id).Delete(&models.User{})
}

// 创建答题记录
func CreateCommitRecord(v *models.CommitRecord) {
	pdb := sql.GetPgsql()
	pdb.Create(v)
}

// 获取答题记录
func GetRecord(q string, uid int64) *models.CommitRecord {
	pdb := sql.GetPgsql()
	v := new(models.CommitRecord)
	pdb.Where("qid=? and uid =?", q, uid).Find(v)
	return v
}
