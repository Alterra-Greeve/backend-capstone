package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChallengeData struct {
	mock.Mock
}

func (m *MockChallengeData) GetAllForUser(userId string) ([]challenges.Challenge, error) {
	args := m.Called(userId)
	return args.Get(0).([]challenges.Challenge), args.Error(1)
}

func (m *MockChallengeData) GetChallengeParticipant(challengeId string) (int, error) {
	args := m.Called(challengeId)
	return args.Int(0), args.Error(1)
}

func (m *MockChallengeData) GetByID(challengeId string) (challenges.Challenge, error) {
	args := m.Called(challengeId)
	return args.Get(0).(challenges.Challenge), args.Error(1)
}

func (m *MockChallengeData) Swipe(userId string, challengeId string, challengeType string) error {
	args := m.Called(userId, challengeId, challengeType)
	return args.Error(0)
}

func (m *MockChallengeData) Create(challenge challenges.Challenge) error {
	args := m.Called(challenge)
	return args.Error(0)
}

func (m *MockChallengeData) GetAllForAdmin(page int) ([]challenges.Challenge, int, error) {
	args := m.Called(page)
	return args.Get(0).([]challenges.Challenge), args.Int(1), args.Error(2)
}

func (m *MockChallengeData) GetUserParticipate(userId string, status string) ([]challenges.ChallengeConfirmation, error) {
	args := m.Called(userId, status)
	return args.Get(0).([]challenges.ChallengeConfirmation), args.Error(1)
}

func (m *MockChallengeData) Update(challenge challenges.Challenge) error {
	args := m.Called(challenge)
	return args.Error(0)
}

func (m *MockChallengeData) IsUserParticipate(userId string, challengeId string) (error, bool) {
	args := m.Called(userId, challengeId)
	return args.Error(0), args.Bool(1)
}

func (m *MockChallengeData) Delete(challengeId string) error {
	args := m.Called(challengeId)
	return args.Error(0)
}

func (m *MockChallengeData) GetChallengeForUserByID(challengeConfirmationId string) (challenges.ChallengeConfirmation, error) {
	args := m.Called(challengeConfirmationId)
	return args.Get(0).(challenges.ChallengeConfirmation), args.Error(1)
}

func (m *MockChallengeData) EditChallengeForUserByID(challengeId string, images []string) error {
	args := m.Called(challengeId, images)
	return args.Error(0)
}

func (m *MockChallengeData) AddToLogs(userId string, challengeId string, status string) error {
	args := m.Called(userId, challengeId, status)
	return args.Error(0)
}

func (m *MockChallengeData) InsertCoinAndExpUser(challengeConfirmationId string) error {
	args := m.Called(challengeConfirmationId)
	return args.Error(0)
}

func TestChallengeService_GetAllForUser(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	expectedChallenges := []challenges.Challenge{
		{ID: "1", Title: "Challenge 1"},
		{ID: "2", Title: "Challenge 2"},
	}

	mockData.On("GetAllForUser", userId).Return(expectedChallenges, nil)

	result, err := service.GetAllForUser(userId)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	mockData.AssertExpectations(t)
}

