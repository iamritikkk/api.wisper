package controllers

import (
	"context"
	"github.com/meanii/api.wisper/clients"
	"github.com/meanii/api.wisper/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	COLLECTION = "users"
)

type UserController struct {
	model      models.User
	collection mongo.Collection
}

func (uc *UserController) LoadCollection() {
	database := clients.GetDatabase()
	database.Collection(COLLECTION)
}

func (uc *UserController) InsertOne(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return uc.collection.InsertOne(ctx, user)
}
