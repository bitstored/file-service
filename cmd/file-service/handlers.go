package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
)

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("Allow %v %v \n", r.Body, r.Header)
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				fmt.Println("Handling preflight OPTIONS request")
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "authorization"}
	w.Header().Add("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
	w.Header().Add("Access-Control-Allow-Methods", strings.Join(methods, ","))
	exposed := []string{"authorization, Accept, Content-Type"}
	w.Header().Add("Access-Control-Expose-Headers", strings.Join(exposed, ","))
	w.Header().Add("Allow", strings.Join(methods, ","))
	glog.Infof("preflight request for %s", r.URL.Path)
}
