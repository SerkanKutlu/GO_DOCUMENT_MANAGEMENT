package mongodb

import (
	"context"
	customerror "documentService/customError"
	"documentService/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (ms *MongoService) GetAllPaths(id *string) (*[]model.DownloadModel, *customerror.CustomError) {
	opts := options.Find().SetProjection(bson.D{{"path", 1}, {"filename", 1}})
	var cursor *mongo.Cursor
	var err error
	if id == nil {
		cursor, err = ms.DocumentCollection.Find(context.Background(), bson.D{}, opts)
	} else {
		cursor, err = ms.DocumentCollection.Find(context.Background(), bson.M{"userid": id}, opts)
	}
	if err != nil {
		return nil, customerror.NewError(err.Error(), 500)
	}

	var entities []model.DownloadModel
	var entity model.DownloadModel
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&entity); err != nil {
			return nil, customerror.InvalidEntity
		}
		entities = append(entities, entity)
	}
	return &entities, nil
}

func (ms *MongoService) GetPathById(id *string, userId *string) (*string, *customerror.CustomError) {
	opts := options.FindOne().SetProjection(bson.D{{"path", 1}})
	var singleResult *mongo.SingleResult
	if userId == nil {
		singleResult = ms.DocumentCollection.FindOne(context.Background(), bson.M{"_id": id}, opts)
	} else {
		singleResult = ms.DocumentCollection.FindOne(context.Background(), bson.M{"_id": id, "userid": userId}, opts)
	}
	if singleResult.Err() != nil {
		return nil, customerror.NotFoundError
	}
	var entity model.Document
	if err := singleResult.Decode(&entity); err != nil {
		return nil, customerror.InvalidEntity
	}
	return &entity.Path, nil
}
