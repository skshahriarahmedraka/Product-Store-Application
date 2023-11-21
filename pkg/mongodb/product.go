package mongodatabase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func CountProduct(name string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.ProductCol.CountDocuments(ctx, bson.M{"name": name})
	return count, err
}

func DublicateBySupplier(name string, supplierId primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := MongoCollection.ProductCol.CountDocuments(ctx, bson.M{"name": name, "supplier_id": supplierId})
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
	res, err := DeleteProductStock(id.Hex())
	_ = res
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil, err

	}
	delRes, err := MongoCollection.ProductCol.DeleteOne(ctx, bson.M{"_id": id})
	return delRes, err
}

func GetProductByName(name string) (Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var product Products
	err := MongoCollection.ProductCol.FindOne(ctx, bson.M{"name": name, "status_id": true}).Decode(&product)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return product, err
	}
	return product, nil

}

func GetProductByPrice(price int) (Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var product Products
	err := MongoCollection.ProductCol.FindOne(ctx, bson.M{"unit_price": price, "status_id": true}).Decode(&product)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return product, err
	}
	return product, nil

}

func GetProductByMulBrand(brandlist Brand_List) ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var productlist []Products
	filter := bson.M{"brand_id": bson.M{"$in": brandlist.Brand_id}}
	cursor, err := MongoCollection.ProductCol.Find(ctx, filter)
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return productlist, err
	}
	defer cursor.Close(ctx)

	// var products []Products
	if err = cursor.All(ctx, &productlist); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return productlist, err
	}
	return productlist, nil
}

func GetProductByCategory(id primitive.ObjectID) ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var productlist []Products
	cursor, err := MongoCollection.ProductCol.Find(ctx, bson.M{"catagory_id": id})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())

		return nil, err

	}
	if err = cursor.All(context.TODO(), &productlist); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil, err
	}
	return productlist, nil
}

func GetProductBySupplier(id primitive.ObjectID) ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var productlist []Products
	cursor, err := MongoCollection.ProductCol.Find(ctx, bson.M{"supplier_id": id})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())

		return nil, err

	}
	if err = cursor.All(context.TODO(), &productlist); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil, err
	}
	return productlist, nil
}

func GetProductByVerifiedSupplier(id primitive.ObjectID) ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	var productlist []Products
	cursor, err := MongoCollection.ProductCol.Find(ctx, bson.M{"supplier_id": id, "status_id": true})
	if err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())

		return nil, err

	}
	if err = cursor.All(context.TODO(), &productlist); err != nil {
		logger.Error().Msg("❌🔥 Error message :" + err.Error())
		return nil, err
	}

	return productlist, nil
}
