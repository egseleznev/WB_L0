package http_server

import (
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"html/template"
	"net/http"
)

type HttpServer struct {
	config *Config
	router *mux.Router
	cache  *cache.Cache
}

func New(config *Config, cache *cache.Cache) *HttpServer {
	return &HttpServer{
		config: config,
		router: mux.NewRouter(),
		cache:  cache,
	}
}

func (h *HttpServer) Start() error {
	h.configureRouter()
	return http.ListenAndServe(h.config.Addr, h.router)
}

func (h *HttpServer) configureRouter() {
	h.router.HandleFunc("/orders", h.handleOrders())
}

func (h *HttpServer) handleOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("./http_server/templates/main.tmpl")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		var formData struct {
			UID  string
			Data string
		}

		formData.UID = r.FormValue("UID")

		v, flag := h.cache.Get(formData.UID)
		if flag == false {
			err = ts.Execute(w, nil)
		} else {
			formData.Data = string(v.([]byte)[:])
			err = ts.Execute(w, formData)
		}
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
