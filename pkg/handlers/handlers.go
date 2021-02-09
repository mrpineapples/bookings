package handlers

import (
	"net/http"

	"github.com/mrpineapples/bookings/pkg/config"
	"github.com/mrpineapples/bookings/pkg/models"
	"github.com/mrpineapples/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home serves the base route
func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	re.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.tmpl", &models.TemplateData{})
}

// About serves the about route
func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := re.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
