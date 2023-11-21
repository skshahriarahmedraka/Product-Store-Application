package mongodatabase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func CountSupplier(email string) (int64, error ){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.SupplierCol.CountDocuments(ctx, bson.M{"email": email})
	return count,err
}

func CountSupplierById(id primitive.ObjectID) (int64, error ){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.SupplierCol.CountDocuments(ctx, bson.M{"_id": id})
	return count,err
}


func InsertSupplier(supplier Suppliers)( *mongo.InsertOneResult,error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.SupplierCol.InsertOne(ctx, supplier)
	return res, err
}

func GetAllSupplier() ([]Suppliers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var supplier []Suppliers
	cursor, err := MongoCollection.SupplierCol.Find(ctx, bson.M{"status_id":true})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
	
		return nil, err
	
	}
	if err = cursor.All(context.TODO(), &supplier); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Users not Found"})
		return nil, err
	}
	return supplier, nil
}

func GetSupplier(id primitive.ObjectID) (Suppliers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var supplier  Suppliers
		
		err := MongoCollection.SupplierCol.FindOne(ctx, bson.M{"_id": id,"status_id":true}).Decode(&supplier)
		if err != nil {
			logger.Error().Msg("❌🔥 Error message :" + err.Error())
			// c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
			return supplier, err
		}
		return supplier, nil		
}

func UpdateSupplier(id primitive.ObjectID,supplier Suppliers) (*mongo.UpdateResult,error){
	prev,err:= GetSupplier(id)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil,err
	}
	supplier.Id = prev.Id
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.SupplierCol.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": supplier})
	return res, err
}

func DeleteSupplier(id primitive.ObjectID) (*mongo.DeleteResult,error){
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	res, err := MongoCollection.SupplierCol.DeleteOne(ctx, bson.M{"_id": id})
	return res, err
}