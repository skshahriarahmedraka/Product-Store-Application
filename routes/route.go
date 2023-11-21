package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine) {

	r.GET("/health", handler.HandlerHealth())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/brand", handler.CreateBrand())
	r.GET("/brand/:id", handler.GetBrand())
	r.PUT("/brand/:id", handler.UpdateBrand())
	r.DELETE("/brand/:id", handler.DeleteBrand())
	r.GET("/brands", handler.GetAllBrands())

	r.POST("/category", handler.CreateCategory())
	r.GET("/category/:id", handler.GetCategory())
	r.PUT("/category/:id", handler.UpdateCategory())
	r.DELETE("/category/:id", handler.DeleteCategory())
	r.GET("/category", handler.GetAllCategories())

	r.POST("/supplier", handler.CreateSupplier())
	r.GET("/supplier/:id", handler.GetSupplier())
	r.PUT("/supplier/:id", handler.UpdateSupplier())
	r.DELETE("/supplier/:id", handler.DeleteSupplier())
	r.GET("/suppliers", handler.GetAllSuppliers())

	r.POST("/product", handler.CreateProduct())
	r.GET("/product/:id", handler.GetProduct())
	r.PUT("/product/:id", handler.UpdateProduct())
	r.DELETE("/product/:id", handler.DeleteProduct())
	r.GET("/products", handler.GetAllProducts())

	// r.POST("/product/stock", CreateProductStock())
	r.GET("/product/stock/:id", handler.GetProductStock())
	// r.PUT("/product/stock/:id", UpdateProductStock())
	// r.DELETE("/product/stock/:id", DeleteProductStock())
	r.GET("/product/stocks", handler.GetAllProductStocks())

	r.POST("/productbyname", handler.GetProductByName())
	r.GET("/productbyprice/:price", handler.GetProductByPrice())
	r.POST("/productbymulbrand", handler.GetProductByMulBrand())
	r.GET("/productbycategory/:id", handler.GetProductByCategory())
	r.GET("/productbysupplier/:id", handler.GetProductBySupplier())
	r.GET("/productbyverifiedsupplier/:id", handler.GetProductByVerifiedSupplier())

}
