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
				iconUrl := "https://storage.googleapis.com/alterra-greeve/greeve/9e01061a-e622-414a-a3f1-27263620c0cbicon2.svg"
				return CreateImpactCategory(db, "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd", "Hemat Uang", 40, description, imageUrl, iconUrl)
			},
		},
		{
			Name: "MengurangiLimbah",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Less Waste</strong> atau <strong>Mengurangi Limbah</strong> adalah pendekatan yang menekankan pada pengurangan jumlah limbah yang dihasilkan, baik oleh individu, komunitas, maupun perusahaan. Tujuannya adalah untuk meminimalisir dampak negatif terhadap lingkungan dengan cara mengelola sumber daya secara lebih efisien dan bertanggung jawab. Strategi ini mencakup berbagai tindakan seperti daur ulang, pengomposan, penggunaan ulang barang, dan pengurangan penggunaan bahan sekali pakai.</p><p>Contoh Dampak Positif Less Waste pada Lingkungan:</p><ul><li>Pengurangan Polusi Tanah dan Air</li><li>Penurunan Emisi Gas Rumah Kaca</li><li>Konservasi Sumber Daya Alam</li><li>Mengurangi Energi yang Dibutuhkan untuk Produksi</li><li>Mengurangi Beban Tempat Pembuangan Akhir</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/91d948ca-3576-49e5-b435-0c3c7eccff7963ca36bac45556620e0ff4ba7ec5790d.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/greeve/25e1e158-3fb0-42ff-97f7-d728167f2a7eicon4.svg"
				return CreateImpactCategory(db, "7d34a5fa-e2cf-466d-9f01-d731f6967082", "Mengurangi Limbah", 50, description, imageUrl, iconUrl)
			},
		},
		{
			Name: "PerluasWawasan",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Expand Your Mind</strong> atau <strong>Perluas Wawasan Anda</strong> adalah konsep yang mendorong individu untuk terbuka terhadap pengetahuan, ide, dan pengalaman baru. Ini melibatkan pembelajaran berkelanjutan, eksplorasi topik yang beragam, dan berpikir kritis. Dengan memperluas wawasan, kita tidak hanya meningkatkan pemahaman kita tentang dunia, tetapi juga mengembangkan empati, kreativitas, dan kemampuan untuk membuat keputusan yang lebih baik.</p><p>Contoh Dampak Positif Save Money pada Lingkungan:</p><ul><li>Kesadaran Lingkungan yang Lebih Tinggi</li><li>Inovasi dalam Solusi Ramah Lingkungan</li><li>Dukungan untuk Kebijakan Lingkungan yang Lebih Baik</li><li>Pengembangan Komunitas Berkelanjutan</li><li>Konsumsi yang Lebih Bertanggung Jawab</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/fd98dcdc-4cea-428a-a008-e1ca6f972ca72f0376e551eec4f1af915d5983a621c1.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/greeve/1f3dbf92-1d5a-4eb1-9e1f-385c6ea89220icon3.svg"
				return CreateImpactCategory(db, "e8e714bd-c34e-4278-980c-39bd1f55b5fb", "Perluas Wawasan", 45, description, imageUrl, iconUrl)
			},
		},
		{
			Name: "MengurangiPemanasanGlobal",
			Run: func(db *gorm.DB) error {
				description := "<div><p><strong>Less Global Warming</strong> atau <strong>Mengurangi Pemanasan Global</strong> adalah upaya untuk menurunkan peningkatan suhu rata-rata bumi yang disebabkan oleh aktivitas manusia, terutama melalui emisi gas rumah kaca seperti karbon dioksida (CO<sub>2</sub>) dan metana (CH<sub>4</sub>). Upaya ini mencakup berbagai tindakan, seperti meningkatkan efisiensi energi, menggunakan energi terbarukan, mengurangi deforestasi, dan mempromosikan praktik pertanian berkelanjutan. Tujuannya adalah untuk memperlambat perubahan iklim, melindungi ekosistem, dan menjaga keseimbangan alami bumi.</p><p>Contoh Dampak Positif Save Money pada Lingkungan:</p><ul><li>Pengurangan Fenomena Cuaca Ekstrem</li><li>Perlindungan Keanekaragaman Hayati</li><li>Stabilisasi Tingkat Permukaan Laut</li><li>Kualitas Udara yang Lebih Baik</li><li>Konservasi Sumber Daya Air</li></ul></div>"
				imageUrl := "https://storage.googleapis.com/alterra-greeve/greeve/38c805dd-9fbf-4fa0-bef7-9adfce99add9d5787157f3d099da70dcbf8f3021c99b.jpeg"
				iconUrl := "https://storage.googleapis.com/alterra-greeve/greeve/eb285b2b-19b2-447e-894e-81a12fbfb517icon1.svg"
				return CreateImpactCategory(db, "b5d07366-3b31-4011-95e3-34735b0b61f8", "Mengurangi Pemanasan Global", 35, description, imageUrl, iconUrl)
			},
		},
		// Impact Categoreis End
		//Product Seeds Start
		{
			Name: "CreateProduct1",
			Run: func(db *gorm.DB) error {
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
				products := product.Product{
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
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/a5f87eb5-b186-4b3c-bca8-9baf76c74904goodearth-goodearthxashiesh-2020-d1.webp",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "7e7a386a-c8c9-430c-b7d6-2e81f1bad678",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				products := product.Product{
					ID:          "883b2f74-5f4a-4e75-8e8d-07d6df2f9bde",
					Name:        "Baterai Isi Ulang",
					Description: "Baterai AA yang dapat diisi ulang, membantu mengurangi limbah baterai sekali pakai.",
					Price:       100000,
					Coin:        25,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "883b2f74-5f4a-4e75-8e8d-07d6df2f9bde",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/2ac41750-85e3-45f8-95fa-46fa9f4a0c06baterai_isi_ulang.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "883b2f74-5f4a-4e75-8e8d-07d6df2f9bde",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
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
				products := product.Product{
					ID:          "7af1e7fa-55b3-4711-b8ed-e91d9c0bb9de",
					Name:        "Lampu LED Hemat Energi",
					Description: "Lampu LED yang hemat energi dan tahan lama, mengurangi konsumsi listrik.",
					Price:       50000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "7af1e7fa-55b3-4711-b8ed-e91d9c0bb9de",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/03afaa45-b95b-41bb-aaf0-fd2b5d1e7c02led_light_bulb.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "7af1e7fa-55b3-4711-b8ed-e91d9c0bb9de",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
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
				products := product.Product{
					ID:          "15c9a672-1c1d-41c4-9e65-5c0d2d92b722",
					Name:        "Botol Minum Stainless Steel",
					Description: "Botol minum yang dapat digunakan kembali dan ramah lingkungan, terbuat dari stainless steel berkualitas tinggi.",
					Price:       75000,
					Coin:        15,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "15c9a672-1c1d-41c4-9e65-5c0d2d92b722",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/02c6f527-c9db-414f-bc08-5d96210b35dfbottle_stainless_steel.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "15c9a672-1c1d-41c4-9e65-5c0d2d92b722",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "15c9a672-1c1d-41c4-9e65-5c0d2d92b722",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
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
				products := product.Product{
					ID:          "6e5a4e94-30f3-4a1e-87d1-0dace147e1bb",
					Name:        "Tas Belanja Lipat",
					Description: "Tas belanja yang dapat dilipat dan digunakan kembali, mengurangi penggunaan kantong plastik sekali pakai.",
					Price:       60000,
					Coin:        12,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "6e5a4e94-30f3-4a1e-87d1-0dace147e1bb",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8c5b6440-79c2-4e58-9fd3-9f6a4a35f4f3foldable_shopping_bag.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "6e5a4e94-30f3-4a1e-87d1-0dace147e1bb",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/8c5b6440-79c2-4e58-9fd3-9f6a4a35f4f3foldable_shopping_bag_open.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "6e5a4e94-30f3-4a1e-87d1-0dace147e1bb",
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
				products := product.Product{
					ID:          "8b51210c-c7a8-4e74-88f6-c59ed50333a4",
					Name:        "Sabun Batang Alami",
					Description: "Sabun batang yang terbuat dari bahan-bahan alami, bebas dari bahan kimia keras.",
					Price:       30000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "8b51210c-c7a8-4e74-88f6-c59ed50333a4",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/89c3d505-e332-44ad-97fa-46c5d03b60ccnatural_soap_bar.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "8b51210c-c7a8-4e74-88f6-c59ed50333a4",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/89c3d505-e332-44ad-97fa-46c5d03b60ccnatural_soap_bar_packaging.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "8b51210c-c7a8-4e74-88f6-c59ed50333a4",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
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
				products := product.Product{
					ID:          "27a731c4-e539-4a13-8b7f-6cf1a7c5d6d5",
					Name:        "Sikat Gigi Bambu",
					Description: "Sikat gigi yang terbuat dari bambu yang dapat terurai secara alami, alternatif yang lebih baik dari sikat gigi plastik.",
					Price:       20000,
					Coin:        8,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "27a731c4-e539-4a13-8b7f-6cf1a7c5d6d5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/50e0040b-0c4a-4d3b-8325-8490b8b96740bamboo_toothbrush.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "27a731c4-e539-4a13-8b7f-6cf1a7c5d6d5",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/50e0040b-0c4a-4d3b-8325-8490b8b96740bamboo_toothbrush_packaging.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "27a731c4-e539-4a13-8b7f-6cf1a7c5d6d5",
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
				products := product.Product{
					ID:          "ab6d0e7b-8c08-498e-8d91-39c530bb6c62",
					Name:        "Kompos Bin",
					Description: "Kompos bin yang mudah digunakan untuk mendaur ulang limbah organik di rumah.",
					Price:       150000,
					Coin:        30,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "ab6d0e7b-8c08-498e-8d91-39c530bb6c62",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/67e9c472-2b85-4a55-b499-3ff46f524f31compost_bin.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "ab6d0e7b-8c08-498e-8d91-39c530bb6c62",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/67e9c472-2b85-4a55-b499-3ff46f524f31compost_bin_open.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "ab6d0e7b-8c08-498e-8d91-39c530bb6c62",
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
				products := product.Product{
					ID:          "1e570071-cfe2-46b1-a78b-e02b9ae35c0b",
					Name:        "Kantong Sampah Kompos",
					Description: "Kantong sampah yang dapat terurai secara alami dan cocok untuk kompos.",
					Price:       50000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "1e570071-cfe2-46b1-a78b-e02b9ae35c0b",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/f7e9d9d4-7f68-4741-bdf3-c5f2ef10e3a2compostable_garbage_bags.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "1e570071-cfe2-46b1-a78b-e02b9ae35c0b",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
				products := product.Product{
					ID:          "4e0740c7-cb0e-43a3-aea6-66bdf5f9e57e",
					Name:        "Gelas Kopi Bawa Pulang",
					Description: "Gelas kopi yang dapat digunakan kembali dan ramah lingkungan, cocok untuk kopi bawa pulang.",
					Price:       75000,
					Coin:        15,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "4e0740c7-cb0e-43a3-aea6-66bdf5f9e57e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/78d75fa0-f104-4d69-aed9-fd60b46ff8a7reusable_coffee_cup.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "4e0740c7-cb0e-43a3-aea6-66bdf5f9e57e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/78d75fa0-f104-4d69-aed9-fd60b46ff8a7reusable_coffee_cup_lid.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "4e0740c7-cb0e-43a3-aea6-66bdf5f9e57e",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "4e0740c7-cb0e-43a3-aea6-66bdf5f9e57e",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
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
				products := product.Product{
					ID:          "ec7f68b0-f109-4cb3-8a68-5f7f1b3c593a",
					Name:        "Penghilang Bau Alami",
					Description: "Penghilang bau yang terbuat dari bahan-bahan alami, bebas dari bahan kimia berbahaya.",
					Price:       40000,
					Coin:        10,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "ec7f68b0-f109-4cb3-8a68-5f7f1b3c593a",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/1c3a77de-718f-4a91-b06e-4a819ef86dbdnatural_deodorizer.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "ec7f68b0-f109-4cb3-8a68-5f7f1b3c593a",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
					},
					{
						ID:               uuid.New().String(),
						ProductID:        "ec7f68b0-f109-4cb3-8a68-5f7f1b3c593a",
						ImpactCategoryID: "b5d07366-3b31-4011-95e3-34735b0b61f8",
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
				products := product.Product{
					ID:          "f3d5b5d5-7d3f-48b4-9283-4741aef6a8a8",
					Name:        "Talenan Kayu",
					Description: "Talenan yang terbuat dari kayu berkualitas tinggi, alternatif yang ramah lingkungan dari talenan plastik.",
					Price:       120000,
					Coin:        20,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "f3d5b5d5-7d3f-48b4-9283-4741aef6a8a8",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/4b17a6de-3f37-4e0b-8d6f-3fda5167f849wooden_cutting_board.jpg",
						Position:  1,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "f3d5b5d5-7d3f-48b4-9283-4741aef6a8a8",
						ImpactCategoryID: "83808762-e2b8-4b34-a1eb-0ed8d4fda3dd",
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
				products := product.Product{
					ID:          "aedc1ef4-b4aa-45a4-b1e8-345ca7e55f0e",
					Name:        "Tempat Penyimpanan Kain",
					Description: "Tempat penyimpanan yang terbuat dari kain yang dapat digunakan kembali dan tahan lama.",
					Price:       85000,
					Coin:        15,
				}
				images := []product.ProductImage{
					{
						ID:        uuid.New().String(),
						ProductID: "aedc1ef4-b4aa-45a4-b1e8-345ca7e55f0e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/9e65e2a2-5d64-4ae1-87d7-8f9e3f3c607bcloth_storage_bin.jpg",
						Position:  1,
					},
					{
						ID:        uuid.New().String(),
						ProductID: "aedc1ef4-b4aa-45a4-b1e8-345ca7e55f0e",
						ImageURL:  "https://storage.googleapis.com/alterra-greeve/greeve/9e65e2a2-5d64-4ae1-87d7-8f9e3f3c607bcloth_storage_bin_open.jpg",
						Position:  2,
					},
				}
				categories := []product.ProductImpactCategory{
					{
						ID:               uuid.New().String(),
						ProductID:        "aedc1ef4-b4aa-45a4-b1e8-345ca7e55f0e",
						ImpactCategoryID: "7d34a5fa-e2cf-466d-9f01-d731f6967082",
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
					Title:       "Menanam Pohon",
					Difficulty:  "Sedang",
					Description: "Tanam pohon di halaman belakang atau taman komunitas",
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
				challenges := challenge.Challenge{
					ID:          "ba2621ac-6f99-43cb-ab8a-7db626b7c4e4",
					Title:       "Minggu Tanpa Limbah",
					Difficulty:  "Sulit",
					Description: "Tidak menghasilkan limbah selama seminggu dan daur ulang segala sesuatu yang mungkin",
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
				challenges := challenge.Challenge{
					ID:          "5086c59a-e2f9-44d0-bcc6-f7fdb58f07a4",
					Title:       "Kerja Bakti Membersihkan Lingkungan",
					Difficulty:  "Mudah",
					Description: "Ikut serta dalam acara kerja bakti membersihkan lingkungan setempat",
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
				challenges := challenge.Challenge{
					ID:          "aa013dc7-4fb3-4cf1-9cc0-e1a0ad60414d",
					Title:       "Bersepeda ke Tempat Kerja",
					Difficulty:  "Sedang",
					Description: "Berangkat kerja dengan sepeda selama seminggu",
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
				challenges := challenge.Challenge{
					ID:          "5550e03b-15c6-43c9-997c-4086adc8b573",
					Title:       "Bulan Hemat Energi",
					Difficulty:  "Sulit",
					Description: "Kurangi konsumsi energi rumah tangga Anda sebesar 20% selama sebulan",
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
				challenges := challenge.Challenge{
					ID:          "9b21b3e5-46ff-42a1-92cc-4969c078ff83",
					Title:       "Membuat Produk Ramah Lingkungan",
					Difficulty:  "Sedang",
					Description: "Buat sendiri dan gunakan tiga produk rumah tangga ramah lingkungan",
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
		// End Voucher
		// Start Challenge log
		// End Challenge

		// Start Challenge Confirmation
		// End Challenge Confirmation
	}
	return seeds
}
