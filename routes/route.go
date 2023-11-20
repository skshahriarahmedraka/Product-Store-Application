package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
	// "github.com/skshahriarahmedraka/Product-Store-Application/pkg/monogodb"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// swaggerFiles "github.com/swaggo/files"
)

func Routes(r *gin.Engine) {
	// mongoConn := monodb.MongodbConnection()
	// H := monodb.DatabaseInitialization(mongoConn)

	r.GET("/health",handler.HandlerHealth())
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.GET("/public", Home())
	// // r.POST("/login", H.Login())
	// // r.POST("/register", H.Register())
	// // r.GET("/logout", H.Logout())
	// // r.GET("/:id", H.UserData())
	// // r.GET("/alluser", H.AllUserData())

	// r.POST("/brand", CreateBrand())
	// r.GET("/brand/:id", GetBrand())
	// r.PUT("/brand/:id", UpdateBrand())
	// r.DELETE("/brand/:id", DeleteBrand())
	// r.GET("/brands", GetAllBrands())

	r.POST("/category", handler.CreateCategory())
	r.GET("/category/:id", handler.GetCategory())
	r.PUT("/category/:id", handler.UpdateCategory())
	r.DELETE("/category/:id", handler.DeleteCategory())
	r.GET("/category", handler.GetAllCategories())
	
	// r.POST("/supplier", CreateSupplier())
	// r.GET("/supplier/:id", GetSupplier())
	// r.PUT("/supplier/:id", UpdateSupplier())
	// r.DELETE("/supplier/:id", DeleteSupplier())
	// r.GET("/suppliers", GetAllSuppliers())



	// r.POST("/product", CreateProduct())
	// r.GET("/product/:id", GetProduct())
	// r.PUT("/product/:id", UpdateProduct())
	// r.DELETE("/product/:id", DeleteProduct())
	// r.GET("/products", GetAllProducts())

	// r.POST("/product/stock", CreateProductStock())
	// r.GET("/product/stock/:id", GetProductStock())
	// r.PUT("/product/stock/:id", UpdateProductStock())
	// r.DELETE("/product/stock/:id", DeleteProductStock())
	// r.GET("/product/stocks", GetAllProductStocks())
}
