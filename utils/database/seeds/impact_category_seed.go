package seeds

import (
	ic "backendgreeve/features/impactcategory/data"

	"gorm.io/gorm"
)


func CreateImpactCategory(db *gorm.DB, id string, name string, impact_point int, icon_url string) error {
	impactCategory := ic.ImpactCategory{
		ID:         id,
		Name:       name,
		ImpactPoint: impact_point,
		IconURL:     icon_url,
	}
	return db.Where("id = ?", id).FirstOrCreate(&impactCategory).Error
}
