package seeds

import (
	"backendgreeve/utils/database/seed"

	"gorm.io/gorm"
)

func Seeds() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "CreateAdmin1",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "16930c07-bdb5-49d2-8a81-32591833241b", "admin", "admin", "admin@greeve.store", "admin")
			},
		},
		{
			Name: "CreateAdmin2",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "14adafd7-de6c-4586-a35e-3cf17ef3d351", "admin2", "admin2", "admin2@greeve.store", "admin2")
			},
		},
	}
	return seeds
}
