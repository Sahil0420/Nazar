package utils

import (
	"regexp"
	"strings"
)

func Slugify(title string) string {
	// Lowercase
	slug := strings.ToLower(title)
	// Replace non-alphanumeric with hyphens
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	// Remove leading/trailing hyphens
	slug = strings.Trim(slug, "-")
	return slug
}
