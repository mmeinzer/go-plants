package main

import (
	"html/template"
	"path/filepath"

	"mattmeinzer.com/plants/pkg/models"
)

type templateData struct {
	Plant  *models.Plant
	Plants []*models.Plant
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		name := filepath.Base(page)
		cache[name] = ts
	}

	return cache, nil
}
