package handlers

import (
	"html/template"
	"net/http"
)

var (
	templates       *template.Template
	adminPassphrase string
	blogName        string
)

func InitTemplates(t *template.Template) {
	templates = t
}

func SetAdminPassphrase(pass string) {
	adminPassphrase = pass
}

func SetBlogName(name string) {
	blogName = name
}

func Render(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	// Always inject BlogName if not already present
	if _, exists := data["BlogName"]; !exists {
		data["BlogName"] = blogName
	}

	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
