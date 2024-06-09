package seeds

import (
	ic "backendgreeve/features/impactcategory/data"

	"gorm.io/gorm"
)

func CreateImpactCategory(db *gorm.DB, id string, name string, impact_point int, icon_url, image_url, description string) error {
	impactCategory := ic.ImpactCategory{
		ID:          id,
		Name:        name,
		ImpactPoint: impact_point,
		IconURL:     icon_url,
		ImageURL:    image_url,
		Description: description,
	}
	return db.Where("id = ?", id).FirstOrCreate(&impactCategory).Error
}
