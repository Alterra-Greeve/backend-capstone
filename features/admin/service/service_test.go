package service

import (
	"fmt"
	"testing"

	"backendgreeve/constant"
	"backendgreeve/features/admin"
	"backendgreeve/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminData struct {
	mock.Mock
	Users   map[string]admin.Admin
	Err     error
	Updated bool
}

type MockJWT struct {
	mock.Mock
	Token string
	Err   error
}

type MockMailer struct {
	mock.Mock
}

func (m *MockAdminData) Login(user admin.Admin) (admin.Admin, error) {
	args := m.Called(user)
	return args.Get(0).(admin.Admin), args.Error(1)
}

func (m *MockAdminData) Update(user admin.AdminUpdate) (admin.Admin, error) {
	args := m.Called(user)
	return args.Get(0).(admin.Admin), args.Error(1)
}

func (m *MockAdminData) Delete(user admin.Admin) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAdminData) GetAdminByID(id string) (admin.Admin, error) {
	args := m.Called(id)
	return args.Get(0).(admin.Admin), args.Error(1)
}

func (m *MockAdminData) IsEmailExist(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

func (m *MockJWT) GenerateAdminJWT(admin helper.AdminJWT) (string, error) {
	args := m.Called(admin)
	return args.String(0), args.Error(1)
}

func (m *MockJWT) ExtractAdminToken(token *jwt.Token) map[string]interface{} {
	args := m.Called(token)
	return args.Get(0).(map[string]interface{})
}

func (m *MockJWT) GenerateAdminToken(admin helper.AdminJWT) string {
	args := m.Called(admin)
	return args.String(0)
}

func (m *MockJWT) GenerateUserJWT(user helper.UserJWT) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockJWT) GenerateUserToken(user helper.UserJWT) string {
	args := m.Called(user)
	return args.String(0)
}

func (m *MockJWT) ExtractUserToken(token *jwt.Token) map[string]interface{} {
	args := m.Called(token)
	return args.Get(0).(map[string]interface{})
}

func (m *MockJWT) ValidateToken(id string) (*jwt.Token, error) {
	args := m.Called(id)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

func (m *MockMailer) Send(to, subject string) error {
	args := m.Called(to, subject)
	return args.Error(0)
}

func TestLoginEmptyInput(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	adminUser := admin.Admin{
		Email:    "",
		Password: "",
	}

	_, err := service.Login(adminUser)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrEmptyEmailandPasswordAdmin, err)
}

func TestLoginErrorJWT(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := &MockJWT{Err: fmt.Errorf("JWT error")}
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	adminUser := admin.Admin{
		Email:    "john@example.com",
		Password: "password123",
	}

	mockData.On("Login", adminUser).Return(admin.Admin{}, nil)
	mockJWT.On("GenerateAdminJWT", mock.Anything).Return("", fmt.Errorf("JWT error"))

	_, err := service.Login(adminUser)
	assert.Error(t, err)
	assert.Equal(t, "JWT error", err.Error())
}

