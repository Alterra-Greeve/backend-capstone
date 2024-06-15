package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/forums"
	"backendgreeve/helper"
)

type ForumService struct {
	d forums.ForumDataInterface
}

func New(data forums.ForumDataInterface) forums.ForumServiceInterface {
	return &ForumService{
		d: data,
	}
}

func (s *ForumService) GetAllForum() ([]forums.Forum, error) {
	return s.d.GetAllForum()
}

func (s *ForumService) GetAllByPage(page int) ([]forums.Forum, int, error) {
	return s.d.GetAllByPage(page)
}

func (s *ForumService) GetForumByID(ID string) (forums.Forum, error) {
	if ID == "" {
		return forums.Forum{}, constant.ErrGetMessageByID
	}
	return s.d.GetForumByID(ID)
}

func (s *ForumService) PostForum(forum forums.Forum) error {
	if forum.Title == "" || forum.Description == "" {
		return constant.ErrFieldForumCannotBeEmpty
	}
	if !helper.IsValidInput(forum.Title) || !helper.IsValidInput(forum.Description) {
		return constant.ErrFieldData
	}
	return s.d.PostForum(forum)
}

func (s *ForumService) UpdateForum(forum forums.EditForum) error {
	if forum.ID == "" {
		return constant.ErrGetForumByID
	}
	if forum.Title == "" && forum.Description == " " {
		return constant.ErrEditForum
	}

	if !helper.IsValidInput(forum.Title) && !helper.IsValidInput(forum.Description) {
		return constant.ErrFieldData
	}
	return s.d.UpdateForum(forum)
}

func (s *ForumService) DeleteForum(forumID string) error {
	return s.d.DeleteForum(forumID)
}

func (s *ForumService) PostMessageForum(messageForum forums.MessageForum) error {
	if messageForum.ForumID == "" {
		return constant.ErrGetForumByID
	}
	if messageForum.Message == "" {
		return constant.ErrFieldForumMessageCannotBeEmpty
	}

	if !helper.IsValidInput(messageForum.Message) {
		return constant.ErrFieldData
	}
	return s.d.PostMessageForum(messageForum)
}

func (s *ForumService) DeleteMessageForum(productId string) error {
	return s.d.DeleteMessageForum(productId)
}

func (s *ForumService) GetMessageForumByID(ID string) (forums.MessageForum, error) {
	if ID == "" {
		return forums.MessageForum{}, constant.ErrGetMessageByID
	}
	return s.d.GetMessageForumByID(ID)
}

func (s *ForumService) UpdateMessageForum(message forums.EditMessage) error {
	if message.ID == "" {
		return constant.ErrGetMessageByID
	}
	if message.Message == "" {
		return constant.ErrEditMessage
	}

	if !helper.IsValidInput(message.Message) {
		return constant.ErrFieldData
	}
	return s.d.UpdateMessageForum(message)
}

func (s *ForumService) GetForumByUserID(ID string, page int) ([]forums.Forum, int, error) {
	return s.d.GetForumByUserID(ID, page)
}

func (s *ForumService) GetMessagesByForumID(forumID string) ([]forums.MessageForum, error) {
	return s.d.GetMessagesByForumID(forumID)
}

func (s *ForumService) GetMessagesByForumIDWithPagination(forumID string, page int, pageSize int) ([]forums.MessageForum, error) {
	return s.d.GetMessagesByForumIDWithPagination(forumID, page, pageSize)
}
