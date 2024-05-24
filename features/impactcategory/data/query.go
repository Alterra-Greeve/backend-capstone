package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/impactcategory"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImpactCategoryData struct {
	db *gorm.DB
}

func New(db *gorm.DB) impactcategory.ImpactCategoryDataInterface {
	return &ImpactCategoryData{
		db: db,
	}
}

func (icd *ImpactCategoryData) GetAll() ([]impactcategory.ImpactCategory, error) {
	var categories []impactcategory.ImpactCategory
	err := icd.db.Find(&categories).Error
	if err != nil {
		return nil, constant.ErrImpactCategoryNotFound
	}
	return categories, nil
}

func (icd *ImpactCategoryData) GetByID(ID string) (impactcategory.ImpactCategory, error) {
	var category impactcategory.ImpactCategory
	err := icd.db.First(&category, "id = ?", ID).Error
	if err != nil {
		return impactcategory.ImpactCategory{}, constant.ErrImpactCategoryNotFound
	}
	return category, nil
}

func (icd *ImpactCategoryData) Create(category impactcategory.ImpactCategory) error {
	category.ID = uuid.New().String()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	err := icd.db.Create(&category).Error
	if err != nil {
		return constant.ErrCreateImpactCategory
	}
	return nil
}

func (icd *ImpactCategoryData) Update(category impactcategory.ImpactCategoryUpdate) error {
	err := icd.db.Table("impact_categories").Where("id = ?", category.ID).Updates(category).Error
	if err != nil {
		return constant.ErrUpdateImpactCategory
	}
	return nil
}

func (icd *ImpactCategoryData) Delete(category impactcategory.ImpactCategory) error {
	result := icd.db.Table("impact_categories").Delete(&category)
	if result.Error != nil {
		return constant.ErrDeleteImpactCategory
	}
	if result.RowsAffected == 0 {
		return constant.ErrImpactCategoryNotFound
	}
	return nil
}
