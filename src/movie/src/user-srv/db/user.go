package db

import (
	"database/sql"
	"fmt"
	"movie/src/user-srv/entity"
)

// SelectUserById  根据 id 查询用户
func SelectUserById(id int32) (*entity.User, error){
	fmt.Println("-------SelectUserById , id : ", id)
	// 创建 User 对象，用于返回
	user := new(entity.User)
	// 执行查询
	err := db.Get(user, "SELECT name, address, phone FROM user WHERE id = ?", id)
	if err != nil {
		if err != sql.ErrNoRows {   // 表示其它数据库查询错误
			return nil, err
		}
		return nil, nil  // 表示没有查到该用户
	}

	return user, nil  // 表示查到了该用户
}

// 添加用户
func InsertUser(user *entity.User) (int64, error) {

	rep, err := db.Exec("INSERT INTO `user` (`name`, `address`, `phone`)VALUE(?,?,?)",
						user.Name, user.Address, user.Phone)

	if err != nil {
		return 0, err
	}
	return rep.LastInsertId()
}

// 修改用户
func ModifyUser(user *entity.User) error {
	_, err := db.Exec("UPDATE `user` set `name`=?, `address`=?, `phone`=? WHERE `id`=?",
		user.Name,user.Address,user.Phone, user.Id)
	if err != nil {
		return err
	}
	return nil
}

// 删除用户
func DeleteUser(id int32) error {
	_, err := db.Exec("DELETE FROM `user` WHERE `id`=?",id )
	if err != nil {
		return err
	}
	return nil
}

