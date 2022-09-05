package mongodb

import (
	"context"
	customerror "documentService/customError"
	"documentService/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (ms *MongoService) GetAll() (*[]model.Document, *customerror.CustomError) {
	cursor, err := ms.DocumentCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, customerror.NewError(err.Error(), 500)
	}
	var entities []model.Document
	var entity model.Document
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&entity); err != nil {
			return nil, customerror.InvalidEntity
		}
		entities = append(entities, entity)
	}
	return &entities, nil
}
func (ms *MongoService) GetOfUsers(userId string) (*[]model.Document, *customerror.CustomError) {
	cursor, err := ms.DocumentCollection.Find(context.Background(), bson.M{"userid": userId})
	if err != nil {
		return nil, customerror.NewError(err.Error(), 500)
	}
	var entities []model.Document
	var entity model.Document
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&entity); err != nil {
			return nil, customerror.InvalidEntity
		}
		entities = append(entities, entity)
	}
	return &entities, nil
}
func (ms *MongoService) Insert(document *model.Document) *customerror.CustomError {
	_, err := ms.DocumentCollection.InsertOne(context.Background(), document)
	if err != nil {
		return customerror.NewError(err.Error(), 500)
	}
	return nil

}
func (ms *MongoService) InsertMany(documents *[]model.Document) *customerror.CustomError {

	list := make([]interface{}, len(*documents))

	for i := range *documents {
		list[i] = (*documents)[i]
	}
	_, err := ms.DocumentCollection.InsertMany(context.Background(), list)
	if err != nil {
		return customerror.NewError(err.Error(), 500)
	}
	return nil

}
func (ms *MongoService) Delete(id string) (*string, *customerror.CustomError) {
	foundEntity := ms.DocumentCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	if foundEntity.Err() != nil {
		return nil, customerror.NotFoundError
	}
	var document model.Document
	err := foundEntity.Decode(document)
	if err != nil {
		return nil, customerror.NewError("Found a invalid entity at the database. Decode Error.", 500)
	}
	return &document.Path, nil
}
func (ms *MongoService) DeleteWithUserId(id string, userId string) (*string, *customerror.CustomError) {
	foundEntity := ms.DocumentCollection.FindOneAndDelete(context.TODO(), bson.M{"id": id})
	//findoneanddelete method can not be working
	if foundEntity.Err() != nil {
		return nil, customerror.NotFoundError
	}
	var document model.Document
	err := foundEntity.Decode(document)
	if err != nil {
		return nil, customerror.NewError("Found a invalid entity at the database. Decode Error.", 500)
	}
	return &document.Path, nil

}
