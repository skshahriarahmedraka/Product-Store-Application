package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CreateBrand() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqBrand mongodatabase.Brands
		err := c.BindJSON(&reqBrand)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", reqBrand)

		reqBrand.Id = primitive.NewObjectID()
		reqBrand.Created_at = time.Now().UTC()
		count, err := mongodatabase.CountBrand(reqBrand.Name)

		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			logger.Error().Msg("❌🔥 Error message :" + "already registered Brand")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered brand"})
			return
		}

		res, err := mongodatabase.InsertBrand(reqBrand)
		if err == nil {
			logger.Info().Msg("📢 Info message :" + "successfully registered Brand")
		}
		_ = res
		fmt.Println("🚀 ~ file: category.go ~ line 46 ~ returnfunc ~ res : ", res)
		c.JSON(http.StatusAccepted, res)
	}
}

func GetAllBrands() gin.HandlerFunc {
	return func(c *gin.Context) {
		brands, err := mongodatabase.GetAllBrand()
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, brands)
	}
}
func GetBrand() gin.HandlerFunc {

	return func(c *gin.Context) {
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		brand, err := mongodatabase.GetBrand(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, brand)
	}
}

func UpdateBrand() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBrand mongodatabase.Brands
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := c.BindJSON(&reqBrand)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqBrand.Created_at = time.Now().UTC()

		brand, err := mongodatabase.UpdateBrand(objectID, reqBrand)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, brand)
	}
}

func DeleteBrand() gin.HandlerFunc {
	return func(c *gin.Context) {

		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		res, err := mongodatabase.DeleteBrand(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
