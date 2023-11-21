package mongodatabase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

func CountProduct(name string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.ProductCol.CountDocuments(ctx, bson.M{"name": name})
	return count, err
}

func InsertProduct(product Products) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.ProductCol.InsertOne(ctx, product)
	return res, err
}

func GetAllProduct() ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var products []Products
	cursor, err := MongoCollection.ProductCol.Find(ctx, bson.M{"status_id": true})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())

		return nil, err

	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Users not Found"})
		return nil, err
	}
	return products, nil
}

func GetProduct(id primitive.ObjectID) (Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var product Products

	err := MongoCollection.ProductCol.FindOne(ctx, bson.M{"_id": id, "status_id": true}).Decode(&product)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
		return product, err
	}
	return product, nil
}

func UpdateProduct(id primitive.ObjectID, product Products) (*mongo.UpdateResult, error) {
	prev, err := GetBrand(id)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil, err
	}
	product.Id = prev.Id
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.ProductCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
	return res, err
}

func DeleteProduct(id primitive.ObjectID) (*mongo.DeleteResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.ProductCol.DeleteOne(ctx, bson.M{"_id": id})
	return res, err
}

func GetProductByName(name string) (Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var product Products
	err := MongoCollection.ProductCol.FindOne(ctx, bson.M{"name": name, "status_id": true}).Decode(&product)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
		return product, err
	}
	return product, nil

}


func GetProductByPrice(price int) (Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var product Products
	err := MongoCollection.ProductCol.FindOne(ctx, bson.M{"Unit_price": price, "status_id": true}).Decode(&product)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
		return product, err
	}
	return product, nil

}
