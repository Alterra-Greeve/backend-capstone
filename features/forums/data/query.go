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

func (u *ForumData) GetForumByID(ID string) (forums.Forum, error) {
	var forum forums.Forum
	if err := u.DB.Preload("User").Where("id = ?", ID).First(&forum).Error; err != nil {
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
func (u *ForumData) DeleteForum(forum forums.Forum) error {
	result := u.DB.Table("forums").Delete(&forum)
	if result.Error != nil {
		return constant.ErrDeleteForum
	}

	if result.RowsAffected == 0 {
		return constant.ErrForumNotFound
	}
	return nil
}

func (u *ForumData) GetForumByUserID(ID string) (forums.Forum, error) {
	var forum forums.Forum
	if err := u.DB.Where("user_id = ?", ID).First(&forum).Error; err != nil {
		return forums.Forum{}, err
	}
	return forum, nil
}

// Message
func (u *ForumData) PostMessageForum(messageForum forums.MessageForum) error {
	if err := u.DB.Create(&messageForum).Error; err != nil {
		return err
	}
	return nil
}

func (u *ForumData) DeleteMessageForum(messageForum forums.MessageForum) error {
	result := u.DB.Table("message_forums").Delete(&messageForum)
	if result.Error != nil {
		return constant.ErrDeleteForum
	}

	if result.RowsAffected == 0 {
		return constant.ErrMessgaeNotFound
	}
	return nil
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
	if err := r.DB.Where("id = ?", message.ID).Find(&messageForum).Error; err != nil {
		return err
	}

	err := r.DB.Model(&messageForum).Updates(message).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *ForumData) GetMessagesByForumID(forumID string) ([]forums.MessageForum, error) {
	var messages []forums.MessageForum
	if err := u.DB.Where("forum_id = ?", forumID).Find(&messages).Error; err != nil {
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
