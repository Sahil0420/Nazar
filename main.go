package main

import (
	"html/template"
	"log"
	"nazar/db"
	"nazar/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// --- Configuration ---
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	adminPass := os.Getenv("ADMIN_PASSPHRASE")
	if adminPass == "" {
		log.Fatal("ADMIN_PASSPHRASE not set")
	}

	// --- Database ---
	db.ConnectMongo()

	// function  that you wanna use in html
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"len": func(s []string) int { return len(s) },
	}
	// --- Template Parsing ---
	templates := template.Must(
		template.New("").Funcs(funcMap).ParseFiles(
			// Layouts and Components
			"templates/base.html",
			"templates/components/header.html",
			"templates/components/add_article_form.html",
			"templates/components/add_category_form.html",
			"templates/index.html",
			"templates/article.html",
			"templates/admin_login.html",
			"templates/dashboard.html",
		),
	)

	// --- Dependency Injection ---
	// Give the handlers access to the parsed templates and the passphrase.
	handlers.InitTemplates(templates)
	handlers.SetAdminPassphrase(adminPass)

	// --- Server Setup ---
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/article/", handlers.ArticleHandler) // Added trailing slash for consistency
	http.HandleFunc("/admin/login/", handlers.AdminLoginHandler)
	http.HandleFunc("/admin/dashboard/", handlers.AdminDashboardHandler)
	http.HandleFunc("/admin/logout/", handlers.LogoutHandler)
	http.HandleFunc("/admin/add-category/", handlers.AddCategoryHandler)
	http.HandleFunc("/admin/add-article/", handlers.AddArticleHandler)
	http.HandleFunc("/admin/delete-article/", handlers.DeleteArticleHandler)
	http.HandleFunc("/category/", handlers.CategoryHandler)

	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// The "How to Train Your Dragon" franchise tells the story of Hiccup, a young Viking living in a world where dragons are feared and hunted. Hiccup, however, finds himself questioning his tribe's traditions when he encounters a rare and injured dragon, a Night Fury, who he eventually names Toothless. Their unlikely bond becomes a catalyst for change, opening Hiccup's eyes and challenging the long-held beliefs of his people.
//https://m.media-amazon.com/images/M/MV5BMjA5NDQyMjc2NF5BMl5BanBnXkFtZTcwMjg5ODcyMw@@._V1_FMjpg_UX1000_.jpg
