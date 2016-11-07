package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var KYP_GATEWAY_ENDPOINT = os.Getenv("KYP_GATEWAY_ENDPOINT")
var KYP_PROXY_PORT = os.Getenv("KYP_PROXY_PORT")

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   KYP_GATEWAY_ENDPOINT,
	})

	proxy.Director = director

	log.Print("Starting Kyp Proxy at " + KYP_PROXY_PORT)

	http.ListenAndServe(":"+KYP_PROXY_PORT, proxy)
}

func director(r *http.Request) {
	r.URL.Host = KYP_GATEWAY_ENDPOINT
	r.URL.Scheme = "http"
	r.Header.Set("Kyp-Proxy", "KypProxy")
}
