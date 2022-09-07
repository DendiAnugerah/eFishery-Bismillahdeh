package main

import (
	"Fish/config"
	categorycontroller "Fish/controller/category"
	productcontroller "Fish/controller/product"
	categoryrepository "Fish/repository/category"
	productrepository "Fish/repository/product"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Temen-temen sudah dapat tema ya, temanya e-commerce. Tugasnya bikin REST API yang isinya:

// 1. Detail produk -> nama , foto, harga, deskripsi
// 2. Filter produk -> berdasarkan kategori, berdasarkan range harga
// 3. Lissting produk -> nama, foto, harga, kategori
// 4. Pembayaran -> bayar sesuai total pembelian dan upload bukti pembayaran (cukup upload saja, tidak perlu verifikasi)
// 5. Tambahkan produk ke troli atau cart

// Catatan:
// 1. Dokumentasi REST API bisa pakai swagger atau buat tabel di spreadsheet
// 2. Tidak perlu UI. Semu via REST API
// 3. Demo bisa menggunakan Postman
// 4. Lama pengerjaan 1 minggu
// 5. Akan ada sesi speeddating dengan eFisherian dan User

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	e := echo.New()
	DB, _ := config.Conection()

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	productR := productrepository.NewProdRepository(DB)
	CategoryR := categoryrepository.NewCatRepository(DB)

	productC := productcontroller.NewProdController(productR, CategoryR)
	CategoryC := categorycontroller.NewCatController(CategoryR)

	// Product
	e.POST("/product/create", productC.CreateProduct)
	e.GET("/product/get", productC.GetProducts)
	e.GET("/product/get/:id", productC.GetProduct)
	e.PUT("/product/update/:id", productC.UpdateProduct)
	e.DELETE("/product/delete/:id", productC.DeleteProduct)

	// Category
	e.POST("/category/create", CategoryC.CreateCategory)
	e.GET("/category/get", CategoryC.GetAllCategory)
	e.GET("/category/get/:id", CategoryC.GetCategory)
	e.PUT("/category/update/:id", CategoryC.UpdateCategory)
	e.DELETE("/category/delete/:id", CategoryC.DeleteCategory)

	e.Logger.Fatal(e.Start(":9000"))
}
