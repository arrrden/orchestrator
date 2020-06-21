package handlers

import (
	"log"
	"net/http"

	"github.com/arrrden/orchestrator/data"
)

type Catalogue struct {
	l *log.Logger
}

func (c *Catalogue) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getCatalogue(w, r)
		break
	case http.MethodPost:
		c.addWork(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *Catalogue) getCatalogue(w http.ResponseWriter, r *http.Request) {
	wl := data.GetCatalogue()
	err := wl.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
	}
}

func (c *Catalogue) addWork(w http.ResponseWriter, r *http.Request) {
	work := &data.Catalogue{}
	err := work.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to dencode JSON", http.StatusBadRequest)
	}
}

func NewCatalogue(l *log.Logger) *Catalogue {
	return &Catalogue{l}
}
