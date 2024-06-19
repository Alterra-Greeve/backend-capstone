package voucher

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Voucher struct {
	ID          string
	Name        string
	Code        string
	Discount    string
	Description string
	ExpiredAt   time.Time
	CreatedAt   time.Time
}

type VoucherEdit struct {
	ID          string
	Name        string
	Code        string
	Discount    string
	Description string
	ExpiredAt   time.Time
	UpdateAt    time.Time
}

type VoucherHandlerInterface interface {
	GetAll() echo.HandlerFunc
	GetByIdVoucher() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type VoucherServiceInterface interface {
	GetAll() ([]Voucher, error)
	GetByIdVoucher(ID string) (Voucher, error)
	Create(Voucher) error
	Update(VoucherEdit) error
	Delete(voucherID string) error
	GetVoucherUsed(voucherId string) (int, error)
}

type VoucherDataInterface interface {
	GetAll() ([]Voucher, error)
	GetByIdVoucher(ID string) (Voucher, error)
	Create(Voucher) error
	Update(VoucherEdit) error
	Delete(voucherID string) error
	GetVoucherUsed(voucherId string) (int, error)
}
