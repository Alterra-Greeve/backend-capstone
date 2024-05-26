package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/forums"
	"backendgreeve/helper"
	"net/http"
	"strconv"

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

		forums, err := h.s.GetAllForum()
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.ObjectFormatResponse(false, err.Error(), nil))
		}

		var forumsResponse []ForumGetAllResponse

		for _, forum := range forums {
			forumsResponse = append(forumsResponse, ForumGetAllResponse{
				ID:          forum.ID,
				Title:       forum.Title,
				Description: forum.Description,
				Author:      Author{ID: forum.Author.ID, Name: forum.Author.Name, Username: forum.Author.Username, Email: forum.Author.Email},
			})
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessGetAll, forumsResponse))
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
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		forumData := forums.Forum{
			ID:          uuid.New().String(),
			Title:       forum.Title,
			Description: forum.Description,
			AuthorID:    userId.(string),
		}
		err = h.s.PostForum(forumData)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.ForumSuccessCreate, nil))
	}
}
func (h *ForumHandler) GetForumByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		forumid := c.Param("id")
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 1 {
			page = 1
		}
		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
		if err != nil || pageSize < 1 {
			pageSize = 10
		}

		forums, err := h.s.GetForumByID(forumid)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.ObjectFormatResponse(false, err.Error(), nil))
		}

		messages, err := h.s.GetMessagesByForumID(forumid)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.ObjectFormatResponse(false, err.Error(), nil))
		}

		var messageResponses []MessageResponse
		for _, msg := range messages {
			messageResponses = append(messageResponses, MessageResponse{
				ID:      msg.ID,
				UserID:  msg.UserID,
				Message: msg.Message,
			})
		}

		forumsResponse := ForumGetDetailResponse{
			ID:            forums.ID,
			Title:         forums.Title,
			Description:   forums.Description,
			Author:        Author{ID: forums.Author.ID, Name: forums.Author.Name, Username: forums.Author.Username, Email: forums.Author.Email},
			ForumMessages: messageResponses,
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessGetByID, forumsResponse))
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

		if existingForum.AuthorID != userId {
			return c.JSON(http.StatusForbidden, helper.FormatResponse(false, string(constant.UnatuhorizeForumAndMessage.Error()), nil))
		}
		var forum EditForumRequest
		if err := c.Bind(&forum); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

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
			helper.UnauthorizedError(c)
		}

		adminData := h.j.ExtractUserToken(token)
		role := adminData[constant.JWT_ROLE]
		if role != constant.RoleAdmin {
			helper.UnauthorizedError(c)
		}

		forumid := c.Param("id")
		var forums forums.Forum
		forums.ID = forumid
		err = h.s.DeleteForum(forums)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.ForumSuccessDelete, nil))
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
		err = h.s.PostMessageForum(forumData)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.MessageSuccessCreate, nil))
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

		messageForumID := c.Param("id")
		existingMessageForum, err := h.s.GetMessageForumByID(messageForumID)
		if err != nil {
			return helper.InternalServerError(c)
		}

		if existingMessageForum.UserID != userId {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, constant.Unauthorized, nil))
		}

		err = h.s.DeleteMessageForum(existingMessageForum)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
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
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, string(constant.ErrForumNotFound.Error()), nil))
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
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
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
			return c.JSON(helper.ConvertResponseCode(err), helper.ObjectFormatResponse(false, err.Error(), nil))
		}

		forumsResponse := MessageResponse{
			ID:      message.ID,
			Message: message.Message,
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessGetByID, forumsResponse))
	}
}
