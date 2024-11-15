package handlers

import (
	"groupie-tracker/data"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

// FormatLocation formats a location string by replacing underscores and hyphens
func FormatLocation(location string) string {
	// Replace underscores with spaces
	location = strings.ReplaceAll(location, "_", " ")
	// Replace hyphens with commas and spaces
	location = strings.ReplaceAll(location, "-", ", ")
	// Capitalize first letter of each word
	words := strings.Fields(location)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		if words[i] == "Usa" {
			words[i] = "USA"
		} else if words[i] == "Uk" {
			words[i] = "UK"
		}
	}
	return strings.Join(words, " ")
}

func FormatDate(date string) string {
	// Remove asterisks (*)
	date = strings.ReplaceAll(date, "*", "")
	// Return the date as is without any other modifications
	return date
}

// Create template functions map
var funcMap = template.FuncMap{
	"FormatLocation": FormatLocation,
	"FormatDate":     FormatDate,
}

// Load the templates with function map
var templates = template.Must(template.New("").Funcs(funcMap).ParseFiles(filepath.Join("templates", "index.html")))

// IndexHandler renders the homepage with artist data
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all data
	indexes, err := data.FetchAllData()
	if err != nil {
		http.Error(w, "Failed to fetch artist data", http.StatusInternalServerError)
		return
	}

	// Render template with fetched data
	err = templates.ExecuteTemplate(w, "index.html", indexes)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/landing.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
