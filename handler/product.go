package handler

import (
	"fmt"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	mongodatabase "github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateProduct() gin.HandlerFunc{
	return func(c *gin.Context){

		var reqProduct mongodatabase.Products 
		err := c.BindJSON(&reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", reqProduct)

		reqProduct.Id =primitive.NewObjectID()
		// reqProduct.Created_at =time.Now().UTC()
		count,err := mongodatabase.CountBrand(reqProduct.Name)

		// SEARCH EMAIL
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			logger.Error().Msg("❌🔥 Error message :" + "already created product")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already created product"})
			return
		}

		res,err:=mongodatabase.InsertProduct(reqProduct)
		if err == nil {
			logger.Info().Msg("📢 Info message :" + "successfully created Product")
		}
		_ = res
        fmt.Println("🚀 ~ file: category.go ~ line 46 ~ returnfunc ~ res : ", res)
		c.JSON(http.StatusAccepted, res)
	}
}


func GetAllProducts() gin.HandlerFunc{
	return func(c *gin.Context){
		product,err := mongodatabase.GetAllProduct()
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}
func GetProduct() gin.HandlerFunc{
	
	return func(c *gin.Context){
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		product,err := mongodatabase.GetProduct(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,product)
	}
}

func UpdateProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		var reqProduct mongodatabase.Products 
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := c.BindJSON(&reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// reqBrand.Created_at =time.Now().UTC()

		product,err := mongodatabase.UpdateProduct(objectID,reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,product)
	}
}


func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context){
		
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		res,err := mongodatabase.DeleteProduct(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,res)
	}
}