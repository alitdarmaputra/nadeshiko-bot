package repositories

import (
	"context"

	"github.com/alitdarmaputra/nadeshiko-bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "app_instagram"

// InstagramRepo implement models.InstagramRepository
type InstagramRepo struct {
	db  *mongo.Database
	ctx context.Context
}

func NewInstagramRepo(db *mongo.Database, ctx context.Context) *InstagramRepo {
	return &InstagramRepo{db: db, ctx: ctx}
}

func (i *InstagramRepo) FindOne(username string) (*models.Instagram, error) {
	var filter = bson.M{"username": username}
	var result models.Instagram
	err := i.db.Collection(COLLECTION_NAME).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (i *InstagramRepo) Save(instagram *models.Instagram) error {
	_, err := i.db.Collection(COLLECTION_NAME).InsertOne(context.TODO(), *instagram)
	if err != nil {
		return err
	}
	return nil
}
