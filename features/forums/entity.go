package forums

import (
	"backendgreeve/features/users"
	"time"

	"github.com/labstack/echo/v4"
)

type Forum struct {
	ID          string
	Title       string
	Description string
	UserID      string
	User        users.User
	CreatedAt   time.Time
}

type MessageForum struct {
	ID        string
	UserID    string
	ForumID   string
	Message   string
	CreatedAt time.Time
}

type EditForum struct {
	ID          string
	Title       string
	UserID      string
	Description string
	UpdatedAt   time.Time
}

type EditMessage struct {
	ID        string
	UserID    string
	Message   string
	UpdatedAt time.Time
}

type ForumHandlerInterface interface {
	// User
	// # Forum
	GetAllForum() echo.HandlerFunc
	GetForumByID() echo.HandlerFunc
	PostForum() echo.HandlerFunc

	//
	UpdateForum() echo.HandlerFunc // Unsure
	DeleteForum() echo.HandlerFunc

	// # Message Forum
	PostMessageForum() echo.HandlerFunc
	DeleteMessageForum() echo.HandlerFunc
	UpdateMessageForum() echo.HandlerFunc
	GetMessageForumByID() echo.HandlerFunc
	GetForumByUserID() echo.HandlerFunc
}

type ForumServiceInterface interface {
	// User
	// # Forum
	GetAllForum() ([]Forum, error)
	GetForumByID(ID string) (Forum, error)
	PostForum(Forum) error
	GetForumByUserID(ID string) (Forum, error)

	//
	UpdateForum(EditForum) error //unsure
	DeleteForum(Forum) error

	// # Message Forum
	PostMessageForum(MessageForum) error
	DeleteMessageForum(MessageForum) error
	GetMessageForumByID(ID string) (MessageForum, error)
	UpdateMessageForum(EditMessage) error
	GetMessagesByForumID(forumID string) ([]MessageForum, error)
	GetMessagesByForumIDWithPagination(forumID string, page int, pageSize int) ([]MessageForum, error)
}

type ForumDataInterface interface {
	// User
	// # Forum
	GetAllForum() ([]Forum, error)
	GetForumByID(ID string) (Forum, error)
	PostForum(Forum) error
	GetForumByUserID(ID string) (Forum, error)

	//
	UpdateForum(EditForum) error //unsure
	DeleteForum(Forum) error

	// # Message Forum
	PostMessageForum(MessageForum) error
	DeleteMessageForum(MessageForum) error
	GetMessageForumByID(ID string) (MessageForum, error)
	UpdateMessageForum(EditMessage) error
	GetMessagesByForumID(forumID string) ([]MessageForum, error)
	GetMessagesByForumIDWithPagination(forumID string, page int, pageSize int) ([]MessageForum, error)
}