func TestAdminService_Login(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	adminUser := admin.Admin{
		Email:    "test@example.com",
		Password: "password123",
	}

	returnedAdmin := admin.Admin{
		ID:       "1",
		Name:     "Test Admin",
		Email:    "test@example.com",
		Username: "testadmin",
	}

	mockData.On("Login", adminUser).Return(returnedAdmin, nil)
	mockJWT.On("GenerateAdminJWT", mock.Anything).Return("mockToken", nil)

	result, err := service.Login(adminUser)

	assert.NoError(t, err)
	assert.Equal(t, "mockToken", result.Token)
	mockData.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

// func TestAdminService_Update(t *testing.T) {
// 	mockData := new(MockAdminData)
// 	mockJWT := new(MockJWT)
// 	mockMailer := new(MockMailer)
// 	service := New(mockData, mockJWT, mockMailer)

// 	adminUpdate := admin.AdminUpdate{
// 		ID:       "1",
// 		Name:     "New Name",
// 		Username: "newusername",
// 		Email:    "newemail@example.com",
// 		Password: "newpassword",
// 	}

// 	updatedAdmin := admin.Admin{
// 		ID:       "1",
// 		Name:     "New Name",
// 		Username: "newusername",
// 		Email:    "newemail@example.com",
// 	}

// 	mockData.On("Update", adminUpdate).Return(updatedAdmin, nil)
// 	mockJWT.On("GenerateAdminJWT", mock.Anything).Return("newMockToken", nil)

// 	result, err := service.Update(adminUpdate)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "newMockToken", result.Token)
// 	mockData.AssertExpectations(t)
// 	mockJWT.AssertExpectations(t)
// }

// func TestAdminService_UpdateInvalidName(t *testing.T) {
// 	mockData := new(MockAdminData)
// 	mockJWT := new(MockJWT)
// 	mockMailer := new(MockMailer)
// 	service := New(mockData, mockJWT, mockMailer)

// 	adminUpdate := admin.AdminUpdate{
// 		ID:   "1",
// 		Name: "Invalid Name!", // Assume special characters are invalid
// 	}

// 	_, err := service.Update(adminUpdate)
// 	assert.Error(t, err)
// 	assert.Equal(t, constant.ErrFieldData, err)
// }

func TestAdminService_UpdateTrimmedFields(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	adminUpdate := admin.AdminUpdate{
		ID:       "1",
		Name:     "  New Name  ",
		Email:    " newemail@example.com ",
		Username: " newusername ",
	}

	updatedAdmin := admin.Admin{
		ID:       "1",
		Name:     "New Name",
		Username: "newusername",
		Email:    "newemail@example.com",
	}

	mockData.On("Update", mock.Anything).Return(updatedAdmin, nil)
	mockJWT.On("GenerateAdminJWT", mock.Anything).Return("newMockToken", nil)

	result, err := service.Update(adminUpdate)

	assert.NoError(t, err)
	assert.Equal(t, "newMockToken", result.Token)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, "newusername", result.Username)
	assert.Equal(t, "newemail@example.com", result.Email)
	mockData.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestAdminService_UpdateError(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	// Test with empty ID
	adminUpdate := admin.AdminUpdate{
		ID: "",
	}

	_, err := service.Update(adminUpdate)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrUpdateUser, err)

	// Test with all empty fields
	adminUpdate = admin.AdminUpdate{
		ID: "1",
	}

	_, err = service.Update(adminUpdate)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrEditAdmin, err)
}

func TestAdminService_Delete(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	adminUser := admin.Admin{
		ID: "1",
	}

	mockData.On("Delete", adminUser).Return(nil)

	err := service.Delete(adminUser)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestAdminService_DeleteError(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	// Test with empty ID
	adminUser := admin.Admin{
		ID: "",
	}

	err := service.Delete(adminUser)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrDeleteUser, err)
}

func TestAdminService_GetAdminData(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	adminUser := admin.Admin{
		ID: "1",
	}

	returnedAdmin := admin.Admin{
		ID:       "1",
		Name:     "Test Admin",
		Email:    "test@example.com",
		Username: "testadmin",
	}

	mockData.On("GetAdminByID", "1").Return(returnedAdmin, nil)

	result, err := service.GetAdminData(adminUser)

	assert.NoError(t, err)
	assert.Equal(t, returnedAdmin, result)
	mockData.AssertExpectations(t)
}

func TestAdminService_GetAdminDataError(t *testing.T) {
	mockData := new(MockAdminData)
	mockJWT := new(MockJWT)
	mockMailer := new(MockMailer)
	service := New(mockData, mockJWT, mockMailer)

	// Test with empty ID
	adminUser := admin.Admin{
		ID: "",
	}

	_, err := service.GetAdminData(adminUser)
	assert.Error(t, err)
	assert.Equal(t, constant.ErrGetDataAdmin, err)
}
