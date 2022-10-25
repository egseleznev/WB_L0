package http_server

import (
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"html/template"
	"net/http"
)

// HttpServer represents entity of HttpServer
// config is a handle to the specified server config
// router is a handle to navigate by endpoints
// cache is a handle to an application cache
type HttpServer struct {
	config *Config
	router *mux.Router
	cache  *cache.Cache
}

// New creates and return a HttpServer handle
func New(config *Config, cache *cache.Cache) *HttpServer {

	return &HttpServer{
		config: config,
		router: mux.NewRouter(),
		cache:  cache,
	}
}

// Start starts a HttpServer with configured endpoints
func (h *HttpServer) Start() error {

	h.configureRouter()

	return http.ListenAndServe(h.config.Addr, h.router)
}

// configureRouter concatenates endpoints and their handles
func (h *HttpServer) configureRouter() {

	h.router.HandleFunc("/orders", h.handleOrders())
}

// handleOrders reads Uid from the form, gets the data from the cache by this Uid
// and return the rendered template
func (h *HttpServer) handleOrders() http.HandlerFunc {

	var formData struct {
		UID  string
		Data string
	} // structure for template data

	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("./http_server/templates/main.tmpl") // parse template

		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		formData.UID = r.FormValue("UID") // read Uid from form

		v, flag := h.cache.Get(formData.UID) // get the cache for this Uid

		if flag == false {
			err = ts.Execute(w, nil) // if cache not found, render form without data
		} else {
			formData.Data = string(v.([]byte)[:]) // otherwise convert byte stream to string
			err = ts.Execute(w, formData)         // then render template with found cache
		}

		if err != nil {
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
