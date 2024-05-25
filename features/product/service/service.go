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

func (s *ProductService) GetById(id string) (product.Product, error) {
    return s.d.GetById(id)
}

func (s *ProductService) GetByCategoryID(categoryID string) ([]product.Product, error) {
    return s.d.GetByCategoryID(categoryID)
}

func (s *ProductService) GetByName(name string) ([]product.Product, error) {
    return s.d.GetByName(name)
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
    return s.d.Update(product)
}

func (s *ProductService) Delete(product product.Product) error {
    return s.d.Delete(product)
}

