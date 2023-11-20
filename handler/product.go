package handler

import (
	// "context"
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
		
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		// defer cancel()

		// // SEARCH EMAIL
		// count, err := H.Mongo.Collection(os.Getenv("PRODUCT_COL")).CountDocuments(ctx, bson.M{"Email": user.Email})
		// if err != nil {
		// 	logger.Error().Msg("❌🔥 Error message :" + err.Error())

		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }
		// if count > 0 {
		// 	logger.Error().Msg("❌🔥 Error message :" + "already registered user")
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
		// 	return
		// }

		c.JSON(http.StatusOK, gin.H{"message": "Product Created Successfully"})



	}
}