package seeds

import (
	challenge "backendgreeve/features/challenges/data"
	product "backendgreeve/features/product/data"
	"backendgreeve/utils/database/seed"
	"time"

	"github.com/google/uuid"
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
		{
			Name: "CreateImpactCategorySustainable",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "a8089bc1-de1c-4eaf-a1fd-a2a6bc3028e2", "Sustainable", 50, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryEcoFriendly",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "1350c308-05e7-4d49-a5ad-3e92b00ecba9", "Eco-Friendly", 40, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryRecyclable",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "078817ea-4fcf-4eeb-96b2-f10961826469", "Recyclable", 30, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryRenewable",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "c0c85c1d-9bb3-4e4d-a2d4-13fd458f15bc", "Renewable", 45, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryBiodegradable",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "30fc3682-a674-42e1-b420-e16fcd0c85a5", "Biodegradable", 35, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryEnergyEfficient",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "0c30941d-a7fc-438f-8ab7-5bba4f1cf57f", "Energy Efficient", 55, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryLowCarbon",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "3655727a-e807-4fc1-bfad-6b7516c89b13", "Low Carbon", 60, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryWaterSaving",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "19ecd259-58f0-4b3c-ba6b-e16f03317048", "Water Saving", 50, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryZeroWaste",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "8559dcee-ffcf-4aa0-b77d-98dd82f4e96b", "Zero Waste", 70, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryOrganic",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "2d9c8fd2-8104-470f-a8d3-4ff34993fabd", "Organic", 40, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryNonToxic",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "4b1e1acd-570d-47d6-8d6a-ea2a277882cb", "Non-Toxic", 50, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryCompostable",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "1cd52acc-f272-4aa8-8abd-e64fb288f64c", "Compostable", 45, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryReusable",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "3ecbf571-79b6-452b-8626-c45e1a2ab671", "Reusable", 60, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryUpcycled",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "14eb3066-5f5c-459b-a84a-c99cbbd114f9", "Upcycled", 35, "https://www.example.com")
			},
		},
		{
			Name: "CreateImpactCategoryFairTrade",
			Run: func(db *gorm.DB) error {
				return CreateImpactCategory(db, "e7f1ae18-bc65-4f29-a8c7-dd0c30ab3977", "Fair Trade", 50, "https://www.example.com")
			},
		},
		//Product Seeds Start
		{
			Name: "CreateProduct1",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "7dcf0f20-d1c6-4d8a-8cf5-de60f5c420be",
					Name:        "Reusable Straw",
					Description: "A set of stainless steel straws with a cleaning brush.",
					Price:       15000,
					Coin:        5,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7dcf0f20-d1c6-4d8a-8cf5-de60f5c420be",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/af266a14-78bd-4076-817d-10ebec2403698158DFd-L5L.__AC_SX300_SY300_QL70_FMwebp_.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "7dcf0f20-d1c6-4d8a-8cf5-de60f5c420be",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/3e4c422e-163d-4916-bfd6-d9bb1768a57d71-3ljQvVDL._AC_SX679_.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7dcf0f20-d1c6-4d8a-8cf5-de60f5c420be",
						ImpactCategoryID: "3ecbf571-79b6-452b-8626-c45e1a2ab671",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct2",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "7c45085a-2466-4f5e-9ccd-bdd139751dac",
					Name:        "Biodegradable Toothbrush",
					Description: "Sikat gigi yang terbuat dari bambu dengan bulu sikat yang dapat terurai secara alami.",
					Price:       10000,
					Coin:        7,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7c45085a-2466-4f5e-9ccd-bdd139751dac",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b27c09ea-ae32-4cfc-85e6-0c2e962305e8818HQZIZgZL.__AC_SY300_SX300_QL70_FMwebp_.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "7c45085a-2466-4f5e-9ccd-bdd139751dac",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/ae7d23a8-3710-4961-a4c5-d510252fe9fc71nbjZF6TxL._AC_SX679_.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7c45085a-2466-4f5e-9ccd-bdd139751dac",
						ImpactCategoryID: "30fc3682-a674-42e1-b420-e16fcd0c85a5",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct3",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "7fbf0a60-57bf-419f-900a-a993d3168cec",
					Name:        "Organic Cotton Bag",
					Description: "An eco-friendly bag made from 100% organic cotton.",
					Price:       45000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7fbf0a60-57bf-419f-900a-a993d3168cec",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/6b391e80-0aef-491e-82be-17716d7047aamessy_hair_amazingy_bag.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7fbf0a60-57bf-419f-900a-a993d3168cec",
						ImpactCategoryID: "2d9c8fd2-8104-470f-a8d3-4ff34993fabd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "7fbf0a60-57bf-419f-900a-a993d3168cec",
						ImpactCategoryID: "3ecbf571-79b6-452b-8626-c45e1a2ab671",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "7fbf0a60-57bf-419f-900a-a993d3168cec",
						ImpactCategoryID: "e7f1ae18-bc65-4f29-a8c7-dd0c30ab3977",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct4",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "b3b4ae7a-ed2e-4163-bf71-1a00b089c8a3",
					Name:        "Eco-Friendly Notebook",
					Description: "A notebook made from recycled paper and eco-friendly ink.",
					Price:       40000,
					Coin:        8,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "b3b4ae7a-ed2e-4163-bf71-1a00b089c8a3",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/c75dc5f6-7510-4d25-90a9-36e61112b1eaECO-notebooks.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "b3b4ae7a-ed2e-4163-bf71-1a00b089c8a3",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/627c1804-50fe-4797-ae83-90f11fbd7276ECO-notebooks-2-980x980.webp",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "b3b4ae7a-ed2e-4163-bf71-1a00b089c8a3",
						ImpactCategoryID: "1350c308-05e7-4d49-a5ad-3e92b00ecba9",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "b3b4ae7a-ed2e-4163-bf71-1a00b089c8a3",
						ImpactCategoryID: "078817ea-4fcf-4eeb-96b2-f10961826469",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct5",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "7dc84308-0330-41f1-8072-ca635e5055c8",
					Name:        "Reusable Water Bottle",
					Description: "A durable water bottle made from stainless steel.",
					Price:       35000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7dc84308-0330-41f1-8072-ca635e5055c8",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/d8707133-71f4-4b7e-aa76-6356f0c455d471t8zZuv+aL._AC_SY300_SX300_.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "7dc84308-0330-41f1-8072-ca635e5055c8",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/cc9059cf-4e00-46bc-9e21-7d2534b5d2d3718B2kwwspL._AC_SX679_.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7dc84308-0330-41f1-8072-ca635e5055c8",
						ImpactCategoryID: "3ecbf571-79b6-452b-8626-c45e1a2ab671",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct6",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "cc8d605c-db48-40ec-b62c-826a6aa5f9a2",
					Name:        "Solar Powered Charger",
					Description: "A portable solar-powered charger for your devices.",
					Price:       150000,
					Coin:        15,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "cc8d605c-db48-40ec-b62c-826a6aa5f9a2",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/55c70474-5251-421b-a4b2-7fcd9bb44272518fd648-46dc-4fbb-a2bc-15b0e57db458.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "cc8d605c-db48-40ec-b62c-826a6aa5f9a2",
						ImpactCategoryID: "c0c85c1d-9bb3-4e4d-a2d4-13fd458f15bc",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "cc8d605c-db48-40ec-b62c-826a6aa5f9a2",
						ImpactCategoryID: "0c30941d-a7fc-438f-8ab7-5bba4f1cf57f",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct7",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "809db141-0071-478e-bf49-9666968c82b5",
					Name:        "Bamboo Cutlery Set",
					Description: "A set of reusable cutlery made from bamboo.",
					Price:       20000,
					Coin:        6,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "809db141-0071-478e-bf49-9666968c82b5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2ea547b6-f906-44b9-af13-80705afd4528images.jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "809db141-0071-478e-bf49-9666968c82b5",
						ImpactCategoryID: "3ecbf571-79b6-452b-8626-c45e1a2ab671",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "809db141-0071-478e-bf49-9666968c82b5",
						ImpactCategoryID: "30fc3682-a674-42e1-b420-e16fcd0c85a5",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct8",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "5d502c20-9e83-4e84-9d93-db901a3508b5",
					Name:        "Fair Trade Coffee",
					Description: "Organic and fair trade coffee beans.",
					Price:       35000,
					Coin:        12,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "5d502c20-9e83-4e84-9d93-db901a3508b5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/58e8b9ea-57c0-4bf2-8f36-be23fb2ccccbimages (1).jpeg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "5d502c20-9e83-4e84-9d93-db901a3508b5",
						ImpactCategoryID: "2d9c8fd2-8104-470f-a8d3-4ff34993fabd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "5d502c20-9e83-4e84-9d93-db901a3508b5",
						ImpactCategoryID: "e7f1ae18-bc65-4f29-a8c7-dd0c30ab3977",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct9",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "f5e9bae2-1e04-4e2f-8ed5-319849ba8b39",
					Name:        "Compostable Plates",
					Description: "Plates made from compostable materials",
					Price:       17000,
					Coin:        8,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "f5e9bae2-1e04-4e2f-8ed5-319849ba8b39",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/a2720d40-a2f0-4bb2-9d96-853d39039047images (2).jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "f5e9bae2-1e04-4e2f-8ed5-319849ba8b39",
						ImpactCategoryID: "1cd52acc-f272-4aa8-8abd-e64fb288f64c",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct10",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "08989b73-113e-4161-8047-9642db33fea3",
					Name:        "Energy Efficient Bulbs",
					Description: "LED bulbs that save energy and last longer.",
					Price:       58000,
					Coin:        20,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "08989b73-113e-4161-8047-9642db33fea3",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/4b5e6cab-59dd-4d81-988e-e1d390cb9b0c24ba49e1-c01c-4bff-bf1c-e26b0e3e660a.jpg.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "08989b73-113e-4161-8047-9642db33fea3",
						ImpactCategoryID: "0c30941d-a7fc-438f-8ab7-5bba4f1cf57f",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "08989b73-113e-4161-8047-9642db33fea3",
						ImpactCategoryID: "3655727a-e807-4fc1-bfad-6b7516c89b13",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct11",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "bc601a3e-513f-4ba5-b1ac-171d27489b90",
					Name:        "Upcycled Drawer",
					Description: "Drawer made from upcycled materials.",
					Price:       30,
					Coin:        900000,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "bc601a3e-513f-4ba5-b1ac-171d27489b90",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/63355302-4eb9-4585-aabb-fb2cf36db1d3upcycle-furniture-ideas.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "bc601a3e-513f-4ba5-b1ac-171d27489b90",
						ImpactCategoryID: "14eb3066-5f5c-459b-a84a-c99cbbd114f9",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct12",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "e5da3310-6ec6-4b43-81c2-a6eb5a0538c8",
					Name:        "Low Carbon Sneakers",
					Description: "Sneakers with a low carbon footprint.",
					Price:       25,
					Coin:        750000,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: products.ID,
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/505e2f53-bffd-4d0b-9986-48c57aa51431seoarticle7desktop1-1645003305937.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: products.ID,
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/4b31515d-913e-4562-b17c-a398c7bf287aseoarticle7desktop2-1645003305957.webp",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "3655727a-e807-4fc1-bfad-6b7516c89b13",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct13",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "a10882e9-ef38-4796-bc25-f56ca2374fed",
					Name:        "Non-Toxic Paint",
					Description: "Paint made from non-toxic and natural ingredients.",
					Price:       18,
					Coin:        80000,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: products.ID,
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b354682c-0d94-4b60-8a5b-81ec3e68d71fimages (3).jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "4b1e1acd-570d-47d6-8d6a-ea2a277882cb",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "1350c308-05e7-4d49-a5ad-3e92b00ecba9",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct14",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "0e1c5843-a637-4a92-bfa0-7c7342f51846",
					Name:        "Recycled Paper",
					Description: "Paper made from 100% recycled materials.",
					Price:       12000,
					Coin:        5,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: products.ID,
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/499b5dd0-1df5-4f50-a909-9cc297dd5dac5d1c5b4a544aa5b3ec360cc6db50b318.jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "078817ea-4fcf-4eeb-96b2-f10961826469",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "1350c308-05e7-4d49-a5ad-3e92b00ecba9",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct15",
			Run: func(db *gorm.DB) error {
				products := product.Product{
					ID:          "de45d430-da60-4701-a51a-ffda02ba69c5",
					Name:        "Reusable Grocery Bags",
					Description: "Strong and durable grocery bags made from recycled materials.",
					Price:       24000,
					Coin:        7,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: products.ID,
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2603e2f9-26f8-44dc-8e8f-92f604f00dffTrashy_Smart_Bag_-_Unfolded.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "3ecbf571-79b6-452b-8626-c45e1a2ab671",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        products.ID,
						ImpactCategoryID: "078817ea-4fcf-4eeb-96b2-f10961826469",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		// Product Ends
		// Challenge Starts
		{
			Name: "CreateChallenge1",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "5fe8f4f6-43ee-4420-8969-bf1a0e19260c",
					Title:       "Plant a Tree",
					Difficulty:  "Medium",
					Description: "Plant a tree in your backyard or a community park",
					Exp:         100,
					Coin:        50,
					DateStart:   time.Date(2024, 6, 10, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 6, 17, 9, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "a8089bc1-de1c-4eaf-a1fd-a2a6bc3028e2",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "30fc3682-a674-42e1-b420-e16fcd0c85a5",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge2",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "ba2621ac-6f99-43cb-ab8a-7db626b7c4e4",
					Title:       "Zero Waste Week",
					Difficulty:  "Hard",
					Description: "Produce no waste for a week and recycle everything possible",
					Exp:         200,
					Coin:        100,
					DateStart:   time.Date(2024, 7, 1, 8, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 7, 8, 8, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "8559dcee-ffcf-4aa0-b77d-98dd82f4e96b",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "078817ea-4fcf-4eeb-96b2-f10961826469",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge3",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "5086c59a-e2f9-44d0-bcc6-f7fdb58f07a4",
					Title:       "Community Clean-Up",
					Difficulty:  "Easy",
					Description: "Join a local community clean-up event",
					Exp:         50,
					Coin:        25,
					DateStart:   time.Date(2024, 8, 15, 10, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 8, 22, 10, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "1350c308-05e7-4d49-a5ad-3e92b00ecba9",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "078817ea-4fcf-4eeb-96b2-f10961826469",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge4",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "aa013dc7-4fb3-4cf1-9cc0-e1a0ad60414d",
					Title:       "Bicycle to Work",
					Difficulty:  "Medium",
					Description: "Commute to work by bicycle for a week",
					Exp:         150,
					Coin:        75,
					DateStart:   time.Date(2024, 9, 5, 7, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 9, 12, 7, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "3655727a-e807-4fc1-bfad-6b7516c89b13",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "0c30941d-a7fc-438f-8ab7-5bba4f1cf57f",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge5",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "5550e03b-15c6-43c9-997c-4086adc8b573",
					Title:       "Energy Saving Month",
					Difficulty:  "Hard",
					Description: "Reduce your household energy consumption by 20% for a month",
					Exp:         300,
					Coin:        150,
					DateStart:   time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 10, 31, 23, 59, 59, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "0c30941d-a7fc-438f-8ab7-5bba4f1cf57f",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "19ecd259-58f0-4b3c-ba6b-e16f03317048",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge6",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "9b21b3e5-46ff-42a1-92cc-4969c078ff83",
					Title:       "DIY Eco-friendly Products",
					Difficulty:  "Medium",
					Description: "Create and use three DIY eco-friendly household products",
					Exp:         120,
					Coin:        60,
					DateStart:   time.Date(2024, 11, 10, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 11, 17, 9, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "4b1e1acd-570d-47d6-8d6a-ea2a277882cb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "14eb3066-5f5c-459b-a84a-c99cbbd114f9",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge7",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "85dd4189-9d2a-42f9-8b26-b03daa7e5869",
					Title:       "Green Commute Challenge",
					Difficulty:  "Hard",
					Description: "Use only public transport or walk for all your commutes for two weeks",
					Exp:         250,
					Coin:        125,
					DateStart:   time.Date(2024, 12, 01, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 12, 15, 8, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "3655727a-e807-4fc1-bfad-6b7516c89b13",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "c0c85c1d-9bb3-4e4d-a2d4-13fd458f15bc",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge8",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "4cc8b34f-9e3e-4f05-a73a-035395ff546d",
					Title:       "Organic Gardening",
					Difficulty:  "Easy",
					Description: "Start and maintain an organic garden",
					Exp:         80,
					Coin:        40,
					DateStart:   time.Date(2024, 12, 20, 10, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 12, 27, 10, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "2d9c8fd2-8104-470f-a8d3-4ff34993fabd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "19ecd259-58f0-4b3c-ba6b-e16f03317048",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge9",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "1e0f4c4f-5b54-4421-b8c1-46a2d382f8b2",
					Title:       "Upcycle Project",
					Difficulty:  "Medium",
					Description: "Complete an upcycling project with materials you already have",
					Exp:         110,
					Coin:        55,
					DateStart:   time.Date(2025, 01, 05, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 01, 12, 9, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "14eb3066-5f5c-459b-a84a-c99cbbd114f9",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "3ecbf571-79b6-452b-8626-c45e1a2ab671",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge10",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "c0aa381d-4c14-4eba-9811-c6356176ab67",
					Title:       "Eco-friendly Shopping",
					Difficulty:  "Easy",
					Description: "Buy only eco-friendly products for a week",
					Exp:         90,
					Coin:        45,
					DateStart:   time.Date(2025, 02, 20, 8, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 03, 01, 0, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "1350c308-05e7-4d49-a5ad-3e92b00ecba9",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e7f1ae18-bc65-4f29-a8c7-dd0c30ab3977",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge11",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "1f6774a8-a173-4161-ad3e-e58400e007b2",
					Title:       "Water Conservation Challenge",
					Difficulty:  "Hard",
					Description: "Reduce your water usage by 30% for a month",
					Exp:         280,
					Coin:        140,
					DateStart:   time.Date(2025, 03, 01, 8, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 03, 23, 59, 59, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "19ecd259-58f0-4b3c-ba6b-e16f03317048",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "a8089bc1-de1c-4eaf-a1fd-a2a6bc3028e2",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge12",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "b0ac9f1c-4e13-4848-977a-02dd540369a7",
					Title:       "Participate in Earth Hour",
					Difficulty:  "Easy",
					Description: "Turn off all lights and electrical appliances for one hour",
					Exp:         40,
					Coin:        20,
					DateStart:   time.Date(2025, 03, 29, 20, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 03, 29, 21, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "0c30941d-a7fc-438f-8ab7-5bba4f1cf57f",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "a8089bc1-de1c-4eaf-a1fd-a2a6bc3028e2",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge13",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "2e80a9cf-5007-4eac-9cc8-9a603d954d6a",
					Title:       "Meatless Monday",
					Difficulty:  "Medium",
					Description: "Go meatless every Monday for a month",
					Exp:         130,
					Coin:        65,
					DateStart:   time.Date(2025, 04, 01, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 04, 30, 59, 59, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "2d9c8fd2-8104-470f-a8d3-4ff34993fabd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "8559dcee-ffcf-4aa0-b77d-98dd82f4e96b",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge14",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "0d1233b8-d6e0-40df-8c6f-2ae1f8f2bfb7",
					Title:       "Create a Compost Bin",
					Difficulty:  "Easy",
					Description: "Create and maintain a compost bin at home",
					Exp:         70,
					Coin:        35,
					DateStart:   time.Date(2025, 05, 01, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 05, 05, 7, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "1cd52acc-f272-4aa8-8abd-e64fb288f64c",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "078817ea-4fcf-4eeb-96b2-f10961826469",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge15",
			Run: func(db *gorm.DB) error {
				challenges := challenge.Challenge{
					ID:          "ce6bf160-34ba-4685-80d7-10df1bd04fd1",
					Title:       "Participate in Car-Free Day",
					Difficulty:  "Medium",
					Description: "Avoid using cars for a whole day",
					Exp:         100,
					Coin:        50,
					DateStart:   time.Date(2025, 06, 05, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 06, 06, 7, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "3655727a-e807-4fc1-bfad-6b7516c89b13",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "c0c85c1d-9bb3-4e4d-a2d4-13fd458f15bc",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		// Challenge Ends
		// Users Starts
		{
			Name: "CreateUsers1",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "5dcdb121-85df-487c-8d4a-b4fd5033c9c0", "oscar.simanjuntak", "secret123", "Oscar Simanjuntak", "oscar.simanjuntak@example.com", "Jl. Kamboja No. 15, Jayapura", "Laki-laki", "081234567804", 350, 140, "")
			},
		},
		{
			Name: "CreateUsers2",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "5b9a23e5-e562-4fae-b11c-787e655a67d7", "ahmad.fauzi", "password123", "Ahmad Fauzi", "ahmad.fauzi@example.com", "Jl. Merdeka No. 1, Jakarta", "Laki-laki", "'081234567890", 200, 129, "")
			},
		},
		{
			Name: "CreateUsers3",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "d4a59d1b-38fa-4736-88d7-00cdc02ef5be", "budi.santoso", "passw0rd", "Budi Santoso", "budi.santoso@example.com", "Jl. Pahlawan No. 2, Surabaya", "Laki-laki", "081234567891", 50, 124, "")
			},
		},
		{
			Name: "CreateUsers4",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "6932e20c-3c55-4d95-91c5-ce316cc5843f", "citra.lestari", "mypassword", "Citra Lestari", "citra.lestari@example.com", "Jl. Melati No. 3, Bandung", "Perempuan", "081234567892", 0, 0, "")
			},
		},
		{
			Name: "CreateUsers5",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "8f25169f-5084-45c3-a99f-40daf59485d4", "dewi.ayu", "securepass", "Dewi Ayu", "dewi.ayu@example.com", "Jl. Mawar No. 4, Yogyakarta", "Perempuan", "'081234567893", 0, 24, "")
			},
		},
		{
			Name: "CreateUsers6",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "b7352ca3-fae1-4f3c-85b9-095de368b706", "eko.prasetyo", "password321", "Eko Prasetyo", "eko.prasetyo@example.com", "Jl. Kenanga No. 5, Semarang", "Laki-laki", "'081234567894", 200, 100, "")
			},
		},
		{
			Name: "CreateUsers7",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "29a1ee6d-05bd-4f11-b42a-a6769fd6a509", "fani.nugraha", "pass1234", "Fani Nugraha", "fani.nugraha@example.com", "Jl. Teratai No. 6, Medan", "Perempuan", "081234567895", 250, 125, "")
			},
		},
		{
			Name: "CreateUsers8",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "fbfccf91-f6ea-4b4f-952e-f371e9c69d9c", "gilang.permana", "qwerty123", "Gilang Permana", "gilang.permana@example.com", "Jl. Anggrek No. 7, Bali", "Laki-laki", "081234567896", 80, 64, "")
			},
		},
		{
			Name: "CreateUsers9",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "4fe8e112-e3d7-4468-bfd0-24058790bb6e", "hana.putri", "letmein", "Hana Putri", "hana.putri@example.com", "Jl. Dahlia No. 8, Makassar", "Perempuan", "081234567897", 110, 55, "")
			},
		},
		{
			Name: "CreateUsers10",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "56a85666-b5aa-4b39-8823-99c7f789ac7b", "indra.maulana", "welcome1", "Indra Maulana", "indra.maulana@example.com", "Jl. Flamboyan No. 9, Malang", "Laki-laki", "081234567898", 90, 45, "")
			},
		},
		{
			Name: "CreateUsers11",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "c0045c29-d1fa-4524-85fd-3e9341f99c6a", "surya.rahmat", "trustno1", "Surya Rahmat", "suryarahmat@example.com", "Jl. Cempaka No. 10, Solo", "Laki-laki", "081234567899", 0, 0, "")
			},
		},
		{
			Name: "CreateUsers12",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "6220c455-f14e-42d3-a73d-2f4573f4cf52", "kiki.amalia", "password456", "Kiki Amalia", "kiki.amalia@example.com", "Jl. Kamboja No. 11, Pekanbaru", "Perempuan", "081234567800", 40, 20, "")
			},
		},
		{
			Name: "CreateUsers13",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "35e243e9-d0b1-40b7-b216-5bc6cc15e243", "lina.marlina", "mypassword2", "Lina Marlina", "lina.marlina@example.com", "Jl. Bougenville No. 12, Palembang", "Tidak Terdefinisi", "081234567801", 0, 0, "")
			},
		},
		{
			Name: "CreateUsers14",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "8f14519b-ea61-457b-bf96-90504221de90", "mulyadi.saputra", "pass2022", "Mulyadi Saputra", "mulyadi.saputra@example.com", "Jl. Tulip No. 13, Balikpapan", "Tidak Terdefinisi", "081234567802", 0, 0, "")
			},
		},
		{
			Name: "CreateUsers15",
			Run: func(db *gorm.DB) error {
				return CreateUsers(db, "3fb1eb90-805f-4727-9f7d-a6fefc6115f9", "nina.agustina", "password789", "Nina Agustina", "nina.agustina@example.com", "Jl. Sakura No. 14, Pontianak", "Tidak Terdefinisi", "081234567803", 100, 50, "")
			},
		},
		// End of Users
		// Start of Vouchers
		{
			Name: "CreateVoucher1",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "84c65b94-163f-4880-badd-ad10ff8bfead", "Christmas Sale", "SUSTXMAS", "20000", "Special discounts on sustainable products", time.Date(2024, 12, 25, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher2",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "f1098036-cfa6-48ae-882c-00726d0c5d57", "Lebaran Bundle", "ECOLEBARAN", "15%", "Bundle deals on eco-friendly items for Lebaran", time.Date(2025, 5, 15, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher3",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "7e9c42d0-707c-4f88-b098-ae47888104fa", "Independence Day Sale", "RECYID", "15000", "Discounts on recyclable products for 17 Agustus", time.Date(2025, 8, 17, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher4",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "268a7851-53dd-4bc3-af91-0c84d00ea727", "New Year Offer", "RENEWNY", "30%", "Special offers on renewable resources for New Year", time.Date(2025, 1, 1, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher5",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "0e07708c-a23e-41a4-a999-58d4576743e4", "Earth Day Deals", "BIODEED", "10000", "Discounts on biodegradable goods for Earth Day", time.Date(2025, 4, 22, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher6",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "fb2643f7-796b-4a20-a3a7-bd009582215e", "Halloween Sale", "ENERHALLOW", "25%", "Savings on energy-efficient products for Halloween", time.Date(2024, 10, 31, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher7",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "756945dc-7906-47b4-95c5-650f428d7767", "Kartini Day Discount", "LOWCKARTINI", "30000", "Discounts on low carbon footprint items for Kartini Day", time.Date(2025, 4, 21, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher8",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "2e419999-cb91-4e36-a363-289976b7e83b", "Eid al-Adha Special", "WATEID", "20%", "Special offers for water-saving products for Eid al-Adha", time.Date(2025, 7, 10, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher9",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "880d707a-744b-41b9-991f-65c306074c03", "National Batik Day", "ZEROBATIK", "35%", "Big discounts for zero waste products for National Batik Day", time.Date(2025, 10, 2, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher10",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "4b7746fd-42ad-4090-8de9-b2fab6e1921c", "Lunar New Year Delights", "ORGLUNAR", "25%", "Discounts on organic and natural products for Lunar New Year", time.Date(2025, 2, 1, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher11",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "485d5b86-c1f9-45c1-a542-7b61c6feeecf", "Valentine's Day Sale", "NONTOXVAL", "12000", "Savings on non-toxic household items for Valentine's Day", time.Date(2025, 2, 14, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher12",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "bfbc3bd1-17fc-4766-b1ae-fa7535b665b3", "Earth Hour Offers", "COMPSEARTH", "15000", "Offers on compostable products for Earth Hour", time.Date(2025, 3, 29, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher13",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "c6613ef8-30a7-43ec-8b1f-78ae2491c7e5", "National Education Day Sale", "REUSEDU", "20%", "Discounts on reusable items for National Education Day", time.Date(2025, 5, 2, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher14",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "83c4db38-d196-49c8-9cb7-b53423e1658f", "National Health Day Special", "UPCYHEALTH", "30000", "Special deals on upcycled products for National Health Day", time.Date(2025, 11, 12, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher15",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "b1b68578-cba7-4b9d-8656-004737edc836", "Pancasila Day Savings", "FAIRTPAN", "15%", "Savings on fair trade certified products for Pancasila Day", time.Date(2025, 6, 1, 23, 59, 59, 0, time.UTC))
			},
		},
	}
	return seeds
}
