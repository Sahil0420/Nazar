package handlers

import (
	"crypto/subtle"
	"log"
	"net/http"
	"time"
)

// AdminLoginHandler remains the same, it's a good router.
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleAdminLoginGet(w, r)
	case http.MethodPost:
		handleAdminLoginPost(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// CORRECTED GET handler
func handleAdminLoginGet(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{"Title": "Admin Access"}

	err := templates.ExecuteTemplate(w, "admin_login.html", data)
	if err != nil {
		log.Println("Template Execution Error :", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleAdminLoginPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	input := r.FormValue("passphrase")

	if subtle.ConstantTimeCompare([]byte(input), []byte(adminPassphrase)) != 1 {
		log.Println("Failed login attempt")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "admin-session",
		Value:    "authenticated",
		Path:     "/admin/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   r.TLS != nil,
		SameSite: http.SameSiteLaxMode,
	})
	http.Redirect(w, r, "/admin/dashboard/", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "admin-session",
		Value:   "",
		Path:    "/admin/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
