module api-gateway-service

go 1.26.4

require github.com/developer-golang/payment-service v0.0.0

require github.com/stripe/stripe-go/v86 v86.2.0 // indirect

replace github.com/developer-golang/payment-service => ../payment-service
