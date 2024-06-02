package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/voucher"
	"log"

	"gorm.io/gorm"
)

type VoucherData struct {
	*gorm.DB
}

func New(db *gorm.DB) voucher.VoucherDataInterface {
	return &VoucherData{
		DB: db,
	}
}

func (u *VoucherData) GetAll() ([]voucher.Voucher, error) {
	var voucher []voucher.Voucher
	err := u.DB.Where("deleted_at IS null").Find(&voucher).Error
	if err != nil {
		return nil, constant.ErrVoucherNotFound
	}
	return voucher, nil
}

func (u *VoucherData) GetByIdVoucher(ID string) (voucher.Voucher, error) {
	var vouchers voucher.Voucher

	if err := u.DB.Where("id = ? AND deleted_at IS null", ID).First(&vouchers).Error; err != nil {
		return voucher.Voucher{}, err
	}
	return vouchers, nil
}

func (u *VoucherData) Create(voucherData voucher.Voucher) error {
	if err := u.DB.Create(&voucherData).Error; err != nil {
		return err
	}
	return nil
}

func (u *VoucherData) Update(voucherData voucher.VoucherEdit) error {
	var existingVoucher voucher.Voucher
	if err := u.DB.Where("id = ? AND deleted_at is null", voucherData.ID).First(&existingVoucher).Error; err != nil {
		return err
	}

	err := u.DB.Model(&existingVoucher).Updates(voucherData).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *VoucherData) Delete(voucher string) error {
	res := u.DB.Begin()

	if err := res.Where("id = ?", voucher).Delete(&Voucher{}); err.Error != nil {
		res.Rollback()
		log.Print(err)
		return constant.ErrVoucherNotFound
	} else if err.RowsAffected == 0 {
		res.Rollback()
		return constant.ErrVoucherNotFound
	}

	return res.Commit().Error
}
