package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ReverseProxy(target string, prefix string) http.Handler {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		r.Host = targetURL.Host
		log.Printf("Proxy forwarding to %s%s", target, r.URL.Path)
		proxy.ServeHTTP(w, r)
	})
}
