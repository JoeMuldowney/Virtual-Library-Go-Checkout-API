package shipping

import (
	"awesomeProject/api/cart"
	"awesomeProject/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Street         string `json:"street"`
	City           string `json:"city"`
	State          string `json:"state"`
	Zip            int    `json:"zip_code"`
	DefaultAddress bool   `json:"ship_default"`
	UserId         int    `json:"user_id"`
}

func AddAddress(w http.ResponseWriter, r *http.Request) {

	userId, err := cart.Verify(w, r)
	if err != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}
	var addAddress Address
	err = json.NewDecoder(r.Body).Decode(&addAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addAddress.UserId = userId
	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO shipping (first_name, last_name,  street, city, state, zip_code, ship_default, user_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(addAddress.FirstName, addAddress.LastName, addAddress.Street, addAddress.City, addAddress.State, addAddress.Zip, addAddress.DefaultAddress, addAddress.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "address inserted successfully")
}

func GetAddress(w http.ResponseWriter, r *http.Request) {

	userId, err := cart.Verify(w, r)
	if err != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var shippingAddress Address
	rows, err := db.Query("SELECT first_name, last_name, street, city, state, zip_code FROM shipping WHERE user_id=$1 AND ship_default=$2", userId, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&shippingAddress.FirstName, &shippingAddress.LastName, &shippingAddress.Street, &shippingAddress.City, &shippingAddress.State, &shippingAddress.Zip); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(shippingAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func GetAllAddress(w http.ResponseWriter, r *http.Request) {

	userId, err := cart.Verify(w, r)
	if err != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var shippingAddress []Address

	rows, err := db.Query("SELECT id, first_name, last_name, street, city, state, zip_code, ship_default FROM shipping WHERE user_id=$1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var location Address
		if err := rows.Scan(&location.Id, &location.FirstName, &location.LastName, &location.Street, &location.City, &location.State, &location.Zip, &location.DefaultAddress); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		shippingAddress = append(shippingAddress, location)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(shippingAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func UpdateShippingAddress(w http.ResponseWriter, r *http.Request) {

	userId, err := cart.Verify(w, r)
	if err != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}
	// Parse the request body to get the new address data
	var newAddress Address
	err = json.NewDecoder(r.Body).Decode(&newAddress)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt2, err := db.Prepare("UPDATE shipping SET ship_default=$1 WHERE user_id=$2 AND ship_default=$3")

	_, err = stmt2.Exec(false, userId, newAddress.DefaultAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("UPDATE shipping SET ship_default=$1 WHERE id=$2 AND user_id=$3")

	_, err = stmt.Exec(newAddress.DefaultAddress, newAddress.Id, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Address shipping updated successfully")
}
