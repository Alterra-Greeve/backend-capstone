package seeds

import (
	challenge "backendgreeve/features/challenges/data"
	forum "backendgreeve/features/forums/data"
	product "backendgreeve/features/product/data"
	transaction "backendgreeve/features/transaction/data"
	user "backendgreeve/features/users/data"
	"backendgreeve/utils/database/seed"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seeds() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "CreateAdmin01",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "16930c07-bdb5-49d2-8a81-32591833241b", "admin", "admin", "admin@greeve.store", "admin")
			},
		},
		{
			Name: "CreateAdmin02",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "14adafd7-de6c-4586-a35e-3cf17ef3d351", "admin2", "admin2", "admin2@greeve.store", "admin2")
			},
		},
		{
			Name: "CreateAdmin3",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "b1e7fdc4-6465-4f06-8914-85b36255368e", "susantowijaya", "Susanto Wijaya", "susantowijaya@example.com", "password123")
			},
		},
		{
			Name: "CreateAdmin4",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "cf98c726-fb72-4687-9c88-e8ce671c32d6", "indrianisetiawati", "Indrianie Setiawati", "indrianisetiawati@example.com", "secret456")
			},
		},
		{
			Name: "CreateAdmin5",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "06db6d61-e7c8-4126-ab0f-f2b6bc1458d1", "ahmadrizkiputra", "Ahmad Rizki Putra", "ahmadrizkiputra@example.com", "qwerty789")
			},
		},
		{
			Name: "CreateAdmin6",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "6a53c2db-2fa5-47df-8462-baea4308d48f", "sucimutiara", "Suci Mutiara", "sucimutiara@example.com", "12345abcd")
			},
		},
		{
			Name: "CreateAdmin7",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "90aae29b-bf80-44bd-ab59-e7388e84e27f", "ariefulhaq", "Arie Ul Haq", "ariefulhaq@example.com", "mypassword")
			},
		},
		{
			Name: "CreateAdmin8",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "6112389a-2e44-4c78-ae64-849a5af60678", "dewianinditya", "Dewi Aninditya", "dewianinditya@example.com", "superstrong")
			},
		},
		{
			Name: "CreateAdmin9",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "9e66633a-cc3b-4887-8c44-dc8e4bf3530a", "rahmandarmawan", "Rahman Darmawan", "rahmandarmawan@example.com", "password!@#")
			},
		},
		{
			Name: "CreateAdmin10",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "b0835feb-2404-4a28-846a-52fafc8e6902", "natasyanurhaliza", "Natasya Nurhaliza", "natasyanurhaliza@example.com", "123456789")
			},
		},
		{
			Name: "CreateAdmin11",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "6562a858-84e7-412b-9e03-542f951fe0a2", "ilhamfauzan", "Ilham Fauzan", "ilhamfauzan@example.com", "password1")
			},
		},
		{
			Name: "CreateAdmin12",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "46a36705-f7d0-48a4-bb69-8cd5826b0da5", "tiaraputri", "Tiara Putri", "tiaraputri@example.com", "mypassword1")
			},
		},
		{
			Name: "CreateAdmin13",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "d1e93661-00b0-41a4-ab51-f2afaef8f86f", "gilangpratama", "Gilang Pratama", "gilangpratama@example.com", "secretpass")
			},
		},
		{
			Name: "CreateAdmin14",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "bb2c878b-2f05-4964-9307-a46d6bbccd31", "rianasafitri", "Riana Safitri", "rianasafitri@example.com", "strongpassword")
			},
		},
		{
			Name: "CreateAdmin15",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "7b8df5e7-c440-44f0-898a-d2714f762842", "yudhistirapratama", "Yudhistira Pratama", "yudhistirapratama@example.com", "password123!")
			},
		},
		{
			Name: "CreateAdmin16",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "98f9e9d2-9b19-4e14-b148-c1bb84104d07", "aisyahnuraini", "Aisyah Nuraini", "aisyahnuraini@example.com", "12345678")
			},
		},
		{
			Name: "CreateAdmin17",
			Run: func(db *gorm.DB) error {
				return CreateAdminLogin(db, "cea38442-5ede-409c-930f-d5f7dd772d5e", "rizkiadnan", "Rizki Adnan", "rizkiadnan@example.com", "mypassword2")
			},
		},
		// End Admin
		// Impact Categories
		{
			Name: "HematUang",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Save Money</strong> atau <strong>Hemat Uang</strong> adalah konsep yang mendorong individu untuk mengelola keuangan mereka dengan bijak. Ini melibatkan praktik-praktik seperti mengurangi pengeluaran yang tidak perlu, mencari cara untuk mendapatkan nilai lebih dari setiap pembelian, dan memaksimalkan efisiensi dalam penggunaan sumber daya. Dengan menghemat uang, individu dapat mencapai stabilitas finansial dan lebih mampu menghadapi situasi darurat atau merencanakan masa depan dengan lebih baik.</p><p>Contoh Dampak Positif Save Money pada Lingkungan:</p><ul><li>Mengurangi emisi gas rumah kaca dari produksi dan transportasi barang</li><li>Menghemat sumber daya alam seperti air, kayu, dan energi</li><li>Mengurangi polusi udara dan air dari limbah industri dan rumah tangga</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/5af169b6-94b5-45a7-aee9-13fee0fcc1a32f0376e551eec4f1af915d5983a621c1.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/hemat-uang.svg"
				return CreateImpactCategory(db, "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd", "Hemat Uang", 40, description, imageUrl, iconUrl)
			},
		},
		{
			Name: "MengurangiLimbah",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Less Waste</strong> atau <strong>Mengurangi Limbah</strong> adalah pendekatan yang menekankan pada pengurangan jumlah limbah yang dihasilkan, baik oleh individu, komunitas, maupun perusahaan. Tujuannya adalah untuk meminimalisir dampak negatif terhadap lingkungan dengan cara mengelola sumber daya secara lebih efisien dan bertanggung jawab. Strategi ini mencakup berbagai tindakan seperti daur ulang, pengomposan, penggunaan ulang barang, dan pengurangan penggunaan bahan sekali pakai.</p><p>Contoh Dampak Positif Less Waste pada Lingkungan:</p><ul><li>Pengurangan Polusi Tanah dan Air</li><li>Penurunan Emisi Gas Rumah Kaca</li><li>Konservasi Sumber Daya Alam</li><li>Mengurangi Energi yang Dibutuhkan untuk Produksi</li><li>Mengurangi Beban Tempat Pembuangan Akhir</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/91d948ca-3576-49e5-b435-0c3c7eccff7963ca36bac45556620e0ff4ba7ec5790d.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/mengurangi-limbah.svg"
				return CreateImpactCategory(db, "7d34a5fa-e2cf-466d-9f01-d731f6967082", "Mengurangi Limbah", 50, description, imageUrl, iconUrl)
			},
		},
		{
			Name: "PerluasWawasan",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Expand Your Mind</strong> atau <strong>Perluas Wawasan Anda</strong> adalah konsep yang mendorong individu untuk terbuka terhadap pengetahuan, ide, dan pengalaman baru. Ini melibatkan pembelajaran berkelanjutan, eksplorasi topik yang beragam, dan berpikir kritis. Dengan memperluas wawasan, kita tidak hanya meningkatkan pemahaman kita tentang dunia, tetapi juga mengembangkan empati, kreativitas, dan kemampuan untuk membuat keputusan yang lebih baik.</p><p>Contoh Dampak Positif Save Money pada Lingkungan:</p><ul><li>Kesadaran Lingkungan yang Lebih Tinggi</li><li>Inovasi dalam Solusi Ramah Lingkungan</li><li>Dukungan untuk Kebijakan Lingkungan yang Lebih Baik</li><li>Pengembangan Komunitas Berkelanjutan</li><li>Konsumsi yang Lebih Bertanggung Jawab</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/fd98dcdc-4cea-428a-a008-e1ca6f972ca72f0376e551eec4f1af915d5983a621c1.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/menambah-wawasan.svg"
				return CreateImpactCategory(db, "e8e714bd-c34e-4278-980c-39bd1f55b5fb", "Perluas Wawasan", 45, description, imageUrl, iconUrl)
			},
		},
		{
			Name: "MengurangiPemanasanGlobal",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Less Global Warming</strong> atau <strong>Mengurangi Pemanasan Global</strong> adalah upaya untuk menurunkan peningkatan suhu rata-rata bumi yang disebabkan oleh aktivitas manusia, terutama melalui emisi gas rumah kaca seperti karbon dioksida (CO<sub>2</sub>) dan metana (CH<sub>4</sub>). Upaya ini mencakup berbagai tindakan, seperti meningkatkan efisiensi energi, menggunakan energi terbarukan, mengurangi deforestasi, dan mempromosikan praktik pertanian berkelanjutan. Tujuannya adalah untuk memperlambat perubahan iklim, melindungi ekosistem, dan menjaga keseimbangan alami bumi.</p><p>Contoh Dampak Positif Save Money pada Lingkungan:</p><ul><li>Pengurangan Fenomena Cuaca Ekstrem</li><li>Perlindungan Keanekaragaman Hayati</li><li>Stabilisasi Tingkat Permukaan Laut</li><li>Kualitas Udara yang Lebih Baik</li><li>Konservasi Sumber Daya Air</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/38c805dd-9fbf-4fa0-bef7-9adfce99add9d5787157f3d099da70dcbf8f3021c99b.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/pemanasan-global-2.svg"
				return CreateImpactCategory(db, "b5d07366-3b31-4011-95e3-34735b0b61f8", "Mengurangi Pemanasan Global", 35, description, imageUrl, iconUrl)
			},
		},
		// Impact Categoreis End
		//Product Seeds Start
		{
			Name: "CreateProduct1",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7dcf0f20-d1c6-4d8a-8cf5-de60f5c420be",
					Name:        "Sedotan yang Dapat Digunakan Kembali",
					Description: "Satu set sedotan stainless steel dengan sikat pembersih.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7fbf0a60-57bf-419f-900a-a993d3168cec",
					Name:        "Tas Katun Organik",
					Description: "Tas ramah lingkungan yang terbuat dari 100% katun organik.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "7fbf0a60-57bf-419f-900a-a993d3168cec",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7c45085a-2466-4f5e-9ccd-bdd139751dac",
					Name:        "Sikat Gigi Biodegradable",
					Description: "Sikat gigi yang terbuat dari bambu dengan bulu sikat yang dapat terurai.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "b3b4ae7a-ed2e-4163-bf71-1a00b089c8a3",
					Name:        "Buku Catatan Ramah Lingkungan",
					Description: "Buku catatan yang terbuat dari kertas daur ulang dan tinta ramah lingkungan.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7dc84308-0330-41f1-8072-ca635e5055c8",
					Name:        "Botol Air yang Dapat Digunakan Kembali",
					Description: "Botol air tahan lama yang terbuat dari stainless steel.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "cc8d605c-db48-40ec-b62c-826a6aa5f9a2",
					Name:        "Pengisi Daya Tenaga Surya",
					Description: "Pengisi daya portabel bertenaga surya untuk perangkat Anda.",
					Price:       100000,
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
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "cc8d605c-db48-40ec-b62c-826a6aa5f9a2",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "809db141-0071-478e-bf49-9666968c82b5",
					Name:        "Set Alat Makan Bambu",
					Description: "Satu set alat makan yang dapat digunakan kembali yang terbuat dari bambu.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "5d502c20-9e83-4e84-9d93-db901a3508b5",
					Name:        "Kopi Perdagangan Adil",
					Description: "Biji kopi organik dan perdagangan adil.",
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
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "5d502c20-9e83-4e84-9d93-db901a3508b5",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "f5e9bae2-1e04-4e2f-8ed5-319849ba8b39",
					Name:        "Piring Kompos",
					Description: "Piring yang terbuat dari bahan kompos.",
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "08989b73-113e-4161-8047-9642db33fea3",
					Name:        "Lampu Hemat Energi",
					Description: "Bola lampu LED yang menghemat energi dan tahan lama.",
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
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "08989b73-113e-4161-8047-9642db33fea3",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "bc601a3e-513f-4ba5-b1ac-171d27489b90",
					Name:        "Laci Daur Ulang",
					Description: "Laci yang terbuat dari bahan daur ulang.",
					Price:       900000,
					Coin:        30,
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
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "e5da3310-6ec6-4b43-81c2-a6eb5a0538c8",
					Name:        "Sepatu dengan Jejak Karbon Rendah",
					Description: "Sepatu dengan jejak karbon rendah.",
					Price:       750000,
					Coin:        25,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "e5da3310-6ec6-4b43-81c2-a6eb5a0538c8",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/505e2f53-bffd-4d0b-9986-48c57aa51431seoarticle7desktop1-1645003305937.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "e5da3310-6ec6-4b43-81c2-a6eb5a0538c8",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/4b31515d-913e-4562-b17c-a398c7bf287aseoarticle7desktop2-1645003305957.webp",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "e5da3310-6ec6-4b43-81c2-a6eb5a0538c8",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "a10882e9-ef38-4796-bc25-f56ca2374fed",
					Name:        "Cat Non-Toksik",
					Description: "Cat yang terbuat dari bahan alami dan tidak beracun.",
					Price:       80000,
					Coin:        18,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "a10882e9-ef38-4796-bc25-f56ca2374fed",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b354682c-0d94-4b60-8a5b-81ec3e68d71fimages (3).jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "a10882e9-ef38-4796-bc25-f56ca2374fed",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "0e1c5843-a637-4a92-bfa0-7c7342f51846",
					Name:        "Kertas Daur Ulang",
					Description: "Kertas yang terbuat dari 100% bahan daur ulang.",
					Price:       12000,
					Coin:        5,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "0e1c5843-a637-4a92-bfa0-7c7342f51846",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/499b5dd0-1df5-4f50-a909-9cc297dd5dac5d1c5b4a544aa5b3ec360cc6db50b318.jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "0e1c5843-a637-4a92-bfa0-7c7342f51846",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "de45d430-da60-4701-a51a-ffda02ba69c5",
					Name:        "Tas Belanja yang Dapat Digunakan Kembali",
					Description: "Tas belanja yang kuat dan tahan lama yang terbuat dari bahan daur ulang.",
					Price:       24000,
					Coin:        7,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "de45d430-da60-4701-a51a-ffda02ba69c5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2603e2f9-26f8-44dc-8e8f-92f604f00dffTrashy_Smart_Bag_-_Unfolded.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "de45d430-da60-4701-a51a-ffda02ba69c5",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct16",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "9d1a9780-55e1-43c0-ba09-9e0b40e7cb81",
					Name:        "Perawatan Kulit Organik",
					Description: "Produk perawatan kulit yang terbuat dari bahan organik.",
					Price:       445000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "9d1a9780-55e1-43c0-ba09-9e0b40e7cb81",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2290ad20-df4c-4bfd-8bb7-94283bf7f920regenerating-cleanser-main.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "9d1a9780-55e1-43c0-ba09-9e0b40e7cb81",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct17",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "9376eb2d-b9b6-4674-b949-8e3bf42616f5",
					Name:        "Cangkir Kopi yang Dapat Digunakan Kembali",
					Description: "Cangkir kopi yang dapat digunakan kembali, terbuat dari bahan ramah lingkungan.",
					Price:       20000,
					Coin:        12,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "9376eb2d-b9b6-4674-b949-8e3bf42616f5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2d11719e-b9b6-434f-b4a1-783b79ffd799no-brand_no-brand_full01.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "9376eb2d-b9b6-4674-b949-8e3bf42616f5",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct18",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "f0f542fb-4810-458a-b679-8f4ab4e08493",
					Name:        "Batang Cokelat Perdagangan Adil",
					Description: "Batang cokelat yang terbuat dari kakao perdagangan adil.",
					Price:       25000,
					Coin:        15,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "f0f542fb-4810-458a-b679-8f4ab4e08493",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/561cd4c0-3fbb-4578-87a0-b127b5a40220images (4).jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "f0f542fb-4810-458a-b679-8f4ab4e08493",
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "f0f542fb-4810-458a-b679-8f4ab4e08493",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct19",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "fae8d2c5-8374-41e3-b2dc-f38ca67f1723",
					Name:        "Sarung Ponsel Biodegradable",
					Description: "Sarung ponsel yang sepenuhnya dapat terurai.",
					Price:       65000,
					Coin:        9,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "fae8d2c5-8374-41e3-b2dc-f38ca67f1723",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/7101ef8d-d9c1-4ed3-a8ea-f46f5cec3d08images (5).jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "fae8d2c5-8374-41e3-b2dc-f38ca67f1723",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct20",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "23cf052f-26ce-4f99-b61c-6b1cba4e630e",
					Name:        "Matras Yoga Ramah Lingkungan",
					Description: "Matras yoga yang terbuat dari bahan ramah lingkungan dan tidak beracun.",
					Price:       72000,
					Coin:        14,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "23cf052f-26ce-4f99-b61c-6b1cba4e630e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8f317cad-8d3e-4341-a253-6d60a8334cb0REECH.jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "23cf052f-26ce-4f99-b61c-6b1cba4e630e",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct21",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "f5dac7fe-afb6-4eff-9203-68f6994be2db",
					Name:        "Pembungkus Makanan yang Dapat Digunakan Kembali",
					Description: "Pembungkus yang terbuat dari lilin lebah dan katun organik.",
					Price:       15000,
					Coin:        6,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "f5dac7fe-afb6-4eff-9203-68f6994be2db",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b95275d1-d50c-4365-a804-a0d86168bd3fHoneycombRoll1_1200x.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "f5dac7fe-afb6-4eff-9203-68f6994be2db",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/a9e852d0-b485-4675-b45f-80a7078cf2b6BW.comPDPLifestyleImages_1000x1000px_19_1200x.webp",
						Position:  2,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "f5dac7fe-afb6-4eff-9203-68f6994be2db",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/a4a917b4-20e4-4150-a7a0-8437ef94c9853_29941402-9cba-4802-9f45-edbaaf64e64b_1200x.webp",
						Position:  3,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "f5dac7fe-afb6-4eff-9203-68f6994be2db",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct22",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "644c996a-5407-416f-86a3-b8fa9233886f",
					Name:        "Lampu Taman Tenaga Surya",
					Description: "Lampu taman yang bertenaga surya.",
					Price:       230000,
					Coin:        20,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "644c996a-5407-416f-86a3-b8fa9233886f",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/d063a9c6-3ade-4df4-b05a-c0efab059a7761z7WoMtSnL.__AC_SY445_SX342_QL70_FMwebp_.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "644c996a-5407-416f-86a3-b8fa9233886f",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b2aa4228-81ef-4d47-bd58-73c9bac47fe871hxa7Mtq+L._AC_SX679_.jpg",
						Position:  2,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "644c996a-5407-416f-86a3-b8fa9233886f",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/1833a1e6-9302-4c92-837c-f1b7e90799ceimages (9).jpeg",
						Position:  3,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "644c996a-5407-416f-86a3-b8fa9233886f",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "644c996a-5407-416f-86a3-b8fa9233886f",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct23",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "30b3c9d3-3304-4da4-8cda-01052e385eb5",
					Name:        "Alat Makan Kompos",
					Description: "Alat makan sekali pakai yang dapat terkompos.",
					Price:       20000,
					Coin:        8,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "30b3c9d3-3304-4da4-8cda-01052e385eb5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/24e88d99-d3ec-4f1e-a08a-433ec28542d741J7+D28yIL._SX342_SY445_.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "30b3c9d3-3304-4da4-8cda-01052e385eb5",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct24",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7ec87ad7-5380-438e-996d-dea7538d3b64",
					Name:        "Dompet Daur Ulang",
					Description: "Dompet yang terbuat dari bahan daur ulang",
					Price:       115000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7ec87ad7-5380-438e-996d-dea7538d3b64",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/c5e5cfd7-accc-4a33-a1e4-2173c1b1a3fcil_1588xN.1925528247_3zpq.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7ec87ad7-5380-438e-996d-dea7538d3b64",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct25",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "02/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "4f164404-0792-470c-89ee-3e48fd53e50c",
					Name:        "Ransel dengan Jejak Karbon Rendah",
					Description: "Ransel dengan jejak karbon rendah,",
					Price:       800000,
					Coin:        25,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "4f164404-0792-470c-89ee-3e48fd53e50c",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/db2ab4ba-4865-4a11-bd99-86c15c057da0814-173-38113.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "4f164404-0792-470c-89ee-3e48fd53e50c",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/732499c3-ccba-49b0-8f64-0220ce13e8f6814-173-38113_1.jpg",
						Position:  2,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "4f164404-0792-470c-89ee-3e48fd53e50c",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8f3f23bd-55f9-46f3-b46d-8ebeed584f6a814-173-38113_2.webp",
						Position:  3,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "4f164404-0792-470c-89ee-3e48fd53e50c",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct26",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "0e65c6d4-15ab-435f-8b63-364df43fd70f",
					Name:        "Tisu Pembersih Biodegradable",
					Description: "Tisu pembersih yang terbuat dari bahan yang dapat terurai.",
					Price:       18000,
					Coin:        5,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "0e65c6d4-15ab-435f-8b63-364df43fd70f",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/5fe5c383-3e39-41dc-9b40-eb0fabcbfe9agreenworks-wipes.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "0e65c6d4-15ab-435f-8b63-364df43fd70f",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct27",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "4d1d4263-491f-4eb9-a9b3-e61c5ba812e5",
					Name:        "Tas Sandwich yang Dapat Digunakan Kembali",
					Description: "Tas sandwich yang terbuat dari bahan yang tahan lama dan bisa dicuci.",
					Price:       12000,
					Coin:        6,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "4d1d4263-491f-4eb9-a9b3-e61c5ba812e5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/0da34de6-7fce-4df0-9d49-6fe9b9d48c8c81evah0JzDL.__AC_SX300_SY300_QL70_FMwebp_.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "4d1d4263-491f-4eb9-a9b3-e61c5ba812e5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/d07e981e-fc49-4a67-9e9f-cce6ca76512371oFThg1eOL._AC_SX679_.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "4d1d4263-491f-4eb9-a9b3-e61c5ba812e5",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct28",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "2dd4ca9b-d2e8-40c9-8a56-964c9db60106",
					Name:        "Handuk Katun Organik",
					Description: "Handuk lembut yang terbuat dari katun organik.",
					Price:       25000,
					Coin:        8,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "2dd4ca9b-d2e8-40c9-8a56-964c9db60106",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/58a83c3d-2b7d-45e1-b188-674e9713f1feORIB-BLU_fddceb2b-aa79-4f09-bad8-c4fa3d7dffc6_768x.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "2dd4ca9b-d2e8-40c9-8a56-964c9db60106",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct29",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "2c0819d5-5c18-4954-b99a-f796e655f8f2",
					Name:        "Deterjen Cucian Ramah Lingkungan",
					Description: "Deterjen cucian yang terbuat dari bahan alami.",
					Price:       30000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "2c0819d5-5c18-4954-b99a-f796e655f8f2",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/30aa05ec-6d97-4807-a1cb-a3ae04fc313cbest-eco-detergents-good-housekeeping-1624544533.png",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "2c0819d5-5c18-4954-b99a-f796e655f8f2",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct30",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "4d844f6d-a1e1-44d4-8627-29ee6d06ec44",
					Name:        "Vas Kaca Daur Ulang",
					Description: "Vas yang terbuat dari kaca daur ulang.",
					Price:       20000,
					Coin:        12,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "4d844f6d-a1e1-44d4-8627-29ee6d06ec44",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/03be69a9-6bfb-4b3b-aaeb-19c8262b53dfimages (6).jpeg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "4d844f6d-a1e1-44d4-8627-29ee6d06ec44",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct31",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "e3f27cb4-dcda-4114-b1aa-d1afd1e65ab3",
					Name:        "Mainan Anak Non-Toksik",
					Description: "Mainan yang terbuat dari bahan yang tidak beracun dan aman.",
					Price:       45000,
					Coin:        15,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "e3f27cb4-dcda-4114-b1aa-d1afd1e65ab3",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/e6945979-9a10-4abc-b64c-bf6a8fa10f663416_-_Packshot_-_01.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "e3f27cb4-dcda-4114-b1aa-d1afd1e65ab3",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/1c71b970-f061-4a45-bfc8-e146a04586683416_-_Packshot_-_02.webp",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "e3f27cb4-dcda-4114-b1aa-d1afd1e65ab3",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct32",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "dfddfce8-553a-45da-b62a-80d3bc6f96ac",
					Name:        "Teh Herbal Organik",
					Description: "Teh herbal yang terbuat dari bahan organik dan perdagangan adil.",
					Price:       22000,
					Coin:        8,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "dfddfce8-553a-45da-b62a-80d3bc6f96ac",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b1b9cea6-a5da-47a9-9bad-09dbe3e2aa00rs=w_600,h_600.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "dfddfce8-553a-45da-b62a-80d3bc6f96ac",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct33",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "03/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7be7821a-c433-4597-bf94-a4e464cbc37e",
					Name:        "Bantalan Penghapus Riasan yang Dapat Digunakan Kembali",
					Description: "Bantalan yang terbuat dari katun organik, dapat digunakan kembali dan bisa dicuci.",
					Price:       15000,
					Coin:        6,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7be7821a-c433-4597-bf94-a4e464cbc37e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b3acfa7d-10bb-43cd-a2df-422b6db1af1441ecjg7T0pL._SY445_SX342_QL70_FMwebp_.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "7be7821a-c433-4597-bf94-a4e464cbc37e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/82520c0e-2159-4bca-8f87-4c4558d6471e91h0aTKT2-L._SX466_.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7be7821a-c433-4597-bf94-a4e464cbc37e",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct34",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "4895d8dd-0231-44dc-a3a5-a7ab395d1094",
					Name:        "Masker Wajah Katun Organik",
					Description: "Masker wajah yang terbuat dari 100% katun organik, lembut, bisa bernapas, dan nyaman untuk penggunaan sehari-hari. Masker ini dirancang untuk dapat digunakan kembali dan dicuci, mengurangi kebutuhan akan masker sekali pakai.",
					Price:       30000,
					Coin:        20,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "4895d8dd-0231-44dc-a3a5-a7ab395d1094",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/38616064-57d5-46a5-b744-9e18cf851413face_mask_organic_cotton_13.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "4895d8dd-0231-44dc-a3a5-a7ab395d1094",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/93e2f710-6e28-4ad9-b5dc-a6b2fa77bc22face_mask_organic_cotton_9.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "4895d8dd-0231-44dc-a3a5-a7ab395d1094",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct35",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "005c310b-015a-4075-b3b5-575b1feb14d6",
					Name:        "Termostat Pintar",
					Description: "Termostat yang dapat diprogram yang menghemat energi dan uang.",
					Price:       1500000,
					Coin:        50,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "005c310b-015a-4075-b3b5-575b1feb14d6",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/7395c073-f751-4017-8662-3da92efe4511sZWtCnFaQH.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "005c310b-015a-4075-b3b5-575b1feb14d6",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "005c310b-015a-4075-b3b5-575b1feb14d6",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct36",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "ba8150f4-103a-438a-8d14-90ad8dc46a30",
					Name:        "Mesin Cuci Hemat Air",
					Description: "Mesin cuci yang hemat energi dan menghemat air.",
					Price:       3800000,
					Coin:        70,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "ba8150f4-103a-438a-8d14-90ad8dc46a30",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/d95092b9-6bf2-44e9-bf2e-81fa18a648ebasko-w2084cw_1.jpeg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "ba8150f4-103a-438a-8d14-90ad8dc46a30",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/284813e4-e6e5-4546-97a8-5a92a1b89530asko-w2084cw_3.jpeg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "ba8150f4-103a-438a-8d14-90ad8dc46a30",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "ba8150f4-103a-438a-8d14-90ad8dc46a30",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct37",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "9d73a948-298b-4896-8d43-2ad828cbd836",
					Name:        "Lengan Laptop Denim Daur Ulang",
					Description: "Perlindungan laptop yang unik dan ramah lingkungan yang terbuat dari denim daur ulang.",
					Price:       250000,
					Coin:        40,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "9d73a948-298b-4896-8d43-2ad828cbd836",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/c37ba60a-952c-45b7-8d14-79a25a2a3c48Rimagined_Laptop_Sleeves_Checkered_Src_4_1800x1800.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "9d73a948-298b-4896-8d43-2ad828cbd836",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/3b3a6921-b046-4e77-9a19-4d1c9f29cc82Rimagined_Laptop_Sleeves_Checkered_Src_8.webp",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "9d73a948-298b-4896-8d43-2ad828cbd836",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct38",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
					Name:        "Kaos Katun Organik Perdagangan Adil",
					Description: "Kaos katun organik yang bersumber secara etis dan berkelanjutan.",
					Price:       400000,
					Coin:        60,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/271fe697-b266-404b-8f39-169819a59bc1front-FS01-WH__37153.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/0331be78-a4be-4d67-bb46-2cc18a9acd51fairtrade_white_t_2__30295.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct39",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "591c7cbe-3820-4ce1-991c-3294169f815b",
					Name:        "Sabun Berbasis Tumbuhan",
					Description: "Sabun yang dibuat dari bahan alami, bebas dari bahan kimia sintetis dan pewangi buatan. Lembut di kulit, cocok untuk semua jenis kulit, termasuk kulit sensitif.",
					Price:       180000,
					Coin:        35,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "591c7cbe-3820-4ce1-991c-3294169f815b",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/c41ebb0d-2ebc-49a6-b9ce-44af452dd01363aec880-1385-4edd-950e-763f339eb6ad.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "591c7cbe-3820-4ce1-991c-3294169f815b",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2aa2e341-7c9d-4e71-9ea1-ad8f3b0efae9a5c9b2a6-926e-44c3-8be1-7d2964081a37.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "591c7cbe-3820-4ce1-991c-3294169f815b",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "591c7cbe-3820-4ce1-991c-3294169f815b",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct40",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "aad3789d-706f-417a-ad86-2d8762d9ce3a",
					Name:        "Lilin Kedelai",
					Description: "Lilin yang bersih terbakar dan berkelanjutan yang terbuat dari lilin kedelai.",
					Price:       150000,
					Coin:        28,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "aad3789d-706f-417a-ad86-2d8762d9ce3a",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/aea1114f-c3f9-4ed3-b0a7-1c635cfa875310ae8479-c692-4960-8b8e-8335548f093d.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "aad3789d-706f-417a-ad86-2d8762d9ce3a",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct41",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "bddb97ac-1212-440a-ae73-00c89a1621b2",
					Name:        "Wadah Persiapan Makanan Kaca",
					Description: "Wadah penyimpanan makanan yang tahan lama dan dapat digunakan kembali.",
					Price:       300000,
					Coin:        45,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "bddb97ac-1212-440a-ae73-00c89a1621b2",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/fc50a02b-ace4-4b6e-872b-5a12543520f671oWCoV+sFL._AC_SX679_.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "bddb97ac-1212-440a-ae73-00c89a1621b2",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/a9e14f84-6978-44ab-82ca-2b2193871b5f8150MJRwgjL._AC_SY879_.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "bddb97ac-1212-440a-ae73-00c89a1621b2",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "bddb97ac-1212-440a-ae73-00c89a1621b2",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct42",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7a371522-d22e-431b-bf25-149e80809c27",
					Name:        "Organizer Meja Bambu",
					Description: "Solusi organisasi meja yang berkelanjutan dan serbaguna.",
					Price:       350000,
					Coin:        32,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7a371522-d22e-431b-bf25-149e80809c27",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/19a98ddb-7c04-4cd2-bfcc-70d0d62ab81e633f74d7ed5b7f6c86198e36-bamboo-desktop-organizer-home-office.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7a371522-d22e-431b-bf25-149e80809c27",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct43",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "627efc9c-a3bf-4850-b4fa-0eac1768d042",
					Name:        "Cat Ramah Lingkungan",
					Description: "Cat tidak beracun, rendah VOC yang terbuat dari bahan berkelanjutan.",
					Price:       500000,
					Coin:        50,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "627efc9c-a3bf-4850-b4fa-0eac1768d042",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/93a90f0a-8e21-4fe7-8f7d-e4f23bb9e7bdeggshell_paint_wb-800x800.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "627efc9c-a3bf-4850-b4fa-0eac1768d042",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct44",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "719d7f9e-7ae5-4e8a-97fe-eda6dcf8a756",
					Name:        "Kotak Makan Bambu",
					Description: "Kotak makan yang tahan lama dan dapat terurai yang terbuat dari bambu yang dapat diperbaharui.",
					Price:       220000,
					Coin:        35,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "719d7f9e-7ae5-4e8a-97fe-eda6dcf8a756",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/86ab5c9f-64ff-4768-b9de-d6e402a4783c71ohIptAImL._AC_UF894,1000_QL80_.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "719d7f9e-7ae5-4e8a-97fe-eda6dcf8a756",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct45",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "8b99ccfa-8517-45bc-b7c5-fd4fd5a56b29",
					Name:        "Sapu Plastik Daur Ulang",
					Description: "Sapu kokoh yang terbuat dari 100% plastik daur ulang untuk pembersihan nol limbah.",
					Price:       170000,
					Coin:        25,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "8b99ccfa-8517-45bc-b7c5-fd4fd5a56b29",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/91856227-4892-408d-8b48-70fcd60f30e6EZAHM37.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "8b99ccfa-8517-45bc-b7c5-fd4fd5a56b29",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct46",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "868790b7-4439-49c4-b960-d48bde0fed0b",
					Name:        "Lampu Teras Tenaga Surya",
					Description: "Manfaatkan energi terbarukan untuk menerangi ruang luar Anda.",
					Price:       550000,
					Coin:        60,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "868790b7-4439-49c4-b960-d48bde0fed0b",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/680ffddc-0275-469c-b2f8-0f92842dd0fe71VOmyQRn0L.__AC_SY445_SX342_QL70_FMwebp_.webp",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "868790b7-4439-49c4-b960-d48bde0fed0b",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/a2e14472-f661-4d22-86e4-67e97120dfc161xuY0H-eIL._AC_SX679_.jpg",
						Position:  2,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "868790b7-4439-49c4-b960-d48bde0fed0b",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/60f95086-7144-4b12-8798-5055acfefd9b71yejp4eNPL._AC_SX679_.jpg",
						Position:  3,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "868790b7-4439-49c4-b960-d48bde0fed0b",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "868790b7-4439-49c4-b960-d48bde0fed0b",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct47",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "0d4dc83a-6a0c-4d38-8ae0-8a0eba8c03dc",
					Name:        "Kaus Kaki Katun Organik",
					Description: "Kaus kaki yang terbuat dari 100% katun organik bersertifikat GOTS,  mengurangi frekuensi penggantian.",
					Price:       40000,
					Coin:        30,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "0d4dc83a-6a0c-4d38-8ae0-8a0eba8c03dc",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/b422f557-fb5f-4bbc-a2c5-649f01642310136311-a18e6b57f5fd4fa488aaaa859c1d891a_900x.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "0d4dc83a-6a0c-4d38-8ae0-8a0eba8c03dc",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct48",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "66bcdf40-2544-4bec-b9be-0283a030e709",
					Name:        "Sweter Katun Organik",
					Description: "Pakaian lembut dan berkelanjutan yang terbuat dari katun organik bersertifikat GOTS.",
					Price:       650000,
					Coin:        55,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "66bcdf40-2544-4bec-b9be-0283a030e709",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/eb0e7a1f-f54f-48a8-8fd3-cc94b629bcf6dg_organic_basics-men-organic_cotton-merch-sweatshirt-grey_melange-packshot-1.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "66bcdf40-2544-4bec-b9be-0283a030e709",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct49",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "b677eb44-63be-4ffb-8741-8d155b912d21",
					Name:        "Kepala Shower Hemat Air",
					Description: "Kurangi konsumsi air dengan kepala shower yang efisien ini.",
					Price:       400000,
					Coin:        45,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "b677eb44-63be-4ffb-8741-8d155b912d21",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/d559f50a-68c0-460a-be9f-7eb909e735e4ZZH_T27573A01_000_02.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "b677eb44-63be-4ffb-8741-8d155b912d21",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/ce1583ab-fbae-4afa-839c-d44732810d7fZZH_T27929A01_000_03.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "b677eb44-63be-4ffb-8741-8d155b912d21",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "b677eb44-63be-4ffb-8741-8d155b912d21",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				products.Images = images
				products.ImpactCategories = categories
				return CreateProduct(db, products)
			},
		},
		{
			Name: "CreateProduct50",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/01/2024 10:00:00")
				if err != nil {
					return err
				}

				products := product.Product{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "7beeb22b-1928-4f72-8e9e-55525d792d89",
					Name:        "Pot Bunga dari Ban Daur Ulang",
					Description: "Beri kehidupan baru pada ban bekas dengan pot bunga unik ini.",
					Price:       200000,
					Coin:        30,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7beeb22b-1928-4f72-8e9e-55525d792d89",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/e0ec28d3-9ed0-49e3-a898-ef497eca9c06Classic-Pot-L-Front.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7beeb22b-1928-4f72-8e9e-55525d792d89",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "7beeb22b-1928-4f72-8e9e-55525d792d89",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "5fe8f4f6-43ee-4420-8969-bf1a0e19260c",
					Title:       "Menanam Pohon",
					Difficulty:  "Sedang",
					Description: "Tanam pohon di halaman belakang atau taman komunitas",
					Exp:         100,
					Coin:        50,
					DateStart:   time.Date(2024, 6, 10, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 6, 17, 9, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/12f1af01-9c5c-4223-a39f-640d4d5c46c2images (7).jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge2",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "ba2621ac-6f99-43cb-ab8a-7db626b7c4e4",
					Title:       "Minggu Tanpa Limbah",
					Difficulty:  "Sulit",
					Description: "Tidak menghasilkan limbah selama seminggu dan daur ulang segala sesuatu yang mungkin",
					Exp:         200,
					Coin:        100,
					DateStart:   time.Date(2024, 7, 1, 8, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 7, 8, 8, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/70126aef-8d2c-4776-8556-de3e14b452b8Zero-Waste-Week-1.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge3",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "5086c59a-e2f9-44d0-bcc6-f7fdb58f07a4",
					Title:       "Kerja Bakti Membersihkan Lingkungan",
					Difficulty:  "Mudah",
					Description: "Ikut serta dalam acara kerja bakti membersihkan lingkungan setempat",
					Exp:         50,
					Coin:        25,
					DateStart:   time.Date(2024, 8, 15, 10, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 8, 22, 10, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/3afd2556-65de-47eb-b461-e84beb9cfe63images (8).jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge4",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "aa013dc7-4fb3-4cf1-9cc0-e1a0ad60414d",
					Title:       "Bersepeda ke Tempat Kerja",
					Difficulty:  "Sedang",
					Description: "Berangkat kerja dengan sepeda selama seminggu",
					Exp:         150,
					Coin:        75,
					DateStart:   time.Date(2024, 9, 5, 7, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 9, 12, 7, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/411ed466-9abf-465a-b351-d84b9d746812images (9).jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge5",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "5550e03b-15c6-43c9-997c-4086adc8b573",
					Title:       "Bulan Hemat Energi",
					Difficulty:  "Sulit",
					Description: "Kurangi konsumsi energi rumah tangga Anda sebesar 20% selama sebulan",
					Exp:         300,
					Coin:        150,
					DateStart:   time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 10, 31, 23, 59, 59, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/21dc1d7c-34e0-4057-9cfc-0e2ac83189cbimages.png",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge6",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "9b21b3e5-46ff-42a1-92cc-4969c078ff83",
					Title:       "Membuat Produk Ramah Lingkungan",
					Difficulty:  "Sedang",
					Description: "Buat sendiri dan gunakan tiga produk rumah tangga ramah lingkungan",
					Exp:         120,
					Coin:        60,
					DateStart:   time.Date(2024, 11, 10, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 11, 17, 9, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/81d13e21-9952-4afb-a68c-6c968d58250emaxresdefault.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge7",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "85dd4189-9d2a-42f9-8b26-b03daa7e5869",
					Title:       "Gunakan Sistem Transportasi Hijau",
					Difficulty:  "Sulit",
					Description: "Gunakan hanya transportasi umum atau jalan kaki untuk semua perjalanan selama dua minggu",
					Exp:         250,
					Coin:        125,
					DateStart:   time.Date(2024, 12, 01, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 12, 15, 8, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/32f31bf9-96ad-41f9-98a7-ec207425e57aunnamed.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge8",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "4cc8b34f-9e3e-4f05-a73a-035395ff546d",
					Title:       "Berkebun Organik",
					Difficulty:  "Mudah",
					Description: "Mulai dan pelihara kebun organik",
					Exp:         80,
					Coin:        40,
					DateStart:   time.Date(2024, 12, 20, 10, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2024, 12, 27, 10, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/3b796884-3c8b-4817-9e32-c31c63963495images (10).jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge9",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "1e0f4c4f-5b54-4421-b8c1-46a2d382f8b2",
					Title:       "Proyek Daur Ulang",
					Difficulty:  "Sedang",
					Description: "Selesaikan proyek daur ulang dengan bahan yang sudah Anda miliki",
					Exp:         110,
					Coin:        55,
					DateStart:   time.Date(2025, 01, 05, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 01, 12, 9, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/4ee25df4-a2c6-42c1-9b33-a9487d775ad3diy-upcycled-trash-ideas-featured-homebnc-v2.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge10",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "c0aa381d-4c14-4eba-9811-c6356176ab67",
					Title:       "Belanja Barang Ramah Lingkungan",
					Difficulty:  "Mudah",
					Description: "Beli hanya produk ramah lingkungan selama seminggu",
					Exp:         90,
					Coin:        45,
					DateStart:   time.Date(2025, 02, 20, 8, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 03, 01, 0, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/dc808bf3-5f9e-4009-9d87-20a6adaca058less-plastic-1195643968.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge11",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "1f6774a8-a173-4161-ad3e-e58400e007b2",
					Title:       "Penghematan Air",
					Difficulty:  "Sulit",
					Description: "Kurangi penggunaan air Anda sebesar 30% selama sebulan",
					Exp:         280,
					Coin:        140,
					DateStart:   time.Date(2025, 03, 01, 8, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 03, 23, 59, 59, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/64198baa-1814-4c1d-8150-156cfef53226images (11).jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge12",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "b0ac9f1c-4e13-4848-977a-02dd540369a7",
					Title:       "Ikut Serta dalam Earth Hour",
					Difficulty:  "Mudah",
					Description: "Matikan semua lampu dan peralatan listrik selama satu jam",
					Exp:         40,
					Coin:        20,
					DateStart:   time.Date(2025, 03, 29, 20, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 03, 29, 21, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/4d8d8751-1ac4-4d28-99b6-9ed9c0f8f766earth-hour-campaign-poster-or-banner-turn-off-your-lights-for-our-planet-60-minutes-free-vector.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge13",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "2e80a9cf-5007-4eac-9cc8-9a603d954d6a",
					Title:       "Senin Tanpa Daging",
					Difficulty:  "Sedang",
					Description: "Tidak makan daging setiap hari Senin selama sebulan",
					Exp:         130,
					Coin:        65,
					DateStart:   time.Date(2025, 04, 01, 0, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 04, 30, 59, 59, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/2641676b-7411-44bb-a9c9-01fd78bc3b61images%20(1).png",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge14",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "0d1233b8-d6e0-40df-8c6f-2ae1f8f2bfb7",
					Title:       "Membuat Kompos Rumahan Sederhana",
					Difficulty:  "Mudah",
					Description: "Buat dan pelihara tempat kompos untuk sampah rumah tangga",
					Exp:         70,
					Coin:        35,
					DateStart:   time.Date(2025, 05, 01, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 05, 05, 7, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/92eb26ef-d764-43ba-9f89-d6766b268bfcimages%20(12).jpeg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
					},
				}
				challenges.ChallengeImpactCategories = challengeCategories
				return CreateChallenge(db, challenges)
			},
		},
		{
			Name: "CreateChallenge15",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "01/01/2024 8:00:00")
				if err != nil {
					return err
				}
				challenges := challenge.Challenge{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:          "ce6bf160-34ba-4685-80d7-10df1bd04fd1",
					Title:       "Ikut Serta dalam Hari Bebas Mobil",
					Difficulty:  "Sedang",
					Description: "Hindari menggunakan mobil dan kendaraan bermotor selama sehari",
					Exp:         100,
					Coin:        50,
					DateStart:   time.Date(2025, 06, 05, 9, 0, 0, 0, time.UTC),
					DateEnd:     time.Date(2025, 06, 06, 7, 0, 0, 0, time.UTC),
					ImageURL:    "https://storage.googleapis.com/alterra-greeve/greeve/399b5b8a-98a6-47f0-acb6-d314882d8b625eeef3334ed2b.jpg",
				}
				challengeCategories := []challenge.ChallengeImpactCategory{
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "e8e714bd-c34e-4278-980c-39bd1f55b5fb",
					},
					{
						ID:               uuid.New().String(),
						ChallengeID:      challenges.ID,
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				createdAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "05/01/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:         "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					Name:       "Oscar Simanjuntak",
					Password:   "secret123",
					Email:      "oscar.simanjuntak@example.com",
					Username:   "oscar.simanjuntak",
					Address:    "Jl. Kamboja No. 15, Jayapura",
					Gender:     "Laki-laki",
					Phone:      "081234567804",
					Coin:       350,
					Exp:        140,
					Membership: true,
					AvatarURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers2",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "11/01/2024 19:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "11/01/2024 19:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:         "5b9a23e5-e562-4fae-b11c-787e655a67d7",
					Name:       "Ahmad Fauzi",
					Password:   "password123",
					Email:      "ahmad.fauzi@example.com",
					Username:   "ahmad.fauzi",
					Address:    "Jl. Merdeka No. 1, Jakarta",
					Gender:     "Laki-laki",
					Phone:      "081234567890",
					Coin:       200,
					Exp:        129,
					Membership: true,
					AvatarURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers3",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "13/01/2024 10:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "13/01/2024 10:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:         "d4a59d1b-38fa-4736-88d7-00cdc02ef5be",
					Name:       "Budi Santoso",
					Password:   "passw0rd",
					Email:      "budi.santoso@example.com",
					Username:   "budi.santoso",
					Address:    "Jl. Pahlawan No. 2, Surabaya",
					Gender:     "Laki-laki",
					Phone:      "081234567891",
					Coin:       50,
					Exp:        124,
					Membership: true,
					AvatarURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers4",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "14/01/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "14/01/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "6932e20c-3c55-4d95-91c5-ce316cc5843f",
					Name:      "Citra Lestari",
					Password:  "mypassword",
					Email:     "citra.lestari@example.com",
					Username:  "citra.lestari",
					Address:   "Jl. Melati No. 3, Bandung",
					Gender:    "Perempuan",
					Phone:     "081234567892",
					Coin:      0,
					Exp:       0,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers5",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "19/01/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "19/01/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:         "8f25169f-5084-45c3-a99f-40daf59485d4",
					Name:       "Dewi Ayu",
					Password:   "securepass",
					Email:      "dewi.ayu@example.com",
					Username:   "dewi.ayu",
					Address:    "Jl. Mawar No. 4, Yogyakarta",
					Gender:     "Perempuan",
					Phone:      "081234567893",
					Coin:       0,
					Exp:        24,
					Membership: true,
					AvatarURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers6",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "21/01/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "21/01/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:         "b7352ca3-fae1-4f3c-85b9-095de368b706",
					Name:       "Eko Prasetyo",
					Password:   "password321",
					Email:      "eko.prasetyo@example.com",
					Username:   "eko.prasetyo",
					Address:    "Jl. Kenanga No. 5, Semarang",
					Gender:     "Laki-laki",
					Phone:      "081234567894",
					Coin:       200,
					Exp:        100,
					Membership: true,
					AvatarURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers7",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "10/02/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "10/02/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "29a1ee6d-05bd-4f11-b42a-a6769fd6a509",
					Name:      "Fani Nugraha",
					Password:  "pass1234",
					Email:     "fani.nugraha@example.com",
					Username:  "fani.nugraha",
					Address:   "Jl. Teratai No. 6, Medan",
					Gender:    "Perempuan",
					Phone:     "081234567895",
					Coin:      250,
					Exp:       125,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers8",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/02/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/02/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "fbfccf91-f6ea-4b4f-952e-f371e9c69d9c",
					Name:      "Gilang Permana",
					Password:  "qwerty123",
					Email:     "gilang.permana@example.com",
					Username:  "gilang.permana",
					Address:   "Jl. Anggrek No. 7, Bali",
					Gender:    "Laki-laki",
					Phone:     "081234567896",
					Coin:      80,
					Exp:       64,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers9",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "25/02/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "25/02/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "4fe8e112-e3d7-4468-bfd0-24058790bb6e",
					Name:      "Hana Putri",
					Password:  "letmein",
					Email:     "hana.putri@example.com",
					Username:  "hana.putri",
					Address:   "Jl. Dahlia No. 8, Makassar",
					Gender:    "Perempuan",
					Phone:     "081234567897",
					Coin:      110,
					Exp:       55,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers10",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "04/03/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "04/03/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "56a85666-b5aa-4b39-8823-99c7f789ac7b",
					Name:      "Indra Maulana",
					Password:  "welcome1",
					Email:     "indra.maulana@example.com",
					Username:  "indra.maulana",
					Address:   "Jl. Flamboyan No. 9, Malang",
					Gender:    "Laki-laki",
					Phone:     "081234567898",
					Coin:      90,
					Exp:       45,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers11",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "06/03/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "06/03/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "c0045c29-d1fa-4524-85fd-3e9341f99c6a",
					Name:      "Surya Rahmat",
					Password:  "trustno1",
					Email:     "suryarahmat@example.com",
					Username:  "surya.rahmat",
					Address:   "Jl. Cempaka No. 10, Solo",
					Gender:    "Laki-laki",
					Phone:     "081234567899",
					Coin:      0,
					Exp:       0,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers12",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "11/03/2024 8:59:59")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "11/03/2024 8:59:59")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "6220c455-f14e-42d3-a73d-2f4573f4cf52",
					Name:      "Kiki Amalia",
					Password:  "password456",
					Email:     "kiki.amalia@example.com",
					Username:  "kiki.amalia",
					Address:   "Jl. Kamboja No. 11, Pekanbaru",
					Gender:    "Perempuan",
					Phone:     "081234567800",
					Coin:      40,
					Exp:       20,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers13",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "11/03/2024 16:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "11/03/2024 16:00:00")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "35e243e9-d0b1-40b7-b216-5bc6cc15e243",
					Name:      "Lina Marlina",
					Password:  "mypassword2",
					Email:     "lina.marlina@example.com",
					Username:  "lina.marlina",
					Address:   "Jl. Bougenville No. 12, Palembang",
					Gender:    "Tidak Terdefinisi",
					Phone:     "081234567801",
					Coin:      0,
					Exp:       0,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers14",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "13/03/2024 16:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "13/03/2024 16:00:00")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "8f14519b-ea61-457b-bf96-90504221de90",
					Name:      "Mulyadi Saputra",
					Password:  "pass2022",
					Email:     "mulyadi.saputra@example.com",
					Username:  "mulyadi.saputra",
					Address:   "Jl. Tulip No. 13, Balikpapan",
					Gender:    "Tidak Terdefinisi",
					Phone:     "081234567802",
					Coin:      0,
					Exp:       0,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		{
			Name: "CreateUsers15",
			Run: func(db *gorm.DB) error {
				createdAt, err := time.Parse("02/01/2006 15:04:05", "20/03/2024 16:00:00")
				if err != nil {
					return err
				}

				updatedAt, err := time.Parse("02/01/2006 15:04:05", "20/03/2024 16:00:00")
				if err != nil {
					return err
				}
				user := user.User{
					Model: &gorm.Model{
						CreatedAt: createdAt,
						UpdatedAt: updatedAt,
					},
					ID:        "3fb1eb90-805f-4727-9f7d-a6fefc6115f9",
					Name:      "Nina Agustina",
					Password:  "password789",
					Email:     "nina.agustina@example.com",
					Username:  "nina.agustina",
					Address:   "Jl. Sakura No. 14, Pontianak",
					Gender:    "Tidak Terdefinisi",
					Phone:     "081234567803",
					Coin:      100,
					Exp:       50,
					AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
				}
				return CreateUsers(db, user)
			},
		},
		// End of Users
		// Start of Vouchers
		{
			Name: "CreateVoucher1",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "84c65b94-163f-4880-badd-ad10ff8bfead", "Promo Natal", "NATAL2024", "20000", "Diskon khusus untuk produk ramah lingkungan", time.Date(2024, 12, 25, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher2",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "f1098036-cfa6-48ae-882c-00726d0c5d57", "Promo Idul Fitri", "PAKETLEBARAN", "15%", "Paket penawaran barang-barang ramah lingkungan untuk Lebaran", time.Date(2025, 5, 15, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher3",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "7e9c42d0-707c-4f88-b098-ae47888104fa", "Promo Hari Kemerdekaan Indonesia", "MERDEKA2024", "15000", "Diskon produk daur ulang untuk 17 Agustus", time.Date(2025, 8, 17, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher4",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "268a7851-53dd-4bc3-af91-0c84d00ea727", "Harga Spesial Tahun Baru", "TAHUNBARU2024", "30%", "Penawaran khusus pada sumber daya terbarukan untuk Tahun Baru", time.Date(2025, 1, 1, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher5",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "0e07708c-a23e-41a4-a999-58d4576743e4", "Promo Hari Bumi", "HARIBUMI2024", "10000", "Diskon barang biodegradable untuk Hari Bumi", time.Date(2025, 4, 22, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher6",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "fb2643f7-796b-4a20-a3a7-bd009582215e", "Diskon Halloween", "HALLOWEEN2024", "25%", "Penghematan produk hemat energi untuk Halloween", time.Date(2024, 10, 31, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher7",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "756945dc-7906-47b4-95c5-650f428d7767", "Potongan Harga Spesial Hari Kartini", "KARTINI2024", "30000", "Diskon item rendah karbon untuk Hari Kartini", time.Date(2025, 4, 21, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher8",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "2e419999-cb91-4e36-a363-289976b7e83b", "Promo Idul Adha", "IDULADHA2024", "20%", "Penawaran spesial produk hemat air untuk Idul Adha", time.Date(2025, 7, 10, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher9",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "880d707a-744b-41b9-991f-65c306074c03", "Batik Nasional", "BATIK2024", "35%", "Diskon besar produk zero waste di Hari Batik Nasional", time.Date(2025, 10, 2, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher10",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "4b7746fd-42ad-4090-8de9-b2fab6e1921c", "Tahun Baru Imlek", "IMLEK2024", "25%", "Diskon produk organik dan alami untuk Tahun Baru Imlek", time.Date(2025, 2, 1, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher11",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "485d5b86-c1f9-45c1-a542-7b61c6feeecf", "Promo Hari Kasih Sayang", "VALENTINE2024", "12000", "Potongan harga untuk Menghemat barang-barang rumah tangga tidak beracun di Hari Valentine", time.Date(2025, 2, 14, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher12",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "bfbc3bd1-17fc-4766-b1ae-fa7535b665b3", "Promo Spesial Earth Hour", "JAMBUMI2024", "15000", "Penawaran produk kompos untuk Earth Hour", time.Date(2025, 3, 29, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher13",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "c6613ef8-30a7-43ec-8b1f-78ae2491c7e5", "Diskon Hari Pendidikan ", "PENDIDIKAN2024", "20%", "Diskon barang pakai ulang dalam rangka Hari Pendidikan Nasional", time.Date(2025, 5, 2, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher14",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "83c4db38-d196-49c8-9cb7-b53423e1658f", "Diskon Hari Kesehatan Nasional", "KESEHATAN2024", "30000", "Penawaran khusus untuk produk daur ulang dalam rangka Hari Kesehatan Nasional", time.Date(2025, 11, 12, 23, 59, 59, 0, time.UTC))
			},
		},
		{
			Name: "CreateVoucher15",
			Run: func(db *gorm.DB) error {
				return CreateVoucher(db, "b1b68578-cba7-4b9d-8656-004737edc836", "Potongan Harga Spesial Hari Pancasila", "PANCASILA2024", "15%", "Penghematan produk bersertifikat perdagangan yang adil untuk Hari Pancasila", time.Date(2025, 6, 1, 23, 59, 59, 0, time.UTC))
			},
		},
		// End Voucher
		// Start Challenge log
		{
			Name: "CreateTransaction1",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "7af9d694-63bf-42cd-9afc-8b00716a0083",
					UserID:    "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					VoucherID: "2e419999-cb91-4e36-a363-289976b7e83b",
					Address:   "Jl. Sudirman No. 123, Jakarta",
					Status:    "Berhasil",
					Total:     105000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction2",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "3b7e3cde-a4ff-4be5-a583-f76bcc3fdf77",
					UserID:    "5b9a23e5-e562-4fae-b11c-787e655a67d7",
					VoucherID: "2e419999-cb91-4e36-a363-289976b7e83b",
					Address:   "Jl. Gajah Mada No. 456, Surabaya",
					Status:    "Berhasil",
					Total:     70000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction3",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "9f180211-bd78-483b-b824-0125376ba593",
					UserID:    "d4a59d1b-38fa-4736-88d7-00cdc02ef5be",
					VoucherID: "7e9c42d0-707c-4f88-b098-ae47888104fa",
					Address:   "Jl. Pahlawan No. 789, Yogyakarta",
					Status:    "Gagal",
					Total:     70000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction4",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "1c16e3ac-9a15-4203-8ff2-09f448a86228",
					UserID:    "6932e20c-3c55-4d95-91c5-ce316cc5843f",
					VoucherID: "7e9c42d0-707c-4f88-b098-ae47888104fa",
					Address:   "Jl. Diponegoro No. 321, Bandung",
					Status:    "Berhasil",
					Total:     1600000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction5",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "f129e567-fc84-437e-a0fd-2b4b94dd90f9",
					UserID:    "8f25169f-5084-45c3-a99f-40daf59485d4",
					VoucherID: "7e9c42d0-707c-4f88-b098-ae47888104fa",
					Address:   "Jl. Brawijaya No. 654, Malang",
					Status:    "Berhasil",
					Total:     20000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction6",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "7671e308-40e3-4c0c-b64a-982d4bb1c7e1",
					UserID:    "fbfccf91-f6ea-4b4f-952e-f371e9c69d9c",
					VoucherID: "7e9c42d0-707c-4f88-b098-ae47888104fa",
					Address:   "Jl. Thamrin No. 987, Makassar",
					Status:    "Berhasil",
					Total:     70000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction7",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "41895902-5859-4f51-8023-5a4473978e08",
					UserID:    "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					VoucherID: "7e9c42d0-707c-4f88-b098-ae47888104fa",
					Address:   "Jl. Sudirman No. 123, Jakarta",
					Status:    "Berhasil",
					Total:     17000,
					Coin:      75,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction8",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "1e166b59-fdd2-4bce-93b6-d81d35723e41",
					UserID:    "5b9a23e5-e562-4fae-b11c-787e655a67d7",
					VoucherID: "880d707a-744b-41b9-991f-65c306074c03",
					Address:   "Jl. Gajah Mada No. 456, Surabaya",
					Status:    "Gagal",
					Total:     174000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction9",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "489b813b-ecdc-4fd6-90ea-3c692528a344",
					UserID:    "d4a59d1b-38fa-4736-88d7-00cdc02ef5be",
					VoucherID: "880d707a-744b-41b9-991f-65c306074c03",
					Address:   "Jl. Pahlawan No. 789, Yogyakarta",
					Status:    "Berhasil",
					Total:     900000,
					Coin:      25,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction10",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "e137483a-3426-4316-801e-4e992a6451c4",
					UserID:    "6932e20c-3c55-4d95-91c5-ce316cc5843f",
					VoucherID: "fb2643f7-796b-4a20-a3a7-bd009582215e",
					Address:   "Jl. Diponegoro No. 321, Bandung",
					Status:    "Berhasil",
					Total:     1500000,
					Coin:      165,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction11",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "0970df7d-7d7d-4d1c-b435-8536a35efb81",
					UserID:    "8f25169f-5084-45c3-a99f-40daf59485d4",
					VoucherID: "fb2643f7-796b-4a20-a3a7-bd009582215e",
					Address:   "Jl. Brawijaya No. 654, Malang",
					Status:    "Berhasil",
					Total:     80000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction12",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "b063bb48-baa7-4ac9-9a1a-db57214dd3fd",
					UserID:    "fbfccf91-f6ea-4b4f-952e-f371e9c69d9c",
					VoucherID: "fb2643f7-796b-4a20-a3a7-bd009582215e",
					Address:   "Jl. Thamrin No. 987, Makassar",
					Status:    "Gagal",
					Total:     245000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction13",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "44c09473-4b70-4849-90a7-8c5db2430e1f",
					UserID:    "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					VoucherID: "83c4db38-d196-49c8-9cb7-b53423e1658f",
					Address:   "Jl. Sudirman No. 123, Jakarta",
					Status:    "Berhasil",
					Total:     24000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction14",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "3cef7e44-efd9-4518-8f11-440f2db92b54",
					UserID:    "5b9a23e5-e562-4fae-b11c-787e655a67d7",
					VoucherID: "84c65b94-163f-4880-badd-ad10ff8bfead",
					Address:   "Jl. Gajah Mada No. 456, Surabaya",
					Status:    "Menunggu Pembayaran",
					Total:     1405000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		{
			Name: "CreateTransaction15",
			Run: func(db *gorm.DB) error {
				transaction := transaction.Transaction{
					ID:        "59a80fc1-c324-4d65-8cf2-fc3e17793412",
					UserID:    "d4a59d1b-38fa-4736-88d7-00cdc02ef5be",
					VoucherID: "84c65b94-163f-4880-badd-ad10ff8bfead",
					Address:   "Jl. Pahlawan No. 789, Yogyakarta",
					Status:    "Berhasil",
					Total:     955000,
					Coin:      0,
				}
				return CreateTransaction(db, transaction)
			},
		},
		// End Challenge
		// Start Forum
		{
			Name: "CreateForum1",
			Run: func(db *gorm.DB) error {
				forums := forum.Forum{
					ID:            "f9bc77cb-2b04-442e-ae9f-aa8f163781e0",
					Title:         "Praktik Berkebun Berkelanjutan",
					UserID:        "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					Description:   "Diskusi tentang metode berkebun yang ramah lingkungan untuk mengurangi dampak lingkungan.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forums)
			},
		},
		{
			Name: "CreateForum2",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "69f4777d-6446-4903-b444-088bcf5a5978",
					Title:         "Pertanian Perkotaan dan Kebun Komunitas",
					UserID:        "6932e20c-3c55-4d95-91c5-ce316cc5843f",
					Description:   "Berbagi pengalaman dan ide-ide untuk membudidayakan makanan di lingkungan perkotaan.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum3",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "e06255df-6eb4-4ff0-bf8f-bf0efbf27e22",
					Title:         "Solusi Energi Terbarukan",
					UserID:        "6932e20c-3c55-4d95-91c5-ce316cc5843f",
					Description:   "Forum untuk berbagi pengetahuan dan pengalaman dengan teknologi energi terbarukan.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum4",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "d76ffcc9-b17c-4c40-b6da-2236a41b1120",
					Title:         "Diskusi Kebijakan Energi Hijau",
					UserID:        "fbfccf91-f6ea-4b4f-952e-f371e9c69d9c",
					Description:   "Terlibat dalam diskusi tentang kebijakan dan peraturan pemerintah terkait energi terbarukan.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum5",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "a3da4123-c266-4e7f-aabf-e3e8835c6b97",
					Title:         "Alternatif Transportasi Berkelanjutan",
					UserID:        "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					Description:   "Diskusikan penggunaan kendaraan listrik, transportasi umum, dan opsi komuter hijau lainnya.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum6",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "4d54c743-74d0-4486-9df8-a3952fc1e1c0",
					Title:         "Daur Ulang dan Manajemen Sampah",
					UserID:        "c0045c29-d1fa-4524-85fd-3e9341f99c6a",
					Description:   "Diskusi tentang praktik terbaik untuk mendaur ulang dan mengurangi sampah secara efektif.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum7",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "7c7c311b-58ff-48da-b649-3a1c17e649f2",
					Title:         "Desain Bangunan Hijau",
					UserID:        "8f14519b-ea61-457b-bf96-90504221de90",
					Description:   "Eksplorasi teknik arsitektur dan konstruksi yang berkelanjutan untuk bangunan ramah lingkungan.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum8",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "84580f82-f68c-46be-9f31-8dd2e7353e45",
					Title:         "Mode dan Tekstil Berkelanjutan",
					UserID:        "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					Description:   "Eksplorasi alternatif ramah lingkungan di industri fashion dan tekstil.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum9",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "71ee6052-75f0-4c5f-802c-ac5a6e785626",
					Title:         "Upaya Konservasi Lingkungan",
					UserID:        "8f14519b-ea61-457b-bf96-90504221de90",
					Description:   "Diskusikan inisiatif dan strategi untuk melindungi habitat alam dan satwa liar.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		{
			Name: "CreateForum10",
			Run: func(db *gorm.DB) error {
				forum := forum.Forum{
					ID:            "c9f69700-7534-4285-81a7-22d71cdfd09b",
					Title:         "Pilihan Gaya Hidup Ramah Lingkungan",
					UserID:        "c0045c29-d1fa-4524-85fd-3e9341f99c6a",
					Description:   "Eksplorasi cara untuk mengadopsi gaya hidup yang lebih berkelanjutan dan ramah lingkungan.",
					LastMessageAt: time.Now(),
				}
				return CreateForum(db, forum)
			},
		},
		// End Forum
		// Start Message Forum
		{
			Name: "CreateMessageForum1",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "771e3b98-cdd4-46e2-8f80-ef1e3136f9f8",
					UserID:  "5dcdb121-85df-487c-8d4a-b4fd5033c9c0",
					ForumID: "c9f69700-7534-4285-81a7-22d71cdfd09b",
					Message: "Haloo guys",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum2",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "2407dc1d-8152-4a99-ad13-a814b9712437",
					UserID:  "5b9a23e5-e562-4fae-b11c-787e655a67d7",
					ForumID: "69f4777d-6446-4903-b444-088bcf5a5978",
					Message: "Pagii teman-teman",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum3",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "a78529e3-d6d6-4212-842a-a083b846b594",
					UserID:  "d4a59d1b-38fa-4736-88d7-00cdc02ef5be",
					ForumID: "f9bc77cb-2b04-442e-ae9f-aa8f163781e0",
					Message: "Apa Kabarnya kalian?",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum4",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "cc9d6cb1-a3a9-43b3-972d-df3ee968b2a4",
					UserID:  "6932e20c-3c55-4d95-91c5-ce316cc5843f",
					ForumID: "7c7c311b-58ff-48da-b649-3a1c17e649f2",
					Message: "Tetap semangat!!",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum5",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "b3b3a835-edc5-4970-b512-11ab983a9b98",
					UserID:  "8f25169f-5084-45c3-a99f-40daf59485d4",
					ForumID: "4d54c743-74d0-4486-9df8-a3952fc1e1c0",
					Message: "Selamat pagi semuanya!",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum6",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "db85c8d1-54b7-43eb-9673-8fef60135526",
					UserID:  "b7352ca3-fae1-4f3c-85b9-095de368b706",
					ForumID: "71ee6052-75f0-4c5f-802c-ac5a6e785626",
					Message: "Bagaimana hari kalian?",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum7",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "cd388472-8641-40eb-b5c3-72b03424aae0",
					UserID:  "29a1ee6d-05bd-4f11-b42a-a6769fd6a509",
					ForumID: "a3da4123-c266-4e7f-aabf-e3e8835c6b97",
					Message: "Semoga harimu menyenangkan.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum8",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "0798fab9-4db7-4e58-8a87-49a475a34add",
					UserID:  "fbfccf91-f6ea-4b4f-952e-f371e9c69d9c",
					ForumID: "e06255df-6eb4-4ff0-bf8f-bf0efbf27e22",
					Message: "Jangan lupa sarapan.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum9",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "c073c6cb-f873-456e-85c1-2b74d696bcf6",
					UserID:  "4fe8e112-e3d7-4468-bfd0-24058790bb6e",
					ForumID: "84580f82-f68c-46be-9f31-8dd2e7353e45",
					Message: "Semangat menjalani hari ini.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum10",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "b454c9de-5d87-462e-8fe7-54b2f8aca53e",
					UserID:  "56a85666-b5aa-4b39-8823-99c7f789ac7b",
					ForumID: "d76ffcc9-b17c-4c40-b6da-2236a41b1120",
					Message: "Ada yang punya rencana seru hari ini?",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum11",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "73cf3dc6-6ddf-44b4-8965-a0fb2873cc01",
					UserID:  "c0045c29-d1fa-4524-85fd-3e9341f99c6a",
					ForumID: "c9f69700-7534-4285-81a7-22d71cdfd09b",
					Message: "Selamat siang teman-teman.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum12",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "710f3757-90ae-4893-b596-5b90f8bb4555",
					UserID:  "6220c455-f14e-42d3-a73d-2f4573f4cf52",
					ForumID: "69f4777d-6446-4903-b444-088bcf5a5978",
					Message: "Jangan lupa minum air putih.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum13",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "8984b04a-57b2-4f18-8ad7-1f44ed6ebe5c",
					UserID:  "35e243e9-d0b1-40b7-b216-5bc6cc15e243",
					ForumID: "f9bc77cb-2b04-442e-ae9f-aa8f163781e0",
					Message: "Sudah makan siang belum?",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum14",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "2dfc2b53-5a6e-4058-868c-384d0f845ac0",
					UserID:  "8f14519b-ea61-457b-bf96-90504221de90",
					ForumID: "7c7c311b-58ff-48da-b649-3a1c17e649f2",
					Message: "Semoga harimu produktif.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum15",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "d75e6c42-2747-4d69-a676-ab6aa8088c48",
					UserID:  "3fb1eb90-805f-4727-9f7d-a6fefc6115f9",
					ForumID: "4d54c743-74d0-4486-9df8-a3952fc1e1c0",
					Message: "Tetap jaga kesehatan.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum16",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "d0644609-f5be-4050-80a5-3651b31efe96",
					UserID:  "6220c455-f14e-42d3-a73d-2f4573f4cf52",
					ForumID: "71ee6052-75f0-4c5f-802c-ac5a6e785626",
					Message: "Bagaimana pekerjaan kalian?",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum17",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "b0257295-327a-4239-b119-d7c5e6ce103d",
					UserID:  "c0045c29-d1fa-4524-85fd-3e9341f99c6a",
					ForumID: "a3da4123-c266-4e7f-aabf-e3e8835c6b97",
					Message: "Jangan terlalu stress, ya.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum18",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "022cb108-ee7f-4b85-a1cb-b99c29190c29",
					UserID:  "35e243e9-d0b1-40b7-b216-5bc6cc15e243",
					ForumID: "e06255df-6eb4-4ff0-bf8f-bf0efbf27e22",
					Message: "Selamat sore semuanya.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum19",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "30e430fb-3e1f-48be-8dfe-89e7e13726cb",
					UserID:  "29a1ee6d-05bd-4f11-b42a-a6769fd6a509",
					ForumID: "84580f82-f68c-46be-9f31-8dd2e7353e45",
					Message: "Apa yang kalian lakukan sore ini?",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
		{
			Name: "CreateMessageForum20",
			Run: func(db *gorm.DB) error {
				messageForum := forum.MessageForum{
					ID:      "48f34810-1b26-49b2-8f73-dd8d83eeaefe",
					UserID:  "b7352ca3-fae1-4f3c-85b9-095de368b706",
					ForumID: "d76ffcc9-b17c-4c40-b6da-2236a41b1120",
					Message: "Jangan lupa istirahat cukup malam ini.",
				}
				return CreateForumMessage(db, messageForum)
			},
		},
	}
	return seeds
}
