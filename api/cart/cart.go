package cart

import (
	"awesomeProject/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type UserCart struct {
	Id           int     `json:"id"`
	UserId       int     `json:"user_id"`
	BookId       int     `json:"book_id"`
	BookTitle    string  `json:"title"`
	Quantity     int     `json:"quantity"`
	Format       string  `json:"format"`
	PurchaseType string  `json:"purchase"`
	Cost         float64 `json:"cost"`
}

type Checkout struct {
	CartItems  []UserCart `json:"cart_items"`
	TotalItems int        `json:"total_items"`
	TotalCost  float64    `json:"total_cost"`
}

func GetCheckOut(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.URL.Query().Get("user")

	fmt.Println("Received userId:", userIdStr)

	if userIdStr == "" {
		http.Error(w, "Missing id or userId parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId parameter", http.StatusBadRequest)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT book_id, book_title, quantity, purchase_type, cost, format FROM cart WHERE user_id = $1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var cartItems []UserCart
	var totalItems int
	var totalCost float64

	for rows.Next() {
		var item UserCart

		if err := rows.Scan(&item.BookId, &item.BookTitle, &item.Quantity, &item.PurchaseType, &item.Cost, &item.Format); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		intToFloat := float64(item.Quantity)
		totalCost += (intToFloat * item.Cost)
		totalItems += item.Quantity
		cartItems = append(cartItems, item)

	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the response struct
	response := Checkout{
		CartItems:  cartItems,
		TotalItems: totalItems,
		TotalCost:  totalCost,
	}
	// Convert cartItems to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func GetCartBook(w http.ResponseWriter, r *http.Request) {

	bookIdStr := r.URL.Query().Get("id")
	userIdStr := r.URL.Query().Get("user")
	fmt.Println("Received bookId:", bookIdStr)
	fmt.Println("Received userId:", userIdStr)

	if bookIdStr == "" || userIdStr == "" {
		http.Error(w, "Missing id or userId parameter", http.StatusBadRequest)
		return
	}

	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, "Invalid bookId parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId parameter", http.StatusBadRequest)
		return
	}
	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM cart where book_id = $1 AND user_id = $2", bookId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var cartItems []UserCart
	for rows.Next() {
		var item UserCart
		if err := rows.Scan(&item.Id, &item.UserId, &item.BookId, &item.BookTitle, &item.Quantity, &item.PurchaseType, &item.Cost, &item.Format); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cartItems = append(cartItems, item)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert cartItems to JSON
	jsonData, err := json.Marshal(cartItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {

	bookIdStr := r.URL.Query().Get("id")
	userIdStr := r.URL.Query().Get("user")
	fmt.Println("Received bookId:", bookIdStr)
	fmt.Println("Received userId:", userIdStr)

	if bookIdStr == "" || userIdStr == "" {
		http.Error(w, "Missing id or userId parameter", http.StatusBadRequest)
		return
	}

	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		http.Error(w, "Invalid bookId parameter", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId parameter", http.StatusBadRequest)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM cart WHERE book_id = $1 AND user_id = $2")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	// Execute the SQL statement
	_, err = stmt.Exec(bookId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "book deleted successfully")
}
func DeleteAllCartItem(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.URL.Query().Get("user")

	fmt.Println("Received userId:", userIdStr)

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId parameter", http.StatusBadRequest)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM cart WHERE user_id = $1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	// Execute the SQL statement
	_, err = stmt.Exec(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "book deleted successfully")
}
func SaveCartItem(w http.ResponseWriter, r *http.Request) {

	// Parse the request body to get the cart item data
	var newItem UserCart
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cart(user_id, book_id, book_title, quantity, format, purchase_type, cost) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer stmt.Close()
	// Execute the SQL statement
	_, err = stmt.Exec(newItem.UserId, newItem.BookId, newItem.BookTitle, newItem.Quantity, newItem.Format, newItem.PurchaseType, newItem.Cost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "book inserted successfully")
}
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {

	// Parse the bookId and userId from the URL
	bookId := r.URL.Path[len("/cartupdate/"):]

	// Parse the request body to get the cart item data
	var updateItem UserCart
	err := json.NewDecoder(r.Body).Decode(&updateItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId := updateItem.UserId
	fmt.Println(userId)
	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("UPDATE cart SET quantity=$1, format=$2, purchase_type=$3, cost=$4 WHERE book_id = $5 AND user_id = $6")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	// Execute the SQL statement
	_, err = stmt.Exec(updateItem.Quantity, updateItem.Format, updateItem.PurchaseType, updateItem.Cost, bookId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "cart updated successfully")
}

func Verify(w http.ResponseWriter, r *http.Request) (int, error) {
	sessionCookie, err := r.Cookie("sessionid")
	if err != nil {
		http.Error(w, "Session cookie not found", http.StatusUnauthorized)
		return 0, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/users/verify/", nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return 0, err
	}
	req.AddCookie(sessionCookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error forwarding cookie", http.StatusInternalServerError)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return 0, err
	}

	// Parse the response body to extract user_id
	var userResponse struct {
		UserID string `json:"user_id"`
		// Add more fields if needed
	}
	if err := json.Unmarshal(body, &userResponse); err != nil {
		http.Error(w, "Error parsing user data", http.StatusInternalServerError)
		return 0, err
	}

	userId, err := strconv.Atoi(userResponse.UserID)
	if err != nil {
		http.Error(w, "Error converting user ID to integer", http.StatusInternalServerError)
		return 0, err
	}

	return userId, nil
}
