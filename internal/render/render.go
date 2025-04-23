package render

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
// 	if _, err := os.Stat(tmpl); os.IsNotExist(err) {
// 		http.Error(w, "Template does not exist: "+filepath.Base(tmpl), http.StatusNotFound)
// 		return err
// 	}

// 	t, err := template.ParseFiles(tmpl)
// 	if err != nil {
// 		http.Error(w, "Failed to parse template: "+err.Error(), http.StatusInternalServerError)
// 		return err
// 	}

// 	if err := t.Execute(w, data); err != nil {
// 		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
// 		return err
// 	}

// 	return nil
// }

func RenderTemplate(w http.ResponseWriter, data interface{}, templates ...string) error {
	if len(templates) == 0 {
		http.Error(w, "No template provided", http.StatusBadRequest)
		return fmt.Errorf("no template provided")
	}

	for _, tmpl := range templates {
		if _, err := os.Stat(tmpl); os.IsNotExist(err) {
			http.Error(w, "Template does not exist: "+filepath.Base(tmpl), http.StatusNotFound)
			return err
		}
	}

	t, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, "Failed to parse templates: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func RenderJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
