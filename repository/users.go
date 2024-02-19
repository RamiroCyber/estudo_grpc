package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

const UserCollection = "users"

type UserRepository interface{}

type userRepository struct {
	UserCollection *mongo.Collection
}

func NewUsersRepository(coon mongo.Client) UserRepository {
	return &userRepository{coon.Database(os.Getenv("MONGO_DB")).Collection(UserCollection)}
}