func TestChallengeService_GetChallengeParticipant(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"
	expectedParticipants := 10

	mockData.On("GetChallengeParticipant", challengeId).Return(expectedParticipants, nil)

	result, err := service.GetChallengeParticipant(challengeId)

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipants, result)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	challengeId := "456"
	challengeType := "Diterima"

	mockData.On("Swipe", userId, challengeId, challengeType).Return(nil)

	err := service.Swipe(userId, challengeId, challengeType)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe_Reject(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	challengeId := "456"
	challengeType := "Ditolak"

	mockData.On("Swipe", userId, challengeId, challengeType).Return(nil)

	err := service.Swipe(userId, challengeId, challengeType)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe_InvalidEmptyUserId(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := ""
	challengeId := "456"
	challengeType := "Diterima"

	err := service.Swipe(userId, challengeId, challengeType)

	assert.EqualError(t, err, constant.ErrChallengeFieldSwipe.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe_InvalidEmptyChallengeId(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	challengeId := ""
	challengeType := "Diterima"

	err := service.Swipe(userId, challengeId, challengeType)

	assert.EqualError(t, err, constant.ErrChallengeFieldSwipe.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe_InvalidEmptyChallengeType(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	challengeId := "456"
	challengeType := ""

	err := service.Swipe(userId, challengeId, challengeType)

	assert.EqualError(t, err, constant.ErrChallengeFieldSwipe.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe_InvalidChallengeType(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	challengeId := "456"
	challengeType := "InvalidType"

	err := service.Swipe(userId, challengeId, challengeType)

	assert.EqualError(t, err, constant.ErrChallengeType.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Create(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	newChallenge := challenges.Challenge{
		ID:               "1",
		Title:            "New Challenge",
		Description:      "Description",
		Exp:              100,
		Coin:             50,
		DateStart:        time.Now(),
		DateEnd:          time.Now().Add(24 * time.Hour),
		Difficulty:       "Easy",
		ImpactCategories: []challenges.ChallengeImpactCategory{{ID: "11111111"}},
	}

	mockData.On("Create", newChallenge).Return(nil)

	err := service.Create(newChallenge)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestChallengeService_GetByID(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"
	expectedChallenge := challenges.Challenge{
		ID:    "123",
		Title: "Existing Challenge",
	}

	mockData.On("GetByID", challengeId).Return(expectedChallenge, nil)

	result, err := service.GetByID(challengeId)

	assert.NoError(t, err)
	assert.Equal(t, expectedChallenge, result)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Delete(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"

	mockData.On("GetByID", challengeId).Return(challenges.Challenge{ID: challengeId}, nil)
	mockData.On("Delete", challengeId).Return(nil)

	err := service.Delete(challengeId)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Delete_ChallengeNotFound(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "456"

	mockData.On("GetByID", challengeId).Return(challenges.Challenge{}, nil)

	err := service.Delete(challengeId)

	assert.EqualError(t, err, constant.ErrChallengeNotFound.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_EditChallengeForUserByID(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"
	images := []string{"image1.jpg", "image2.jpg"}

	mockData.On("EditChallengeForUserByID", challengeId, images).Return(nil)

	err := service.EditChallengeForUserByID(challengeId, images)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestChallengeService_GetAllForAdmin(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	page := 1
	expectedChallenges := []challenges.Challenge{
		{ID: "1", Title: "Challenge 1"},
		{ID: "2", Title: "Challenge 2"},
	}
	totalPages := 1

	mockData.On("GetAllForAdmin", page).Return(expectedChallenges, totalPages, nil)

	result, total, err := service.GetAllForAdmin(page)

	assert.NoError(t, err)
	assert.Equal(t, expectedChallenges, result)
	assert.Equal(t, totalPages, total)
	mockData.AssertExpectations(t)
}
func TestChallengeService_GetUserParticipate(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	status := "completed"
	expectedConfirmations := []challenges.ChallengeConfirmation{
		{ID: "1", Status: "completed"},
		{ID: "2", Status: "completed"},
	}

	mockData.On("GetUserParticipate", userId, status).Return(expectedConfirmations, nil)

	result, err := service.GetUserParticipate(userId, status)

	assert.NoError(t, err)
	assert.Equal(t, expectedConfirmations, result)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Update(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challenge := challenges.Challenge{
		ID:          "123",
		Title:       "Updated Challenge",
		Description: "Updated Description",
		Exp:         100,
		Coin:        50,
		DateStart:   time.Now(),
		DateEnd:     time.Now().Add(24 * time.Hour),
		Difficulty:  "Medium",
		ImpactCategories: []challenges.ChallengeImpactCategory{
			{ID: "11111111"},
		},
	}

	mockData.On("GetByID", challenge.ID).Return(challenges.Challenge{ID: challenge.ID}, nil)
	mockData.On("Update", challenge).Return(nil)

	err := service.Update(challenge)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestChallengeService_Update_InvalidEmptyFields(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challenge := challenges.Challenge{
		ID: "123",
	}

	err := service.Update(challenge)

	assert.EqualError(t, err, constant.ErrChallengeFieldCreate.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Update_ChallengeNotFound(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challenge := challenges.Challenge{
		ID:               "456",
		Title:            "Updated Challenge",
		Description:      "Updated Description",
		Exp:              100,
		Coin:             50,
		DateStart:        time.Now(),
		DateEnd:          time.Now().Add(24 * time.Hour),
		Difficulty:       "Medium",
		ImpactCategories: []challenges.ChallengeImpactCategory{{ID: "11111111"}},
	}

	mockData.On("GetByID", challenge.ID).Return(challenges.Challenge{}, nil)

	err := service.Update(challenge)

	assert.EqualError(t, err, constant.ErrChallengeNotFound.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_GetChallengeForUserByID(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeConfirmationId := "123"
	expectedChallengeConfirmation := challenges.ChallengeConfirmation{
		ID: "123",
	}

	mockData.On("GetChallengeForUserByID", challengeConfirmationId).Return(expectedChallengeConfirmation, nil)

	result, err := service.GetChallengeForUserByID(challengeConfirmationId)

	assert.NoError(t, err)
	assert.Equal(t, expectedChallengeConfirmation, result)
	mockData.AssertExpectations(t)
}

func TestChallengeService_GetByID_Error(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"
	expectedError := errors.New("get by id error")

	mockData.On("GetByID", challengeId).Return(challenges.Challenge{}, expectedError)

	_, err := service.GetByID(challengeId)

	assert.EqualError(t, err, expectedError.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Create_Error(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	newChallenge := challenges.Challenge{
		ID:               "1",
		Title:            "New Challenge",
		Description:      "Description",
		Exp:              100,
		Coin:             50,
		DateStart:        time.Now(),
		DateEnd:          time.Now().Add(24 * time.Hour),
		Difficulty:       "Easy",
		ImpactCategories: []challenges.ChallengeImpactCategory{{ID: "11111111"}},
	}
	expectedError := errors.New("create error")

	mockData.On("Create", newChallenge).Return(expectedError)

	err := service.Create(newChallenge)

	assert.EqualError(t, err, expectedError.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Delete_Error(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"
	expectedError := errors.New("delete error")

	mockData.On("GetByID", challengeId).Return(challenges.Challenge{ID: challengeId}, nil)
	mockData.On("Delete", challengeId).Return(expectedError)

	err := service.Delete(challengeId)

	assert.EqualError(t, err, expectedError.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_EditChallengeForUserByID_Error(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeId := "123"
	images := []string{"image1.jpg", "image2.jpg"}
	expectedError := errors.New("edit challenge for user by id error")

	mockData.On("EditChallengeForUserByID", challengeId, images).Return(expectedError)

	err := service.EditChallengeForUserByID(challengeId, images)

	assert.EqualError(t, err, expectedError.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_GetChallengeForUserByID_Error(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	challengeConfirmationId := "123"
	expectedError := errors.New("get challenge for user by id error")

	mockData.On("GetChallengeForUserByID", challengeConfirmationId).Return(challenges.ChallengeConfirmation{}, expectedError)

	_, err := service.GetChallengeForUserByID(challengeConfirmationId)

	assert.EqualError(t, err, expectedError.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Swipe_Error(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	userId := "123"
	challengeId := "456"
	challengeType := "Diterima"
	expectedError := errors.New("swipe error")

	mockData.On("Swipe", userId, challengeId, challengeType).Return(expectedError)

	err := service.Swipe(userId, challengeId, challengeType)

	assert.EqualError(t, err, expectedError.Error())
	mockData.AssertExpectations(t)
}

func TestChallengeService_Create_InvalidEmptyTitle(t *testing.T) {
	mockData := new(MockChallengeData)
	service := New(mockData)

	newChallenge := challenges.Challenge{
		ID:               "1",
		Title:            "",
		Description:      "Description",
		Exp:              100,
		Coin:             50,
		DateStart:        time.Now(),
		DateEnd:          time.Now().Add(24 * time.Hour),
		Difficulty:       "Easy",
		ImpactCategories: []challenges.ChallengeImpactCategory{{ID: "11111111"}},
	}

	err := service.Create(newChallenge)

	assert.EqualError(t, err, constant.ErrChallengeFieldCreate.Error())
	mockData.AssertExpectations(t)
}
