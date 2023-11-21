package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
)


func GetProductByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		product, err := mongodatabase.GetProductByName(name)
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
		var product mongodatabase.Products
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products, err := product.GetProductByMulBrand(product.Brand)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product mongodatabase.Products
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products, err := product.GetProductByCategory(product.Category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductBySupplier() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product mongodatabase.Products
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products, err := product.GetProductBySupplier(product.Supplier)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetProductByVerifiedSupplier()gin.HandlerFunc{
	return func(c *gin.Context) {
		var product mongodatabase.Products
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products, err := product.GetProductByVerifiedSupplier(product.Supplier)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}