package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqProduct mongodatabase.Products
		err := c.BindJSON(&reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", reqProduct)
		reqProduct.Id = primitive.NewObjectID()
		count, err := mongodatabase.CountBrandById(reqProduct.Brand_id)

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count <= 0 {
			logger.Error().Msg("❌🔥 Error message :" + "Brand Doesn't exist")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Brand Doesn't exist"})
			return
		}

		count, err = mongodatabase.CountCategoryById(reqProduct.Catagory_id)

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count <= 0 {
			logger.Error().Msg("❌🔥 Error message :" + "Category Doesn't exist")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category Doesn't exist"})
			return
		}

		count, err = mongodatabase.CountSupplierById(reqProduct.Supplier_id)

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count <= 0 {
			logger.Error().Msg("❌🔥 Error message :" + "supplier Doesn't exist")
			c.JSON(http.StatusBadRequest, gin.H{"error": "supplier Doesn't exist"})
			return
		}

		count, err = mongodatabase.DublicateBySupplier(reqProduct.Name, reqProduct.Supplier_id)

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			logger.Error().Msg("❌🔥 Error message :" + "Product already exist in this supplier")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product already exist in this supplier"})
			return
		}

		res, err := mongodatabase.AddProductInStock(reqProduct.Id)
		_ = res
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res, err = mongodatabase.InsertProduct(reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		_ = res
		fmt.Println("🚀 ~ file: category.go ~ line 46 ~ returnfunc ~ res : ", res)
		c.JSON(http.StatusAccepted, res)
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		product, err := mongodatabase.GetAllProduct()
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}
func GetProduct() gin.HandlerFunc {

	return func(c *gin.Context) {
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		product, err := mongodatabase.GetProduct(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqProduct mongodatabase.Products
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := c.BindJSON(&reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := mongodatabase.UpdateProduct(objectID, reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		res, err := mongodatabase.DeleteProduct(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
