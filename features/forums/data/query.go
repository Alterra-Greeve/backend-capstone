package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/forums"

	"gorm.io/gorm"
)

type ForumData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) forums.ForumDataInterface {
	return &ForumData{
		DB: db,
	}
}

// Forum
func (u *ForumData) GetAllForum() ([]forums.Forum, error) {
	var forums []forums.Forum
	if err := u.DB.Preload("User").Find(&forums).Error; err != nil {
		return nil, err
	}
	return forums, nil
}

func (u *ForumData) GetAllByPage(page int) ([]forums.Forum, int, error) {
	var forum []forums.Forum

	var total int64
	count := u.DB.Model(&forums.Forum{}).Where("deleted_at IS NULL").Count(&total)
	if count.Error != nil {
		return nil, 0, constant.ErrProductEmpty
	}

	dataforumPerPage := 5
	totalPages := int((total + int64(dataforumPerPage) - 1) / int64(dataforumPerPage))

	tx := u.DB.Model(&Forum{}).Preload("User").Order("GREATEST(COALESCE(last_message_at, '2000-01-01'), created_at) DESC").Offset((page - 1) * dataforumPerPage).Limit(dataforumPerPage).Find(&forum)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetForum
	}
	return forum, totalPages, nil
}

func (u *ForumData) GetForumByID(ID string) (forums.Forum, error) {
	var forum forums.Forum
	if err := u.DB.Model(&Forum{}).Preload("User").Where("id = ?", ID).First(&forum).Error; err != nil {
		return forums.Forum{}, err
	}
	return forum, nil
}

func (u *ForumData) PostForum(forum forums.Forum) error {
	if err := u.DB.Create(&forum).Error; err != nil {
		return err
	}
	return nil
}

func (u *ForumData) UpdateForum(forum forums.EditForum) error {
	var existingForum forums.Forum
	if err := u.DB.Where("id = ?", forum.ID).First(&existingForum).Error; err != nil {
		return err
	}

	err := u.DB.Model(&existingForum).Updates(forum).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *ForumData) DeleteForum(forumID string) error {
	res := u.DB.Begin()

	result := res.Where("forum_id = ?", forumID).Delete(&MessageForum{})
	if result.Error != nil {
		res.Rollback()
		return constant.ErrMessgaeNotFound
	}
	if err := res.Where("id = ?", forumID).Delete(&Forum{}); err.Error != nil {
		res.Rollback()
		return constant.ErrForumNotFound
	} else if err.RowsAffected == 0 {
		res.Rollback()
		return constant.ErrForumNotFound
	}

	return res.Commit().Error
}

func (u *ForumData) GetForumByUserID(userID string, page int) ([]forums.Forum, int, error) {
	var forums []forums.Forum
	var total int64

	dataforumPerPage := 20
	countResult := u.DB.Model(&Forum{}).Where("user_id = ?", userID).Count(&total)
	if countResult.Error != nil {
		return nil, 0, constant.ErrForumNotFound
	}

	if total == 0 {
		return nil, 0, constant.ErrForumNotFound
	}

	totalPages := int((total + int64(dataforumPerPage) - 1) / int64(dataforumPerPage))
	tx := u.DB.Model(&Forum{}).Preload("User").Order("created_at DESC, last_message_at DESC").Offset((page-1)*dataforumPerPage).Limit(dataforumPerPage).Where("user_id = ?", userID).Find(&forums)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetProduct
	}
	return forums, totalPages, nil
}

// Message
func (u *ForumData) PostMessageForum(messageForum forums.MessageForum) error {
	res := u.DB.Begin()
	if err := res.Create(&messageForum).Error; err != nil {
		return err
	}

	if err := res.Model(&Forum{}).Where("id = ? AND deleted_at IS NULL", messageForum.ForumID).Update("last_message_at", messageForum.CreatedAt).Error; err != nil {
		res.Rollback()
		return err
	}

	var forum forums.Forum
	if err := u.DB.Where("id = ? AND deleted_at IS NULL", messageForum.ForumID).First(&forum).Error; err != nil {
		return constant.ErrForumNotFound
	}
	return res.Commit().Error
}

func (u *ForumData) DeleteMessageForum(messageID string) error {
	res := u.DB.Begin()

	if err := res.Where("id = ?", messageID).Delete(&MessageForum{}); err.Error != nil {
		res.Rollback()
		return constant.ErrMessgaeNotFound
	} else if err.RowsAffected == 0 {
		res.Rollback()
		return constant.ErrMessgaeNotFound
	}

	return res.Commit().Error
}

func (r *ForumData) GetMessageForumByID(ID string) (forums.MessageForum, error) {
	var messageForum forums.MessageForum
	if err := r.DB.Where("id = ?", ID).First(&messageForum).Error; err != nil {
		return forums.MessageForum{}, err
	}
	return messageForum, nil
}

func (r *ForumData) UpdateMessageForum(message forums.EditMessage) error {
	var messageForum forums.MessageForum
	if err := r.DB.Where("id = ? AND deleted_at IS NULL", message.ID).First(&messageForum).Error; err != nil {
		return err
	}
	var forum forums.Forum
	if err := r.DB.Where("id = ? AND deleted_at IS NULL", messageForum.ForumID).First(&forum).Error; err != nil {
		return constant.ErrForumNotFound
	}
	err := r.DB.Model(&messageForum).Updates(message).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *ForumData) GetMessagesByForumID(forumID string) ([]forums.MessageForum, error) {
	var messages []forums.MessageForum
	if err := u.DB.Where("forum_id = ? AND deleted_at IS NULL", forumID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (u *ForumData) GetMessagesByForumIDWithPagination(forumID string, page int, pageSize int) ([]forums.MessageForum, error) {
	var messages []forums.MessageForum
	offset := (page - 1) * pageSize
	if err := u.DB.Where("forum_id = ?", forumID).
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
