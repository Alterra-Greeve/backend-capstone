package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/impactcategory"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockImpactData struct {
	mock.Mock
}

func (m *MockImpactData) GetAll() ([]impactcategory.ImpactCategory, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]impactcategory.ImpactCategory), args.Error(1)
}

func (m *MockImpactData) GetByID(id string) (impactcategory.ImpactCategory, error) {
	args := m.Called(id)
	return args.Get(0).(impactcategory.ImpactCategory), args.Error(1)
}

func (m *MockImpactData) Create(impact impactcategory.ImpactCategory) error {
	args := m.Called(impact)
	return args.Error(0)
}

func (m *MockImpactData) Update(impact impactcategory.ImpactCategoryUpdate) error {
	args := m.Called(impact)
	return args.Error(0)
}
func (m *MockImpactData) Delete(impact impactcategory.ImpactCategory) error {
	args := m.Called(impact)
	return args.Error(0)
}

func TestForumService_GetAll(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	expectedForums := []impactcategory.ImpactCategory{
		{ID: "1", Name: "Forum 1", ImpactPoint: 1},
		{ID: "2", Name: "Forum 2", ImpactPoint: 2},
	}

	mockData.On("GetAll").Return(expectedForums, nil)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedForums, result)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_GetAll_Error(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	expectedError := errors.New("failed to get categories")

	mockData.On("GetAll").Return(nil, expectedError)

	result, err := service.GetAll()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_GetByID(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	categoryID := "123"
	expectedCategory := impactcategory.ImpactCategory{ID: categoryID, Name: "Category 1", ImpactPoint: 1}

	mockData.On("GetByID", categoryID).Return(expectedCategory, nil)

	result, err := service.GetByID(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)
	mockData.AssertExpectations(t)
}
func TestImpactCategoryService_GetByID_Error(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	categoryID := "123"
	expectedError := errors.New("failed to get category by ID")

	mockData.On("GetByID", categoryID).Return(impactcategory.ImpactCategory{}, expectedError)

	result, err := service.GetByID(categoryID)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, impactcategory.ImpactCategory{}, result)
	mockData.AssertExpectations(t)
}
func TestImpactCategoryService_Update(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	updateData := impactcategory.ImpactCategoryUpdate{ID: "123", Name: "Updated Category", ImpactPoint: 5}

	mockData.On("Update", updateData).Return(nil)

	err := service.Update(updateData)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}
func TestImpactCategoryService_Update_Error(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	updateData := impactcategory.ImpactCategoryUpdate{ID: "123", Name: "Updated Category", ImpactPoint: 5}
	expectedError := errors.New("update error")

	mockData.On("Update", updateData).Return(expectedError)

	err := service.Update(updateData)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockData.AssertExpectations(t)
}
func TestImpactCategoryService_Create(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	newCategory := impactcategory.ImpactCategory{ID: "1", Name: "New Category", ImpactPoint: 10, IconURL: "http://example.com/icon.png"}

	mockData.On("Create", newCategory).Return(nil)

	err := service.Create(newCategory)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_Create_Error(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	newCategory := impactcategory.ImpactCategory{ID: "1", Name: "New Category", ImpactPoint: 1, IconURL: "http://example.com/icon.png"}
	expectedError := errors.New("create error")

	mockData.On("Create", newCategory).Return(expectedError)

	err := service.Create(newCategory)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_Delete(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	categoryToDelete := impactcategory.ImpactCategory{ID: "123"}

	mockData.On("Delete", categoryToDelete).Return(nil)

	err := service.Delete(categoryToDelete)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}
func TestImpactCategoryService_Delete_Error(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	categoryToDelete := impactcategory.ImpactCategory{ID: "123"}
	expectedError := errors.New("delete error")

	mockData.On("Delete", categoryToDelete).Return(expectedError)

	err := service.Delete(categoryToDelete)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_Create_InvalidInput(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	newCategory := impactcategory.ImpactCategory{ID: "1", Name: "", ImpactPoint: 10, IconURL: "http://example.com/icon.png"}

	err := service.Create(newCategory)

	assert.Error(t, err)
	assert.Equal(t, constant.ErrImpactCategoryField, err)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_GetByID_InvalidID(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	categoryID := ""

	_, err := service.GetByID(categoryID)

	assert.Error(t, err)
	assert.Equal(t, constant.ErrImpactCategoryNotFound, err)
	mockData.AssertExpectations(t)
}

func TestImpactCategoryService_Update_EmptyData(t *testing.T) {
	mockData := new(MockImpactData)
	service := New(mockData)

	updateData := impactcategory.ImpactCategoryUpdate{ID: "123"}

	err := service.Update(updateData)

	assert.Error(t, err)
	assert.Equal(t, constant.ErrImpactCategoryFieldUpdate, err)
	mockData.AssertExpectations(t)
}
