package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqSupplier mongodatabase.Suppliers
		err := c.BindJSON(&reqSupplier)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", reqSupplier)

		reqSupplier.Id = primitive.NewObjectID()
		reqSupplier.Created_at = time.Now().UTC()
		count, err := mongodatabase.CountSupplier(reqSupplier.Email)

		// SEARCH EMAIL
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			logger.Error().Msg("❌🔥 Error message :" + "already registered supplier")
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered supplier"})
			return
		}

		res, err := mongodatabase.InsertSupplier(reqSupplier)
		if err == nil {
			logger.Info().Msg("📢 Info message :" + "successfully registered supplier")
		}
		_ = res
		fmt.Println("🚀 ~ file: category.go ~ line 46 ~ returnfunc ~ res : ", res)
		c.JSON(http.StatusAccepted, res)
	}
}

func GetAllSuppliers() gin.HandlerFunc {
	return func(c *gin.Context) {
		suppliers, err := mongodatabase.GetAllSupplier()
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, suppliers)
	}
}
func GetSupplier() gin.HandlerFunc {

	return func(c *gin.Context) {
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		supplier, err := mongodatabase.GetSupplier(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, supplier)
	}
}

func UpdateSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqSupplier mongodatabase.Suppliers
		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		err := c.BindJSON(&reqSupplier)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqSupplier.Created_at = time.Now().UTC()

		supplier, err := mongodatabase.UpdateSupplier(objectID, reqSupplier)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, supplier)
	}
}

func DeleteSupplier() gin.HandlerFunc {
	return func(c *gin.Context) {

		objectID, _ := primitive.ObjectIDFromHex(c.Param("id"))
		res, err := mongodatabase.DeleteSupplier(objectID)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
