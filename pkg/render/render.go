package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mrpineapples/bookings/pkg/config"
	"github.com/mrpineapples/bookings/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Template renders templates using html/template
func Template(w http.ResponseWriter, tpl string, data *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("there was an error creating the template cache:", err)
		}
	}

	t, ok := tc[tpl]
	if !ok {
		log.Fatalf("template '%s' does not exist in template cache", tpl)
	}

	buf := new(bytes.Buffer)
	data = AddDefaultData(data)
	t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/pages/*.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		const layoutPath = "./templates/layouts/*.tmpl"
		layouts, err := filepath.Glob(layoutPath)
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(layoutPath)
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
