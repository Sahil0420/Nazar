package db

import (
	"context"
	"nazar/models"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options" // Corrected import path
)

func CreateCategories(categories []models.Category) error {
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("categories")

	for _, cat := range categories {
		catName := strings.TrimSpace(cat.Name)
		if catName == "" {
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		filter := bson.M{"name": bson.M{"$regex": "^" + catName + "$", "$options": "i"}}
		update := bson.M{"$setOnInsert": bson.M{"name": catName, "slug": cat.Slug}}
		opts := options.UpdateOne().SetUpsert(true)

		if _, err := coll.UpdateOne(ctx, filter, update, opts); err != nil {
			return err
		}
	}
	return nil
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()

	cursor, err := coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// The .All() method is where the decoding happens. The error points here.
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategoryNameByID(id bson.ObjectID) (string, error) {
	var category models.Category
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	err := coll.FindOne(ctx, filter).Decode(&category)

	if err != nil {
		return "not-found", err
	}

	return category.Name, nil
}

func GetCategoryBySlug(slug string) (models.Category, error) {
	var category models.Category
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("categories")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, bson.M{"slug": slug}).Decode(&category)
	return category, err
}
