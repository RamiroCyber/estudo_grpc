package database

import "go.mongodb.org/mongo-driver/mongo"

var UserCollection *mongo.Collection

func SetupCollections() {
	UserCollection = Client.Database("mongogrpc").Collection("user")
}
