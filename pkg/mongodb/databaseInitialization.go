package mongodatabase

import (
	// "github.com/rs/zerolog"
	// "github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/handler"
	"go.mongodb.org/mongo-driver/mongo"
)


var MongoCollection *mongo.Database 
func DatabaseInitialization() {
	MongoCollection = MongodbConnection()
}
