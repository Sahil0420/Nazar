package handlers

import (
	"log"
	"nazar/db"
	"nazar/models"
	"nazar/utils"
	"net/http"
	"strconv"
	"strings"
)

func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	categoryNames := r.Form["category_name[]"]
	if len(categoryNames) == 0 {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	var categories []models.Category

	for _, name := range categoryNames {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		categories = append(categories, models.Category{
			Name: name,
			Slug: utils.Slugify(name),
		})
	}

	if len(categories) == 0 {
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	if err := db.CreateCategories(categories); err != nil {
		log.Println("Failed to add categories:", err)
		http.Error(w, "Database error while adding categories", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard?status=category_success", http.StatusSeeOther)
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 2 || parts[0] != "category" {
		http.NotFound(w, r)
		return
	}

	categorySlug := parts[1]

	category, err := db.GetCategoryBySlug(categorySlug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	articles, totalPages, err := db.GetArticlesByCategory(category.ID, page, public_page_size)
	if err != nil {
		log.Println("Error fetching category articles:", err)
		http.Error(w, "Error loading category articles", http.StatusInternalServerError)
		return
	}

	for i := range articles {
		articles[i].CategorySlug = category.Slug
	}

	categories, err := db.GetAllCategories()
	if err != nil {
		log.Println("Error fetching categories:", err)
	}

	data := map[string]interface{}{
		"Title":      "Category - " + category.Name,
		"PageTitle":  "Articles in " + category.Name,
		"Articles":   articles,
		"Categories": categories,
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
