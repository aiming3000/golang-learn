package handler

import (
	"context"
	"fmt"
	user_serve "mall/user_serve/proto/admin_user"
)

type Adminuser struct{}

func (a *Adminuser) AdminUserlogin(ctx context.Context, in *user_serve.AdminUserRequest, out *user_serve.AdminUserResponse) error {
	user_name := in.Username
	password := in.Password
	fmt.Println(user_name, password)
	out.Code = 200
	out.Msg = "登录成功"
	return nil
}

func (a *Adminuser) FrontUserList(ctx context.Context, in *user_serve.FrontUsersRequest, out *user_serve.FrontUsersResponse) error {
	currentPage := in.CurrentPage
	pageSize := in.Pagesize
	offset := (currentPage - 1) * pageSize
	fmt.Println(offset)
	out.Code = 200
	out.Msg = "登录成功"
	return nil
}
