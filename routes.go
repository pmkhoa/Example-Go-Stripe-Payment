package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	router := httprouter.New()
	router.GET("/payment", handlePayment)
	router.POST("/payment", handlePayment)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
	})

	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", c.Handler(router)))
}

func handlePayment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stripeToken := r.FormValue("stripeToken")
	stripe.Key = os.Getenv("STRIPE_KEY")

	chargeParams := &stripe.ChargeParams{
		Amount:   45000,
		Currency: "usd",
		Desc:     "Charge for test@example.com",
	}
	chargeParams.SetSource(stripeToken)
	ch, err := charge.New(chargeParams)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(ch)
}
