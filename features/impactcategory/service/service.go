package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/impactcategory"
)

type ImpactCategoryService struct {
	d impactcategory.ImpactCategoryDataInterface
}

func New(d impactcategory.ImpactCategoryDataInterface) impactcategory.ImpactCategoryServiceInterface {
	return &ImpactCategoryService{
		d: d,
	}
}

func (ics *ImpactCategoryService) GetAll() ([]impactcategory.ImpactCategory, error) {
	return ics.d.GetAll()
}

func (ics *ImpactCategoryService) GetByID(ID string) (impactcategory.ImpactCategory, error) {
	if ID == "" {
		return impactcategory.ImpactCategory{}, constant.ErrImpactCategoryNotFound
	}
	return ics.d.GetByID(ID)
}

func (ics *ImpactCategoryService) Create(category impactcategory.ImpactCategory) error {
	if category.Name == "" || category.ImpactPoint == 0 || category.IconURL == "" {
		return constant.ErrImpactCategoryField
	}
	return ics.d.Create(category)
}

func (ics *ImpactCategoryService) Update(category impactcategory.ImpactCategoryUpdate) error {
	if category.ID == "" {
		return constant.ErrImpactCategoryNotFound
	}
	if category.Name == "" && category.ImpactPoint == 0 && category.IconURL == "" {
		return constant.ErrImpactCategoryFieldUpdate
	}

	return ics.d.Update(category)
}

func (ics *ImpactCategoryService) Delete(category impactcategory.ImpactCategory) error {
	return ics.d.Delete(category)
}

