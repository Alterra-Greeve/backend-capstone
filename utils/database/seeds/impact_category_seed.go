package seeds

import (
	ic "backendgreeve/features/impactcategory/data"

	"gorm.io/gorm"
)

func CreateImpactCategory(db *gorm.DB, id string, name string, impact_point int, description string, image_url string, icon_url string) error {
	impactCategory := ic.ImpactCategory{
		ID:          id,
		Name:        name,
		ImpactPoint: impact_point,
		Description: description,
		ImageURL:    image_url,
		IconURL:     icon_url,
	}
	return db.Where("id = ?", id).FirstOrCreate(&impactCategory).Error
}
