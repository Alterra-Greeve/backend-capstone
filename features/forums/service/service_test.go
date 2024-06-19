package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/forums"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockForumData struct {
	mock.Mock
}

func (m *MockForumData) GetAllForum() ([]forums.Forum, error) {
	args := m.Called()
	return args.Get(0).([]forums.Forum), args.Error(1)
}

func (m *MockForumData) GetAllByPage(page int) ([]forums.Forum, int, error) {
	args := m.Called(page)
	return args.Get(0).([]forums.Forum), args.Int(1), args.Error(2)
}

func (m *MockForumData) GetForumByID(ID string) (forums.Forum, error) {
	args := m.Called(ID)
	return args.Get(0).(forums.Forum), args.Error(1)
}

func (m *MockForumData) PostForum(forum forums.Forum) error {
	args := m.Called(forum)
	return args.Error(0)
}

func (m *MockForumData) UpdateForum(forum forums.EditForum) error {
	args := m.Called(forum)
	return args.Error(0)
}

func (m *MockForumData) DeleteForum(forumID string) error {
	args := m.Called(forumID)
	return args.Error(0)
}

func (m *MockForumData) PostMessageForum(message forums.MessageForum) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockForumData) DeleteMessageForum(messageID string) error {
	args := m.Called(messageID)
	return args.Error(0)
}

func (m *MockForumData) GetMessageForumByID(ID string) (forums.MessageForum, error) {
	args := m.Called(ID)
	return args.Get(0).(forums.MessageForum), args.Error(1)
}

func (m *MockForumData) UpdateMessageForum(message forums.EditMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockForumData) GetForumByUserID(ID string, page int) ([]forums.Forum, int, error) {
	args := m.Called(ID, page)
	return args.Get(0).([]forums.Forum), args.Int(1), args.Error(2)
}

func (m *MockForumData) GetMessagesByForumID(forumID string) ([]forums.MessageForum, error) {
	args := m.Called(forumID)
	return args.Get(0).([]forums.MessageForum), args.Error(1)
}

func (m *MockForumData) GetMessagesByForumIDWithPagination(forumID string, page int, pageSize int) ([]forums.MessageForum, error) {
	args := m.Called(forumID, page, pageSize)
	return args.Get(0).([]forums.MessageForum), args.Error(1)
}

func TestForumService_GetAllForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	expectedForums := []forums.Forum{
		{ID: "1", Title: "Forum 1", Description: "Description 1"},
		{ID: "2", Title: "Forum 2", Description: "Description 2"},
	}

	mockData.On("GetAllForum").Return(expectedForums, nil)

	result, err := service.GetAllForum()

	assert.NoError(t, err)
	assert.Equal(t, expectedForums, result)
	mockData.AssertExpectations(t)
}

func TestForumService_GetAllByPage(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	expectedForums := []forums.Forum{
		{ID: "1", Title: "Forum 1", Description: "Description 1"},
	}
	expectedPageCount := 1

	mockData.On("GetAllByPage", 1).Return(expectedForums, expectedPageCount, nil)

	result, pageCount, err := service.GetAllByPage(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedForums, result)
	assert.Equal(t, expectedPageCount, pageCount)
	mockData.AssertExpectations(t)
}

func TestForumService_GetForumByID(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	expectedForum := forums.Forum{
		ID:          "1",
		Title:       "Forum 1",
		Description: "Description 1",
	}

	mockData.On("GetForumByID", "1").Return(expectedForum, nil)

	result, err := service.GetForumByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedForum, result)
	mockData.AssertExpectations(t)
}

func TestForumService_GetForumByID_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	_, err := service.GetForumByID("")

	assert.Error(t, err)
	assert.Equal(t, constant.ErrGetMessageByID, err)
}

func TestForumService_PostForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	newForum := forums.Forum{
		ID:          "1",
		Title:       "New Forum",
		Description: "New Description",
	}

	mockData.On("PostForum", newForum).Return(nil)

	err := service.PostForum(newForum)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestForumService_PostForum_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	invalidForum := forums.Forum{
		ID:    "1",
		Title: "",
	}

	err := service.PostForum(invalidForum)

	assert.Error(t, err)
	assert.Equal(t, constant.ErrFieldForumCannotBeEmpty, err)
}

