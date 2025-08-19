package db

import (
	"context"
	"errors"
	"log"
	"math"
	"nazar/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetAllArticlesPaginated(searchTerm string, page int, pageSize int) ([]models.Article, int, error) {

	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, relying on system environment variables")
		}
	}

	var articles []models.Article
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}

	if searchTerm != "" {
		filter = bson.M{
			// Search in both title and content for the searchTerm
			"$or": []bson.M{
				{"title": bson.M{"$regex": searchTerm, "$options": "i"}},
				{"content": bson.M{"$regex": searchTerm, "$options": "i"}},
			},
		}
	}

	// First, get the total count of documents matching the filter
	totalCount, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	if totalCount == 0 {
		return articles, 0, nil
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	skip := (page - 1) * pageSize
	findOptions.SetSkip(int64(skip))
	findOptions.SetSort(bson.M{"created_at": -1})

	cursor, err := coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &articles); err != nil {
		return nil, 0, err
	}

	catColl := Client.Database(os.Getenv("DB_NAME")).Collection("categories")

	for i, article := range articles {
		var cat struct {
			Slug string `bson:"slug"`
		}
		err := catColl.FindOne(ctx, bson.M{"_id": article.Category}).Decode(&cat)
		if err != nil {
			log.Println("Warning : Failed to fetch the category's slug for article", article.Title, err)
			continue
		}

		articles[i].CategorySlug = cat.Slug
	}

	return articles, totalPages, nil
}

func GetArticleByID(id string) (*models.Article, error) {
	var article models.Article

	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid article ID Format")
	}

	filter := bson.M{"_id": objID}

	err = coll.FindOne(ctx, filter).Decode(&article)
	if err != nil {
		return nil, err
	}

	return &article, nil

}

func GetArticleBySlug(slug string) (*models.Article, error) {
	var article models.Article

	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := coll.FindOne(ctx, bson.M{"slug": slug}).Decode(&article)

	if err != nil {
		return nil, err
	}

	return &article, nil
}

func GetArticleByCategoryAndSlug(category string, slug string) (*models.Article, error) {
	var article models.Article

	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{"category": category, "slug": slug}
	err := coll.FindOne(ctx, filter).Decode(&article)

	if err != nil {
		return nil, err
	}

	return &article, nil

}

var ErrArticleExists = errors.New("article with this slug already exists")

func CreateArticle(article models.Article) error {
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if an article with the same slug already exists
	count, err := coll.CountDocuments(ctx, bson.M{"slug": article.Slug})
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrArticleExists
	}

	_, err = coll.InsertOne(ctx, article)
	return err
}

func DeleteArticleByID(id string) error {
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid article ID format")
	}

	log.Printf("ObjectID converted in db func %s", objID)

	log.Printf("Database name %s Coll Name %s", os.Getenv("DB_NAME"), coll.Name())

	filter := bson.M{"_id": objID}

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		// In v2, there is no mongo.ErrNoDocuments for deletes
		log.Printf("No article found with ID %s", id)
		return nil
	}

	return nil
}

func GetArticlesByCategory(categoryID bson.ObjectID, page, pageSize int) ([]models.Article, int, error) {
	var articles []models.Article
	coll := Client.Database(os.Getenv("DB_NAME")).Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"category_id": categoryID}

	totalCount, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	if totalCount == 0 {
		return articles, 0, nil
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	findOptions := options.Find().
		SetLimit(int64(pageSize)).
		SetSkip(int64((page - 1) * pageSize)).
		SetSort(bson.M{"created_at": -1})

	cursor, err := coll.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &articles); err != nil {
		return nil, 0, err
	}

	return articles, totalPages, nil
}
