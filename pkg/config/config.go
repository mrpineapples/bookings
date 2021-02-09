package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	InfoLog       *log.Logger
	IsProduction  bool
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
	UseCache      bool
}
