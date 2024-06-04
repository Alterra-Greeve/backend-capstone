package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/product"
)

type ProductService struct {
	d product.ProductDataInterface
}

func New(d product.ProductDataInterface) product.ProductServiceInterface {
	return &ProductService{
		d: d,
	}
}

func (s *ProductService) Get() ([]product.Product, error) {
	return s.d.Get()
}

func (s *ProductService) GetByPage(page int) ([]product.Product, int, error) {
	if page <= 0 {
		return nil, 0, constant.ErrPageInvalid
	}
	products, total, err := s.d.GetByPage(page)
	if err != nil {
		return nil, 0, err
	}
	if page > total {
		return nil, 0, constant.ErrPageInvalid
	}
	return products, total, nil
}

func (s *ProductService) GetById(id string) (product.Product, error) {
	return s.d.GetById(id)
}

func (s *ProductService) GetByCategory(category string, page int) ([]product.Product, int, error) {
	if page <= 0 {
		return nil, 0, constant.ErrPageInvalid
	}
	products, total, err := s.d.GetByCategory(category, page)
	if err != nil {
		return nil, 0, err
	}
	if page > total {
		return nil, 0, constant.ErrPageInvalid
	}
	return products, total, nil
}

func (s *ProductService) GetByName(name string, page int) ([]product.Product, int, error) {
	if page <= 0 {
		return nil, 0, constant.ErrPageInvalid
	}
	products, total, err := s.d.GetByName(name, page)
	if err != nil {
		return nil, 0, err
	}
	if page > total {
		return nil, 0, constant.ErrPageInvalid
	}
	return products, total, nil
}

func (s *ProductService) Create(product product.Product) error {
	if product.Name == "" {
		return constant.ErrProductNameEmpty
	}
	if product.Price == 0 {
		return constant.ErrProductPriceEmpty
	}
	if product.Images == nil {
		return constant.ErrProductImagesEmpty
	}
	if product.ImpactCategories == nil {
		return constant.ErrProductImpactCategoriesEmpty
	}
	return s.d.Create(product)
}

func (s *ProductService) Update(product product.Product) error {
	if product.ID == "" {
		return constant.ErrProductIDEmpty
	}
	if product.Name == "" && product.Price == 0 && product.Coin == 0 && product.Images == nil && product.ImpactCategories == nil {
		return constant.ErrProductUpdateEmpty
	}

	return s.d.Update(product)
}

func (s *ProductService) Delete(productId string) error {
	return s.d.Delete(productId)
}
