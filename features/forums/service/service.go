package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/forums"
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

func (s *ForumService) GetForumByID(ID string) (forums.Forum, error) {
	if ID == "" {
		return forums.Forum{}, constant.ErrGetMessageByID
	}
	return s.d.GetForumByID(ID)
}

func (s *ForumService) PostForum(forum forums.Forum) error {
	if forum.Title == "" || forum.Description == "" {
		return constant.ErrCreateMessage
	}
	return s.d.PostForum(forum)
}

func (s *ForumService) UpdateForum(forum forums.EditForum) error {
	if forum.ID == "" {
		return constant.ErrEditForum
	}
	if forum.Title == "" && forum.Description == "" {
		return constant.ErrEditForum
	}
	return s.d.UpdateForum(forum)
}

func (s *ForumService) DeleteForum(forumData forums.Forum) error {
	if forumData.ID == "" {
		return constant.ErrDeleteMessage
	}
	return s.d.DeleteForum(forumData)
}

func (s *ForumService) PostMessageForum(messageForum forums.MessageForum) error {
	if messageForum.ForumID == "" || messageForum.Message == "" {
		return constant.ErrCreateMessage
	}
	return s.d.PostMessageForum(messageForum)
}

func (s *ForumService) DeleteMessageForum(messageForum forums.MessageForum) error {
	if messageForum.ID == "" {
		return constant.ErrDeleteMessage
	}
	return s.d.DeleteMessageForum(messageForum)
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
		return constant.ErrEditForum
	}
	return s.d.UpdateMessageForum(message)
}

func (s *ForumService) GetForumByUserID(ID string) (forums.Forum, error) {
	if ID == "" {
		return forums.Forum{}, constant.ErrGetMessageByID
	}
	return s.d.GetForumByUserID(ID)
}

func (s *ForumService) GetMessagesByForumID(forumID string) ([]forums.MessageForum, error) {
	return s.d.GetMessagesByForumID(forumID)
}

func (s *ForumService) GetMessagesByForumIDWithPagination(forumID string, page int, pageSize int) ([]forums.MessageForum, error) {
	return s.d.GetMessagesByForumIDWithPagination(forumID, page, pageSize)
}