func TestForumService_UpdateForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	updateForum := forums.EditForum{
		ID:          "1",
		Title:       "Updated Title",
		Description: "Updated Description",
	}

	mockData.On("UpdateForum", updateForum).Return(nil)

	err := service.UpdateForum(updateForum)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestForumService_UpdateForum_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	invalidForum := forums.EditForum{
		ID: "",
	}

	err := service.UpdateForum(invalidForum)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrGetForumByID, err)

	invalidForum = forums.EditForum{
		ID:          "1",
		Title:       "",
		Description: "",
	}

	err = service.UpdateForum(invalidForum)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrFieldForumCannotBeEmpty, err)

	invalidForum = forums.EditForum{
		ID:          "1",
		Title:       "!!Title",
		Description: "!Description",
	}

	err = service.UpdateForum(invalidForum)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrFieldData, err)
}

func TestForumService_DeleteForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	mockData.On("DeleteForum", "1").Return(nil)

	err := service.DeleteForum("1")

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestForumService_DeleteForum_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	mockData.On("DeleteForum", "").Return(errors.New("some error"))

	err := service.DeleteForum("")
	assert.Error(t, err)

	mockData.AssertExpectations(t)
}

func TestForumService_PostMessageForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	messageForum := forums.MessageForum{
		ForumID: "1",
		Message: "New Message",
	}

	mockData.On("PostMessageForum", messageForum).Return(nil)

	err := service.PostMessageForum(messageForum)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

// func TestForumService_PostMessageForum_Error(t *testing.T) {
// 	mockData := new(MockForumData)
// 	service := New(mockData)

// 	invalidMessage := forums.MessageForum{
// 		ForumID: "",
// 		Message: "",
// 	}

// 	err := service.PostMessageForum(invalidMessage)
// 	assert.Error(t, err)
// 	assert.Equal(t, constant.ErrFieldForumMessageCannotBeEmpty, err)
// }

func TestForumService_GetMessageForumByID(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	expectedMessage := forums.MessageForum{
		ID:      "1",
		ForumID: "1",
		Message: "Test Message",
	}

	mockData.On("GetMessageForumByID", "1").Return(expectedMessage, nil)

	result, err := service.GetMessageForumByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, result)
	mockData.AssertExpectations(t)
}

func TestForumService_GetMessageForumByID_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	_, err := service.GetMessageForumByID("")

	assert.Error(t, err)
	assert.Equal(t, constant.ErrGetMessageByID, err)
}

func TestForumService_UpdateMessageForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	updateMessage := forums.EditMessage{
		ID:      "1",
		Message: "Updated Message",
	}

	mockData.On("UpdateMessageForum", updateMessage).Return(nil)

	err := service.UpdateMessageForum(updateMessage)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestForumService_UpdateMessageForum_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	invalidMessage := forums.EditMessage{
		ID: "",
	}

	err := service.UpdateMessageForum(invalidMessage)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrGetMessageByID, err)

	invalidMessage = forums.EditMessage{
		ID:      "1",
		Message: "",
	}

	err = service.UpdateMessageForum(invalidMessage)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrEditMessage, err)
}

func TestForumService_DeleteMessageForum(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	mockData.On("DeleteMessageForum", "1").Return(nil)

	err := service.DeleteMessageForum("1")

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestForumService_DeleteMessageForum_Error(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	mockData.On("DeleteMessageForum", "").Return(errors.New("some error"))

	err := service.DeleteMessageForum("")
	assert.Error(t, err)

	mockData.AssertExpectations(t)
}

func TestForumService_GetForumByUserID(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	expectedForums := []forums.Forum{
		{ID: "1", Title: "Forum 1", Description: "Description 1"},
	}
	expectedPageCount := 1

	mockData.On("GetForumByUserID", "1", 1).Return(expectedForums, expectedPageCount, nil)

	result, pageCount, err := service.GetForumByUserID("1", 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedForums, result)
	assert.Equal(t, expectedPageCount, pageCount)
	mockData.AssertExpectations(t)
}

func TestForumService_GetMessagesByForumID(t *testing.T) {
	mockData := new(MockForumData)
	service := New(mockData)

	expectedMessages := []forums.MessageForum{
		{ID: "1", ForumID: "1", Message: "Message 1"},
	}

	mockData.On("GetMessagesByForumID", "1").Return(expectedMessages, nil)

	result, err := service.GetMessagesByForumID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, result)
	mockData.AssertExpectations(t)
}
