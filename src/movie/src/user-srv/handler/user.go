package handler

import (
	"context"
	"fmt"
	"movie/src/share/pb"
	"movie/src/user-srv/db"
	"movie/src/user-srv/entity"
)

// 定义绑定方法的结构体
type UserHandler struct {

}

// 创建结构体对象
func NewUserHandler() *UserHandler {
	return &UserHandler{

	}
}

func (c *UserHandler) InsertUser (ctx context.Context, req *pb.InsertUserReq, resp *pb.InsertUserResp) error  {
	fmt.Println("InsertUser ...")

	user := &entity.User{
		Id:      req.Id,
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}

	insertId, err := db.InsertUser(user)
	if err != nil {
		fmt.Println("添加用户错误")
		return err
	}

	resp.Id = int32(insertId)
	return nil
}

func (c *UserHandler) DeleteUser (ctx context.Context, req *pb.DeleteUserReq, resp *pb.DeleteUserResp) error  {
	fmt.Println("DeleteUser ...")
	err := db.DeleteUser(req.GetId())
	if err != nil {
		fmt.Println("删除用户错误")
		return err
	}
	return nil
}

func (c *UserHandler) ModifyUser (ctx context.Context, req *pb.ModifyUserReq, resp *pb.ModifyUserResp) error {
	fmt.Println("ModifyUser ...")
	user := &entity.User{
		Id:      req.Id,
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}

	err := db.ModifyUser(user)
	if err != nil {
		fmt.Println("修改用户错误")
		return err
	}
	return nil
}

func (c *UserHandler) SelectUser (ctx context.Context, req *pb.SelectUserReq, resp *pb.SelectUserResp) error {
	fmt.Println("SelectUser ...")
	user, err := db.SelectUserById(req.GetId())
	if err != nil {
		fmt.Println("查询用户错误")
		return err
	}
	if user != nil {
		resp.Users = user.ToProtoUser()
	}
	return nil
}
