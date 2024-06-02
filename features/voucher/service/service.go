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

func (s *VoucherService) GetByIdVoucher(id string) (voucher.Voucher, error) {
	if id == "" {
		return voucher.Voucher{}, constant.ErrGetVoucherById
	}
	return s.d.GetByIdVoucher(id)
}

func (s *VoucherService) Create(voucher voucher.Voucher) error {
	if voucher.Name == "" || voucher.Discount == "" || voucher.Code == "" || voucher.ExpiredAt == (time.Time{}) {
		return constant.ErrCreateVoucher
	}
	return s.d.Create(voucher)
}

func (s *VoucherService) Update(voucher voucher.VoucherEdit) error {
	if voucher.ID == "" {
		return constant.ErrGetVoucherById
	}
	return s.d.Update(voucher)
}

func (s *VoucherService) Delete(voucherID string) error {
	return s.d.Delete(voucherID)
}
