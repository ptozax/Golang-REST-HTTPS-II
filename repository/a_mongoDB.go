package repository

import (
	"https/driver"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  GET   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func GET_Mongo_All(collection_name string) (*mongo.Cursor, error) {
	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	cur, err := collection.Find(*ctx, bson.M{})
	if err != nil {
		client.Disconnect(*ctx)
		return cur, err
	}
	client.Disconnect(*ctx)
	return cur, err
}

func GET_Mongo_FindOne_ByID(collection_name string, id primitive.ObjectID, data interface{}) error {
	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	err := collection.FindOne(*ctx, bson.D{{"_id", id}}).Decode(&data)
	if err != nil {
		client.Disconnect(*ctx)
		return err
	}
	client.Disconnect(*ctx)
	return nil

}

func GET_Mongo_FindOne_withFilter(collection_name string, filter bson.M, data interface{}) error {
	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	err := collection.FindOne(*ctx, filter).Decode(data)

	if err != nil {
		client.Disconnect(*ctx)
		return err
	}
	return nil
}

func GET_Mongo_Many(collection_name string, filter bson.M) (*mongo.Cursor, error) {
	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	cur, err := collection.Find(*ctx, filter)
	if err != nil {
		client.Disconnect(*ctx)
		return cur, err
	}
	return cur, nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  POST   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func POST_Mongo_One(collection_name string, data interface{}) error {
	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	_, err := collection.InsertOne(*ctx, data)
	if err != nil {
		client.Disconnect(*ctx)
		return err
	}
	client.Disconnect(*ctx)
	return nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  PUT   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func PUT_Mongo_One_ByID(id primitive.ObjectID, collection_name string, data interface{}) error {

	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	_, err := collection.UpdateOne(*ctx, bson.D{{"_id", id}}, bson.M{"$set": data})
	if err != nil {
		client.Disconnect(*ctx)
		return err
	}

	client.Disconnect(*ctx)
	return nil
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX  DELETE   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func DELETE_Mongo_One_ByID(id primitive.ObjectID, collection_name string) error {

	client, collection, ctx := driver.GET_Mongo_Collection(collection_name)
	_, err := collection.DeleteOne(*ctx, bson.M{"_id": id})
	if err != nil {
		client.Disconnect(*ctx)
		return err
	}
	client.Disconnect(*ctx)
	return nil
}
