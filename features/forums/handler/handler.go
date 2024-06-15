package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/forums"
	"backendgreeve/helper"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ForumHandler struct {
	s forums.ForumServiceInterface
	j helper.JWTInterface
}

func New(s forums.ForumServiceInterface, j helper.JWTInterface) forums.ForumHandlerInterface {
	return &ForumHandler{
		s: s,
		j: j,
	}
}

func (h *ForumHandler) GetAllForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		var totalPages int

		var forums []forums.Forum
		forums, totalPages, err = h.s.GetAllByPage(page)
		metadata := MetadataResponse{
			CurrentPage: page,
			TotalPage:   totalPages,
		}

		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response []ForumGetAllResponse
		for _, f := range forums {
			response = append(response, ForumGetAllResponse{
				ID:          f.ID,
				Title:       f.Title,
				Description: f.Description,
				Author:      Author{ID: f.User.ID, Name: f.User.Name, Username: f.User.Username, Email: f.User.Email, AvatarURL: f.User.AvatarURL},
			})
		}
		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, constant.ForumSuccessGetAll, metadata, response))
	}
}

func (h *ForumHandler) PostForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]
		var forum CreateForumRequest
		c.Bind(&forum)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrCreateForum.Error(), nil))
		}

		forum.Title = strings.TrimSpace(forum.Title)
		forum.Description = strings.TrimSpace(forum.Description)

		forumData := forums.Forum{
			ID:          uuid.New().String(),
			Title:       forum.Title,
			Description: forum.Description,
			UserID:      userId.(string),
		}
		err = h.s.PostForum(forumData)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.ForumSuccessCreate, nil))
	}
}
func (h *ForumHandler) GetForumByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		forumID := c.Param("id")

		forum, err := h.s.GetForumByID(forumID)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ObjectFormatResponse(false, constant.ErrGetForumByID.Error(), nil))
		}

		messages, err := h.s.GetMessagesByForumID(forumID)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.ErrGetMessage.Error(), nil))
		}

		var messageResponses []MessageResponse
		for _, msg := range messages {
			messageResponses = append(messageResponses, MessageResponse{
				ID: msg.ID,
				User: AuthorMessage{
					ID:        msg.User.ID,
					Name:      msg.User.Name,
					Username:  msg.User.Username,
					Email:     msg.User.Email,
					AvatarURL: msg.User.AvatarURL,
				},
				Message: msg.Message,
			})
		}

		forumResponse := ForumGetDetailResponse{
			ID:            forum.ID,
			Title:         forum.Title,
			Description:   forum.Description,
			Author:        Author{ID: forum.User.ID, Name: forum.User.Name, Username: forum.User.Username, Email: forum.User.Email, AvatarURL: forum.User.AvatarURL},
			ForumMessages: messageResponses,
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessGetByID, forumResponse))
	}
}

func (h *ForumHandler) UpdateForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		forumid := c.Param("id")
		existingForum, err := h.s.GetForumByID(forumid)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, string(constant.ErrForumNotFound.Error()), nil))
		}

		if existingForum.UserID != userId {
			return c.JSON(http.StatusForbidden, helper.FormatResponse(false, string(constant.UnatuhorizeForumAndMessage.Error()), nil))
		}
		var forum EditForumRequest
		if err := c.Bind(&forum); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		forum.Title = strings.TrimSpace(forum.Title)
		forum.Description = strings.TrimSpace(forum.Description)
		forumsResponse := forums.EditForum{
			ID:          forumid,
			Title:       forum.Title,
			Description: forum.Description,
		}
		if err := h.s.UpdateForum(forumsResponse); err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessUpdate, nil))
	}
}
func (h *ForumHandler) DeleteForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		adminData := h.j.ExtractAdminToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			return helper.UnauthorizedError(c)
		}

		forumID := c.Param("id")

		err = h.s.DeleteForum(forumID)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.ForumSuccessDelete, nil))
	}
}

func (h *ForumHandler) GetForumByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		var totalPages int
		forums, totalPages, err := h.s.GetForumByUserID(userId, page)
		if err != nil {
			if errors.Is(err, constant.ErrForumNotFound) {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.ErrForumDataNotFound.Error(), nil))
			}
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		metadata := MetadataResponse{
			CurrentPage: page,
			TotalPage:   totalPages,
		}

		var forumsResponse []ForumGetAllResponse

		for _, forum := range forums {
			forumsResponse = append(forumsResponse, ForumGetAllResponse{
				ID:          forum.ID,
				Title:       forum.Title,
				Description: forum.Description,
				Author:      Author{ID: forum.User.ID, Name: forum.User.Name, Username: forum.User.Username, Email: forum.User.Email, AvatarURL: forum.User.AvatarURL},
			})
		}

		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, constant.ForumSuccessGetAll, metadata, forumsResponse))
	}
}

// Message
func (h *ForumHandler) PostMessageForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		var forumMessage CreateMessageForumRequest
		if err := c.Bind(&forumMessage); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		forumData := forums.MessageForum{
			ID:      uuid.New().String(),
			ForumID: forumMessage.ForumID,
			UserID:  userId,
			Message: forumMessage.Messages,
		}
		_, err = h.s.GetForumByID(forumData.ForumID)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.ErrGetForumByID.Error(), nil))
		}

		err = h.s.PostMessageForum(forumData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrCreateMessage.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.MessageSuccessCreate, nil))
	}
}

func (h *ForumHandler) DeleteMessageForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)
		if userId == "" {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, constant.Unauthorized, nil))
		}
		isAdmin := userData[constant.JWT_ROLE] == constant.RoleAdmin

		messageForumID := c.Param("id")
		existingMessageForum, err := h.s.GetMessageForumByID(messageForumID)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, string(constant.ErrGetMessageByID.Error()), nil))
		}

		if existingMessageForum.UserID != userId && !isAdmin {
			return c.JSON(http.StatusForbidden, helper.FormatResponse(false, constant.Unauthorized, nil))
		}

		err = h.s.DeleteMessageForum(messageForumID)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, constant.ErrDeleteMessage.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.MessageSuccessDelete, nil))
	}
}

func (h *ForumHandler) UpdateMessageForum() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		messageId := c.Param("id")
		existingMessage, err := h.s.GetMessageForumByID(messageId)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, string(constant.ErrMessgaeNotFound.Error()), nil))
		}

		if existingMessage.UserID != userId {
			return c.JSON(http.StatusForbidden, helper.FormatResponse(false, string(constant.UnatuhorizeForumAndMessage.Error()), nil))
		}
		var messageForum EditMessageRequest
		if err := c.Bind(&messageForum); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		messageResponse := forums.EditMessage{
			ID:      messageId,
			Message: messageForum.Messages,
		}
		if err := h.s.UpdateMessageForum(messageResponse); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrUpdateMessage.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.MessageSuccessUpdate, nil))
	}

}

func (h *ForumHandler) GetMessageForumByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		messageId := c.Param("id")
		message, err := h.s.GetMessageForumByID(messageId)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ObjectFormatResponse(false, constant.ErrGetMessageByID.Error(), nil))
		}

		messagesResponse := MessageResponse{
			ID:      message.ID,
			User:    AuthorMessage{ID: message.UserID, Name: message.User.Name, Username: message.User.Username, Email: message.User.Email, AvatarURL: message.User.AvatarURL},
			Message: message.Message,
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.MessageSuccessGetByID, messagesResponse))
	}
}
