package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/skshahriarahmedraka/Product-Store-Application/pkg/mongodb"
)



func CreateProduct()gin.HandlerFunc{
	return func (c *gin.Context){
		var reqProduct mongodatabase.Products

		err := c.BindJSON(&reqProduct)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", reqProduct)
		c.JSON(http.StatusOK, gin.H{"message": "Product Created Successfully"})
	}
}