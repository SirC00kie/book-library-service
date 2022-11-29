package db

import (
	"book-library-service/internal/book-library-service/apperror"
	"book-library-service/internal/book-library-service/book"
	"book-library-service/pkg/logging"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, book book.Book) (string, error) {
	d.logger.Debug("create book")
	result, err := d.collection.InsertOne(ctx, book)
	if err != nil {
		return "", fmt.Errorf("failed to create new book: %v", err)
	}

	d.logger.Debug("convert InsertedID to ObjectId")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(book)

	return "", fmt.Errorf("failed to convert book to hex. probably oid: %s", oid)
}

func (d *db) FindAll(ctx context.Context) (books []book.Book, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if err != nil {
		return books, fmt.Errorf("failed to find all book. error: %v", err)
	}

	if err = cursor.All(ctx, &books); err != nil {
		return books, fmt.Errorf("failed to read all documents from cursor. error: %v", err)
	}

	return books, nil
}

func (d *db) FindOne(ctx context.Context, id string) (book book.Book, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, fmt.Errorf("failed to convert hex to objectId: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return book, apperror.ErrNotFound
		}
		return book, fmt.Errorf("failed to find one book by id: %s due to error %v", id, result.Err())
	}

	if err = result.Decode(&book); err != nil {
		return book, fmt.Errorf("failed to decode one book ( id: %s) from DB due to error %v", id, err)
	}

	return book, nil
}

func (d *db) Update(ctx context.Context, book book.Book) error {
	objectID, err := primitive.ObjectIDFromHex(book.ID)
	if err != nil {
		return fmt.Errorf("failed to convert book id to objectId: ID = %s", book.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(book)
	if err != nil {
		return fmt.Errorf("failed to marshal book: %v", err)
	}

	var updateBookObj bson.M

	err = bson.Unmarshal(userBytes, &updateBookObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal book bytes. error: %v", err)
	}

	delete(updateBookObj, "_id")

	update := bson.M{
		"$set": updateBookObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update book query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return apperror.ErrNotFound
	}

	d.logger.Tracef("Mathced %d documents and modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert book id to objectId: ID = %s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to to delete book: error: %v", err)
	}
	if result.DeletedCount == 0 {
		return apperror.ErrNotFound
	}

	d.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil

}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) book.Storage {

	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}

}
