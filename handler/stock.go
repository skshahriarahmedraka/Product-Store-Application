package handler

import (
	// "fmt"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	mongodatabase "github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)




func GetAllProductStocks() gin.HandlerFunc{
	return func(c *gin.Context){
		brands,err := mongodatabase.GetAllBrand()
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, brands)
	}
}
func GetProductStock() gin.HandlerFunc{
	
	return func(c *gin.Context){
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		brand,err := mongodatabase.GetBrand(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,brand)
	}
}

