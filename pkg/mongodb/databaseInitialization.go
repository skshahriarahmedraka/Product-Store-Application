package mongodatabase

import (
	// "github.com/rs/zerolog"
	// "github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/handler"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)


type DatabaseCollections struct {
	MongoDB *mongo.Database
	ProductCol *mongo.Collection
	CatagoryCol *mongo.Collection
	BrandCol *mongo.Collection
	SupplierCol *mongo.Collection
	Product_StockCol *mongo.Collection
}

var MongoCollection DatabaseCollections 

func DatabaseInitialization() {
	MongoCollection.MongoDB = MongodbConnection()
	MongoCollection.ProductCol  =  MongoCollection.MongoDB.Collection(os.Getenv("PRODUCT_COL"))
	MongoCollection.BrandCol =  MongoCollection.MongoDB.Collection(os.Getenv("BRAND_COL"))
	MongoCollection.CatagoryCol  =  MongoCollection.MongoDB.Collection(os.Getenv("CATEGORY_COL"))
	MongoCollection.SupplierCol  =  MongoCollection.MongoDB.Collection(os.Getenv("SUPPLIER_COL"))
	MongoCollection.Product_StockCol  =  MongoCollection.MongoDB.Collection(os.Getenv("PRODUCT_STOCK_COL"))
	
		
}
