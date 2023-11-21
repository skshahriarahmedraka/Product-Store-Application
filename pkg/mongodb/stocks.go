package mongodatabase

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CountProductStockById(id primitive.ObjectID) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var stock Product_Stocks
	MongoCollection.Product_StockCol.FindOne(ctx, bson.M{"product_id": id}).Decode(&stock)
	return stock.Stock_quantity
}

func AddProductInStock(id primitive.ObjectID) (*mongo.InsertOneResult, error) {
	count := CountProductStockById(id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	if count == 0 {
		var stock Product_Stocks
		stock.Id = primitive.NewObjectID()
		stock.Product_id = id
		stock.Stock_quantity = 1
		stock.Updated_at = time.Now().UTC()
		res, err := MongoCollection.Product_StockCol.InsertOne(ctx, stock)
		return res, err
	} else {
		_, err := MongoCollection.Product_StockCol.UpdateOne(ctx, bson.M{"product_id": id}, bson.M{"$inc": bson.M{"stock_quantity": 1}})
		return nil, err
	}
}

func DeleteProductStock(id string) (*mongo.UpdateResult, error) {
	hexid, _ := primitive.ObjectIDFromHex(id)
	count := CountProductStockById(hexid)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	if count > 0 {

		res, err := MongoCollection.Product_StockCol.UpdateOne(ctx, bson.M{"product_id": hexid}, bson.M{"$inc": bson.M{"stock_quantity": -1}})
		return res, err
	}
	return nil, fmt.Errorf("stock not found")
}

func GetAllStock() ([]Brands, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var brands []Brands
	cursor, err := MongoCollection.BrandCol.Find(ctx, bson.M{"status_id": true})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())

		return nil, err

	}
	if err = cursor.All(context.TODO(), &brands); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil, err
	}
	return brands, nil
}

func GetStock(id primitive.ObjectID) (Brands, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var brand Brands

	err := MongoCollection.BrandCol.FindOne(ctx, bson.M{"_id": id, "status_id": true}).Decode(&brand)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return brand, err
	}
	return brand, nil
}
