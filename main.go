package main

import (
	"awesomeProject/api/billing"
	"awesomeProject/api/cart"
	"awesomeProject/api/shipping"
	"github.com/rs/cors"
	"net/http"
	"strconv"
)

// verifyHandler function to use in the server
func verifyHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := cart.Verify(w, r)
	if err != nil {
		// Error has already been handled in Verify, so just return
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(userId)))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/data/{id}", cart.GetCartBook)
	mux.HandleFunc("/cart", cart.SaveCartItem)
	mux.HandleFunc("/checkout", cart.GetCheckOut)
	mux.HandleFunc("/delete/{id}", cart.DeleteCartItem)
	mux.HandleFunc("/deleteall", cart.DeleteAllCartItem)
	mux.HandleFunc("/cartupdate/{id}", cart.UpdateCartItem)
	mux.HandleFunc("/card", billing.AddCard)
	mux.HandleFunc("/allcard", billing.GetAllCard)
	mux.HandleFunc("/billing", billing.GetCard)
	mux.HandleFunc("/updatecard", billing.UpdateCardPayment)
	mux.HandleFunc("/address", shipping.AddAddress)
	mux.HandleFunc("/shipping", shipping.GetAddress)
	mux.HandleFunc("/allshipping", shipping.GetAllAddress)
	mux.HandleFunc("/updateshipping", shipping.UpdateShippingAddress)
	//mux.HandleFunc("/verify", verifyHandler)

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
