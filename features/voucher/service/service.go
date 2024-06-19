package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/voucher"
	"time"
)

type VoucherService struct {
	d voucher.VoucherDataInterface
}

func New(data voucher.VoucherDataInterface) voucher.VoucherServiceInterface {
	return &VoucherService{
		d: data,
	}
}

func (s *VoucherService) GetAll() ([]voucher.Voucher, error) {
	return s.d.GetAll()
}

func (s *VoucherService) GetVoucherUsed(voucherId string) (int, error) {
	return s.d.GetVoucherUsed(voucherId)
}

func (s *VoucherService) GetByIdVoucher(id string) (voucher.Voucher, error) {
	if id == "" {
		return voucher.Voucher{}, constant.ErrGetVoucherById
	}
	return s.d.GetByIdVoucher(id)
}

func (s *VoucherService) Create(voucher voucher.Voucher) error {
	if voucher.Name == "" || voucher.Discount == "" || voucher.Description == "" || voucher.Code == "" || voucher.ExpiredAt == (time.Time{}) {
		return constant.ErrVoucherField
	}

	if len(voucher.Code) != 9 {
		return constant.ErrCodeVoucher
	}
	return s.d.Create(voucher)
}

func (s *VoucherService) Update(vouchers voucher.VoucherEdit) error {
	if vouchers.ID == "" {
		return constant.ErrGetVoucherById
	}

	var existingVoucher voucher.Voucher
	if vouchers.Code != existingVoucher.Code {
		if len(vouchers.Code) != 9 {
			return constant.ErrCodeVoucher
		}
	}
	return s.d.Update(vouchers)
}

func (s *VoucherService) Delete(voucherID string) error {
	return s.d.Delete(voucherID)
}
