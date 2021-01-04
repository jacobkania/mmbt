package handlers

import (
	"log"
	"mmbt/configuration"
	"mmbt/constants"
	"mmbt/handlers/core"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	mux "github.com/gorilla/mux"
)

// SetRoutes sets all routes on the provided mux router
func SetRoutes(r *mux.Router) {
	/* API Routes */

	r.HandleFunc("/login", core.LoginHandler)
	r.HandleFunc("/register", core.RegisterAccountHandler)

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	RouterV1(apiV1)

	/* Frontend Routes */

	if configuration.Config.Environment == constants.Environment.Development {
		r.HandleFunc("/{rest:.*}", handleFrontendDevServer)
	} else {
		r.Handle("/{rest:.*}", http.FileServer(http.Dir("./js")))
	}
}

func handleFrontendDevServer(w http.ResponseWriter, r *http.Request) {
	// Proxy frontend routes to dev server when environment is Development
	redirectURL, err := url.Parse("http://" + configuration.Config.Host + ":" + strconv.Itoa(configuration.Config.FrontendFwdPort))
	if err != nil {
		log.Fatal("FRONTEND_FWD_PORT env must be set when in dev environment")
	}

	proxy := httputil.NewSingleHostReverseProxy(redirectURL)

	r.URL.Path = mux.Vars(r)["rest"]
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(w, r)
}
