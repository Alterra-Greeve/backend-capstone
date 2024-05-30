package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/adminusers"
	"backendgreeve/features/users"
	userData "backendgreeve/features/users/data"
	userbyAdmin "backendgreeve/features/users/data"

	"gorm.io/gorm"
)

type UserDatabyAdmin struct {
	*gorm.DB
}

func New(db *gorm.DB) adminusers.UserDatabyAdminInterface {
	return &UserDatabyAdmin{
		DB: db,
	}
}

func (u *UserDatabyAdmin) GetAllUsers() ([]users.User, error) {
	var users []users.User
	res := u.DB.Find(&users).Where("deleted_at IS NULL")
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (u *UserDatabyAdmin) GetAllByPage(page int) ([]users.User, int, error) {
	var forum []users.User

	var total int64
	count := u.DB.Model(&users.User{}).Count(&total)
	if count.Error != nil {
		return nil, 0, constant.ErrUserDataEmpty
	}

	dataforumPerPage := 20
	totalPages := int((total + int64(dataforumPerPage) - 1) / int64(dataforumPerPage))

	tx := u.DB.Model(&userData.User{}).Offset((page - 1) * dataforumPerPage).Limit(dataforumPerPage).Find(&forum)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetUser
	}
	if tx.RowsAffected == 0 {
		return nil, 0, constant.ErrUserNotFound
	}
	return forum, totalPages, nil
}

func (u *UserDatabyAdmin) GetUserByID(id string) (users.User, error) {
	var users users.User
	res := u.DB.Model(&userbyAdmin.User{}).Where(" id = ? AND deleted_at IS NULL", id).First(&users)
	if res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

func (u *UserDatabyAdmin) UpdateUser(user adminusers.UpdateUserByAdmin) error {
	var existingUser users.User
	if err := u.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		return err
	}

	err := u.DB.Model(&existingUser).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDatabyAdmin) DeleteUser(userID string) error {
	res := u.DB.Begin()

	if err := res.Where("id = ?", userID).Delete(&userData.User{}); err.Error != nil {
		res.Rollback()
		return constant.ErrUserIDNotFound
	} else if err.RowsAffected == 0 {
		res.Rollback()
		return constant.ErrUserIDNotFound
	}

	return res.Commit().Error
}
