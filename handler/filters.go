package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)


func GetProductByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var req struct {
			Name string `json:"name"`
		}
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product, err := mongodatabase.GetProductByName(req.Name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func GetProductByPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		price := c.Param("price")
		priceInt, _ := strconv.Atoi(price)
		product, err := mongodatabase.GetProductByPrice(priceInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func GetProductByMulBrand() gin.HandlerFunc {
	return func(c *gin.Context) {
		var brandlist mongodatabase.Brand_List
		if err := c.BindJSON(&brandlist); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products, err := mongodatabase.GetProductByMulBrand(brandlist)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categoryid primitive.ObjectID
		id:= c.Param("id")
		categoryid, _ = primitive.ObjectIDFromHex(id)
		products, err := mongodatabase.GetProductByCategory(categoryid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductBySupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var supplierid primitive.ObjectID
		id:= c.Param("id")
		supplierid, _ = primitive.ObjectIDFromHex(id)
		products, err := mongodatabase.GetProductBySupplier(supplierid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductByVerifiedSupplier()gin.HandlerFunc{
	return func(c *gin.Context) {
		var supplierid primitive.ObjectID
		id:= c.Param("id")
		supplierid, _ = primitive.ObjectIDFromHex(id)
		products, err := mongodatabase.GetProductByVerifiedSupplier(supplierid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}