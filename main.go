package main

import (
	"awesomeProject/api/billing"
	"awesomeProject/api/cart"
	"awesomeProject/api/shipping"
	"github.com/rs/cors"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/data/{id}", cart.GetCartBook)
	mux.HandleFunc("/cart", cart.SaveCartItem)
	mux.HandleFunc("/checkout", cart.GetCheckOut)
	mux.HandleFunc("/delete/{id}/{userId}", cart.DeleteCartItem)
	mux.HandleFunc("/deleteall/{userId}", cart.DeleteAllCartItem)
	mux.HandleFunc("/membershipcard", billing.AddMembershipCard)
	mux.HandleFunc("/cartupdate/{id}", cart.UpdateCartItem)
	mux.HandleFunc("/card", billing.AddCard)
	mux.HandleFunc("/allcard/{userId}", billing.GetAllCard)
	mux.HandleFunc("/billing/{userId}", billing.GetCard)
	mux.HandleFunc("/updatecard/{userId}", billing.UpdateCardPayment)
	mux.HandleFunc("/address", shipping.AddAddress)
	mux.HandleFunc("/shipping/{userId}", shipping.GetAddress)
	mux.HandleFunc("/allshipping/{userId}", shipping.GetAllAddress)
	mux.HandleFunc("/updateshipping/{userId}", shipping.UpdateShippingAddress)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://18.218.222.138", "http://18.220.48.41:3000", "http://18.220.48.41:8000", "http://localhost:3000", "http://localhost:8000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		Debug:            true,
	})

	// Insert the middleware
	handler := c.Handler(mux)

	err := http.ListenAndServe(":8020", handler)
	if err != nil {
		return
	}

}
