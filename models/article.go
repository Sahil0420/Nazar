package models

import (
	"html/template"

	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Article struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title           string        `bson:"title" json:"title"`
	Slug            string        `bson:"slug"`
	Author          string        `bson:"author" json:"author"`
	Category        bson.ObjectID `bson:"category_id" json:"category_id"`
	CategorySlug    string        `bson:"-" json:"-"`
	ImageUrl        string        `bson:"image_url" json:"image_url"`
	Description     string        `bson:"description" json:"-"`
	DescriptionHTML template.HTML `bson:"-" json:"-"`
	CreatedAt       time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time     `bson:"updated_at" json:"updated_at"`
}
