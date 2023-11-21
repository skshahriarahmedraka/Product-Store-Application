package mongodatabase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CountBrand(name string) (int64, error ){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.BrandCol.CountDocuments(ctx, bson.M{"name": name})
	return count,err
}


func InsertBrand(brand Brands)( *mongo.InsertOneResult,error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.BrandCol.InsertOne(ctx, brand)
	return res, err
}

func GetAllBrand() ([]Brands, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var brands []Brands
	cursor, err := MongoCollection.BrandCol.Find(ctx, bson.M{"status_id":true})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	
		return nil, err
	
	}
	if err = cursor.All(context.TODO(), &brands); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Users not Found"})
		return nil, err
	}
	return brands, nil
}

func GetBrand(id primitive.ObjectID) (Brands, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var brand  Brands
		
		err := MongoCollection.BrandCol.FindOne(ctx, bson.M{"_id": id,"status_id":true}).Decode(&brand)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			// c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
			return brand, err
		}
		return brand, nil		
}

func UpdateBrand(id primitive.ObjectID,brand Brands) (*mongo.UpdateResult,error){
	prev,err:= GetBrand(id)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil,err
	}
	brand.Id = prev.Id
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.BrandCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": brand})
	return res, err
}

func DeleteBrand(id primitive.ObjectID) (*mongo.DeleteResult,error){
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.BrandCol.DeleteOne(ctx, bson.M{"_id": id})
	return res, err
}