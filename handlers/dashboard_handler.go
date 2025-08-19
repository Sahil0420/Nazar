package handlers

import (
	"log"
	"nazar/db"
	"net/http"
	"strconv"
)

type PaginationData struct {
	HasPrev     bool
	PrevPage    int
	HasNext     bool
	NextPage    int
	CurrentPage int
	TotalPages  int
}

const PageSize = 10

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("admin-session")
	if err == http.ErrNoCookie || cookie.Value != "authenticated" {
		http.Redirect(w, r, "/admin/login/", http.StatusSeeOther)
		return
	}

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	searchTerm := r.URL.Query().Get("q")

	articles, totalPages, err := db.GetAllArticlesPaginated(searchTerm, page, PageSize)
	if err != nil {
		log.Println("Error fetching paginated articles:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	categories, err := db.GetAllCategories()
	if err != nil {
		log.Println("Error fetching categories for dashboard:", err)
	}

	data := map[string]interface{}{
		"Title":      "Admin Dashboard",
		"Articles":   articles,
		"Categories": categories,
		"SearchTerm": searchTerm,
		"Pagination": PaginationData{
			HasPrev:     page > 1,
			PrevPage:    page - 1,
			HasNext:     page < totalPages,
			NextPage:    page + 1,
			CurrentPage: page,
			TotalPages:  totalPages,
		},
	}

	status := r.URL.Query().Get("status")
	switch status {
	case "add_success":
		data["SuccessMessage"] = "Article added successfully!"
	case "delete_success":
		data["SuccessMessage"] = "Article deleted successfully!"
	case "category_success":
		data["CategorySuccessMessage"] = "New Categories Added Successfully"
	case "title_already_exists":
		data["ErrorMessage"] = "An article with this title already exists!"
	}

	err = templates.ExecuteTemplate(w, "dashboard.html", data)
	if err != nil {
		log.Println("Dashboard Template Execution Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
