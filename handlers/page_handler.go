package handlers

import (
	"fmt"
	"log"
	"nazar/db"
	"nazar/utils"
	"net/http"
	"strconv"
	"strings"
)

const public_page_size = 15

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	searchTerm := r.URL.Query().Get("q")

	articles, totalPages, err := db.GetAllArticlesPaginated(searchTerm, page, public_page_size)
	if err != nil {
		log.Println("Error loading articles: ", err)
		http.Error(w, "Error loading articles", http.StatusInternalServerError)
		return
	}

	categories, err := db.GetAllCategories()
	if err != nil {
		log.Println("Error fetching categories: ", err)
		http.Error(w, "Error loading categories", http.StatusInternalServerError)
		return
	}

	pageTitle := "Recent Articles"
	if searchTerm != "" {
		pageTitle = "Search results for : " + searchTerm
	}

	data := map[string]interface{}{
		"Title":      "Nazar - Home",
		"PageTitle":  pageTitle,
		"Categories": categories,
		"Articles":   articles,
		"Pagination": PaginationData{
			HasPrev:     page > 1,
			PrevPage:    page - 1,
			HasNext:     page < totalPages,
			NextPage:    page + 1,
			CurrentPage: page,
			TotalPages:  totalPages,
		},
	}

	Render(w, "base.html", data)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 3 || parts[0] != "article" {
		http.NotFound(w, r)
		return
	}

	categorySlug := parts[1]
	articleSlug := parts[2]

	article, err := db.GetArticleBySlug(articleSlug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	categoryName, err := db.GetCategoryNameByID(article.Category)
	if err != nil {
		http.Error(w, "Category not Found", http.StatusInternalServerError)
		return
	}

	if utils.Slugify(categoryName) != categorySlug {
		correctURL := fmt.Sprintf("/article/%s/%s", utils.Slugify(categoryName), article.Slug)
		http.Redirect(w, r, correctURL, http.StatusMovedPermanently)
		return
	}

	article.DescriptionHTML = utils.RenderMarkdown(article.Description)

	err = templates.ExecuteTemplate(w, "article.html", article)
	if err != nil {
		log.Println("Article Template Error: ", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
