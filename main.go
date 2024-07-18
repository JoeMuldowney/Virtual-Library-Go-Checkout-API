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

	mux.HandleFunc("/go/getcartbook", cart.GetCartBook)
	mux.HandleFunc("/go/cart", cart.SaveCartItem)
	mux.HandleFunc("/go/checkout", cart.GetCheckOut)
	mux.HandleFunc("/go/delete", cart.DeleteCartItem)
	mux.HandleFunc("/go/deleteall", cart.DeleteAllCartItem)
	mux.HandleFunc("/go/membershipcard", billing.AddMembershipCard)
	mux.HandleFunc("/go/cartupdate/{id}", cart.UpdateCartItem)
	mux.HandleFunc("/go/addcard", billing.AddCard)
	mux.HandleFunc("/go/allcard", billing.GetAllCard)
	mux.HandleFunc("/go/billing", billing.GetCard)
	mux.HandleFunc("/go/updatecard", billing.UpdateCardPayment)
	mux.HandleFunc("/go/address", shipping.AddAddress)
	mux.HandleFunc("/go/shipping", shipping.GetAddress)
	mux.HandleFunc("/go/allshipping", shipping.GetAllAddress)
	mux.HandleFunc("/go/updateshipping", shipping.UpdateShippingAddress)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://3.129.70.204", "http://18.220.48.41", "https://csjoeportfolio.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		Debug:            true,
	})

	handler := c.Handler(mux)

	err := http.ListenAndServe(":8020", handler)
	if err != nil {
		return
	}

}
