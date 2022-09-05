package mongodb

import (
	"context"
	"documentService/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (ms *MongoService) ShowAll() (*[]model.Document, error) {
	cursor, err := ms.DocumentCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var entities []model.Document
	var entity model.Document
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return &entities, nil

}
