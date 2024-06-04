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
					Title:        "Plant a Tree",
					Difficulty:  "Medium",
					Description: "Plant a tree in your backyard or a community park",
					Exp:         100,
					Coin:        50,
					DateStart:   time.Date(2024, 6, 10, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 6, 17, 9, 0, 0, 0, time.UTC),
					ImageURL:     "https://storage.googleapis.com/alterra-greeve/greeve/170a0d9d-5c57-4d41-aa73-6e0ce1335345images.jpeg",
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
		// Challenge Ends
	}
	return seeds
}
