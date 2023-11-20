package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	mongodatabase "github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateCategory() gin.HandlerFunc{
	return func(c *gin.Context){

		var reqCategory mongodatabase.Catagories 
		err := c.BindJSON(&reqCategory)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", reqCategory)

		reqCategory.Id =primitive.NewObjectID()
		reqCategory.Created_at =time.Now().UTC()
		count,err := mongodatabase.CountCategory(reqCategory.Name)

		// SEARCH EMAIL
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			logger.Error().Msg("❌🔥 Error message :" + "already registered user")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
			return
		}

		res,err:=mongodatabase.InsertCategory(reqCategory)
		if err == nil {
			logger.Info().Msg("📢 Info message :" + "successfully registered user")
		}
		_ = res
        fmt.Println("🚀 ~ file: category.go ~ line 46 ~ returnfunc ~ res : ", res)
		c.JSON(http.StatusAccepted, res)
	}
}


func GetAllCategories() gin.HandlerFunc{
	return func(c *gin.Context){
		categories,err := mongodatabase.GetAllCategories()
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}
func GetCategory() gin.HandlerFunc{
	
	return func(c *gin.Context){
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		category,err := mongodatabase.GetCategory(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,category)
	}
}

func UpdateCategory() gin.HandlerFunc{
	return func(c *gin.Context){
		var reqCategory mongodatabase.Catagories 
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := c.BindJSON(&reqCategory)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqCategory.Created_at =time.Now().UTC()

		category,err := mongodatabase.UpdateCategory(objectID,reqCategory)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,category)
	}
}


func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context){
		
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		res,err := mongodatabase.DeleteCategory(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,res)
	}
}