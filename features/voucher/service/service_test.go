package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/voucher"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVoucherData struct {
	mock.Mock
}

func (m *MockVoucherData) GetAll() ([]voucher.Voucher, error) {
	args := m.Called()
	return args.Get(0).([]voucher.Voucher), args.Error(1)
}

func (m *MockVoucherData) GetByIdVoucher(ID string) (voucher.Voucher, error) {
	args := m.Called(ID)
	return args.Get(0).(voucher.Voucher), args.Error(1)
}

func (m *MockVoucherData) Create(v voucher.Voucher) error {
	args := m.Called(v)
	return args.Error(0)
}

func (m *MockVoucherData) Update(v voucher.VoucherEdit) error {
	args := m.Called(v)
	return args.Error(0)
}

func (m *MockVoucherData) Delete(voucherID string) error {
	args := m.Called(voucherID)
	return args.Error(0)
}

func TestVoucherService_GetAll(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	expectedVouchers := []voucher.Voucher{
		{ID: "1", Name: "Voucher 1", Description: "Description 1", ExpiredAt: time.Now(), Discount: "10%", Code: "ABC123456"},
		{ID: "2", Name: "Voucher 2", Description: "Description 2", ExpiredAt: time.Now(), Discount: "20%", Code: "DEF987654"},
	}

	mockData.On("GetAll").Return(expectedVouchers, nil)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedVouchers, result)
	mockData.AssertExpectations(t)
}

func TestVoucherService_GetByIdVoucher_Success(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	expectedVoucher := voucher.Voucher{
		ID:          "1",
		Name:        "Test Voucher",
		Description: "Test Description",
		ExpiredAt:   time.Now(),
		Discount:    "15%",
		Code:        "XYZ987654",
	}

	mockData.On("GetByIdVoucher", "1").Return(expectedVoucher, nil)

	result, err := service.GetByIdVoucher("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedVoucher, result)
	mockData.AssertExpectations(t)
}

func TestVoucherService_GetByIdVoucher_EmptyID(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	expectedError := constant.ErrGetVoucherById

	result, err := service.GetByIdVoucher("")

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, voucher.Voucher{}, result)
}

func TestVoucherService_Create_Success(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	voucherToCreate := voucher.Voucher{
		Name:        "New Voucher",
		Description: "New Voucher Description",
		ExpiredAt:   time.Now().AddDate(0, 0, 7),
		Discount:    "25%",
		Code:        "EXIST1234",
	}

	mockData.On("Create", voucherToCreate).Return(nil)

	err := service.Create(voucherToCreate)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}

func TestVoucherService_Create_InvalidFields(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	invalidVoucher := voucher.Voucher{
		Name: "Invalid Voucher",
	}

	expectedError := constant.ErrVoucherField

	err := service.Create(invalidVoucher)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockData.AssertNotCalled(t, "Create")
}

// func TestVoucherService_Update_Success(t *testing.T) {
// 	mockData := new(MockVoucherData)
// 	service := New(mockData)

// 	existingVoucher := voucher.Voucher{
// 		ID:          "1",
// 		Name:        "Existing Voucher",
// 		Description: "Existing Voucher Description",
// 		ExpiredAt:   time.Now().AddDate(0, 0, 14),
// 		Discount:    "30%",
// 		Code:        "EXIST1234",
// 	}

// 	voucherEdit := voucher.VoucherEdit{
// 		ID:   "1",
// 		Code: "NEWCODE45",
// 	}

// 	mockData.On("GetByIdVoucher", "1").Return(existingVoucher, nil)

// 	mockData.On("Update", voucherEdit).Return(nil)

// 	err := service.Update(voucherEdit)

// 	assert.NoError(t, err)
// 	mockData.AssertExpectations(t)
// }

func TestVoucherService_Update_InvalidCodeLength(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	existingVoucher := voucher.Voucher{
		ID:          "1",
		Name:        "Existing Voucher",
		Description: "Existing Voucher Description",
		ExpiredAt:   time.Now().AddDate(0, 0, 14),
		Discount:    "30%",
		Code:        "EXIST12345",
	}

	voucherEdit := voucher.VoucherEdit{
		ID:   "1",
		Code: "INVALID",
	}

	mockData.On("GetByIdVoucher", "1").Return(existingVoucher, nil)

	expectedError := constant.ErrCodeVoucher

	err := service.Update(voucherEdit)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockData.AssertNotCalled(t, "Update")
}

func TestVoucherService_Update_EmptyID(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	voucherEdit := voucher.VoucherEdit{
		ID:   "",
		Code: "NEWCODE456",
	}

	expectedError := constant.ErrGetVoucherById

	err := service.Update(voucherEdit)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockData.AssertNotCalled(t, "Update")
}

func TestVoucherService_Delete(t *testing.T) {
	mockData := new(MockVoucherData)
	service := New(mockData)

	voucherID := "1"

	mockData.On("Delete", voucherID).Return(nil)

	err := service.Delete(voucherID)

	assert.NoError(t, err)
	mockData.AssertExpectations(t)
}
