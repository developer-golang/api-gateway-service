package routes

import (
	"api-gateway-service/proxy"
	"net/http"
)

func NewRouter() http.Handler {

	orderProxy := proxy.NewReverseProxy(
		"http://localhost:8082",
	)
	// Product and Order routes
	http.Handle("/createorder", orderProxy)

	payproxy := proxy.NewReverseProxy(
		"https://088dptzh-4242.inc1.devtunnels.ms",
	)
	// Stripe routes
	http.Handle("/api/create-checkout-session", payproxy)
	http.Handle("/createstripesession", payproxy)
	http.Handle("/success", payproxy)
	http.Handle("/api/cancel", payproxy)
	http.Handle("/webhook", payproxy)

	productproxy := proxy.NewReverseProxy(
		"http://localhost:8083",
	)
	http.Handle("/createproduct", productproxy)
	http.Handle("/getproducts", productproxy)
	return http.DefaultServeMux
}
