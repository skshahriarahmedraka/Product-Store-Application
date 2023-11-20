package mongodatabase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CountCategory(name string) (int64, error ){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.CatagoryCol.CountDocuments(ctx, bson.M{"name": name})
	return count,err
}


func InsertCategory(category Catagories)( *mongo.InsertOneResult,error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.CatagoryCol.InsertOne(ctx, category)
	return res, err
}

func GetAllCategories() ([]Catagories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var categories []Catagories
	cursor, err := MongoCollection.CatagoryCol.Find(ctx, bson.M{"status_id":true})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	
		return nil, err
	
	}
	if err = cursor.All(context.TODO(), &categories); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Users not Found"})
		return nil, err
	}
	return categories, nil
}

func GetCategory(id primitive.ObjectID) (Catagories, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var category Catagories
		
		err := MongoCollection.CatagoryCol.FindOne(ctx, bson.M{"_id": id,"status_id":true}).Decode(&category)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			// c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
			return category, err
		}
		return category, nil
		
}

func UpdateCategory(id primitive.ObjectID,category Catagories) (*mongo.UpdateResult,error){
	prev,err:= GetCategory(id)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil,err
	}
	category.Id = prev.Id
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.CatagoryCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": category})
	return res, err
}

func DeleteCategory(id primitive.ObjectID) (*mongo.DeleteResult,error){
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.CatagoryCol.DeleteOne(ctx, bson.M{"_id": id})
	return res, err
}