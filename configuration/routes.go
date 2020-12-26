package configuration

import (
	"log"
	"mmbt/constants"
	"mmbt/handlers"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	mux "github.com/gorilla/mux"
)

// SetRoutes sets all routes on the provided mux router
func setRoutes(server *Server) {
	var r *mux.Router = server.Router

	/* API Routes */
	r.HandleFunc("/api", handlers.IndexHandler)

	/* Frontend Routes */
	if server.Config.Environment == constants.Environment.Development {
		// Proxy frontend routes to dev server when environment is Development
		redirectURL, err := url.Parse("http://" + server.Config.Host + ":" + strconv.Itoa(server.Config.DevFrontendPort))
		if err != nil {
			log.Panic("Dev Frontend Port env variable must be set when in dev environment")
		}
		handleFrontendDevBuild(r, redirectURL)
	} else {
		r.Handle("/{rest:.*}", http.FileServer(http.Dir("./js/build")))
	}
}

func handleFrontendDevBuild(r *mux.Router, redirectToURL *url.URL) {
	r.HandleFunc("/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {

		proxy := httputil.NewSingleHostReverseProxy(redirectToURL)

		r.URL.Path = mux.Vars(r)["rest"]
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		proxy.ServeHTTP(w, r)
	})
}
