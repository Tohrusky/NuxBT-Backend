package db

import (
	"fmt"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/dal/query"
)

// CreateUser 新建用户
func CreateUser(user *model.User) (err error) {
	q := query.User
	err = q.Create(user)
	return err
}

// PatchUser 更新用户信息，根据 userID 和 details 更新用户信息
func PatchUser(userID int32, details *model.User) (err error) {
	q := query.User
	ResultInfo, err := q.Where(q.UserID.Eq(userID)).Updates(details)
	if err != nil {
		return err
	}
	if ResultInfo.RowsAffected == 0 {
		return fmt.Errorf("no rows affected, nothing will be updated, userID: %v", userID)
	}
	return nil
}

// GetUserByEmail 根据 email 获取用户
func GetUserByEmail(email string) (user *model.User, err error) {
	q := query.User
	user, err = q.Where(q.Email.Eq(email)).First()
	return user, err
}

// GetUserByID 根据 userID 获取用户
func GetUserByID(userID int32) (user *model.User, err error) {
	q := query.User
	user, err = q.Where(q.UserID.Eq(userID)).First()
	return user, err
}

// GetUserRolesByID 根据 userID 获取用户角色列表
func GetUserRolesByID(userID int32) (roles []string, err error) {
	q := query.UserRole
	user, err := q.Where(q.UserID.Eq(userID)).Find()
	if err != nil {
		return nil, err
	}
	for _, v := range user {
		roles = append(roles, v.Role)
	}
	return roles, nil
}
