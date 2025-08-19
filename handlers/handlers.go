package handlers

import (
	"html/template"
)

var (
	templates       *template.Template
	adminPassphrase string
)

func InitTemplates(t *template.Template) {
	templates = t
}

func SetAdminPassphrase(pass string) {
	adminPassphrase = pass
}
