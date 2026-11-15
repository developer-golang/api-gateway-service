package routes

import (
	"api-gateway-service/proxy"
	"net/http"
)

func NewRouter() http.Handler {
   

    userProxy := proxy.NewReverseProxy(
        "http://localhost:8083",
    )

    http.Handle("/createproduct", userProxy)
    http.Handle("/getproducts", userProxy)

    return http.DefaultServeMux
}