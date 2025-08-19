package handlers

import (
	"errors"
	"log"
	"nazar/db"
	"nazar/models"
	"nazar/utils"
	"net/http" // Import the strings package
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	// 1. FIXED authentication logic.
	cookie, err := r.Cookie("admin-session")
	if err == http.ErrNoCookie || cookie.Value != "authenticated" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid Form Data", http.StatusBadRequest)
		log.Fatal("Invalid Form data", err)
		return
	}

	categoriesIDStr := r.FormValue("category_id")

	if categoriesIDStr == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		log.Fatal("Invalid CategoryID", categoriesIDStr)
		return
	}

	categoryObjectID, err := bson.ObjectIDFromHex(categoriesIDStr)
	if err != nil {
		http.Error(w, "Invalid category ID format", http.StatusBadRequest)
		log.Fatal("Invalid Category ID Format")
		return
	}

	slug_title := utils.Slugify(r.FormValue("title"))

	article := models.Article{
		Title:       r.FormValue("title"),
		Slug:        slug_title,
		Description: r.FormValue("description"),
		Author:      r.FormValue("author"),
		ImageUrl:    r.FormValue("image_url"),
		Category:    categoryObjectID, // Assign the cleaned slice
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = db.CreateArticle(article)
	if err != nil {
		if errors.Is(err, db.ErrArticleExists) {
			http.Redirect(w, r, "/admin/dashboard?status=title_already_exists", http.StatusSeeOther)
			return
		}
		log.Println("Failed to create article", err)
		http.Error(w, "Failed to create article in database", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard?status=add_success", http.StatusSeeOther)
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("admin-session")
	if err == http.ErrNoCookie || cookie.Value != "authenticated" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	articleIDStr := r.FormValue("article_id")
	if articleIDStr == "" {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	if err := db.DeleteArticleByID(articleIDStr); err != nil {
		log.Printf("Failed to delete article with ID %s : %v\n", strings.TrimSpace(articleIDStr), err)
		http.Error(w, "Database error while deleting article", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard?status=delete_success", http.StatusSeeOther)
}
