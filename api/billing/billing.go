package billing

import (
	"awesomeProject/config"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Card struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CardNum     string `json:"card_num"`
	PaymentType string `json:"payment_type"`
	ExpDate     string `json:"exp_date"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip_code"`
	PayDefault  bool   `json:"pay_default"`
	UserId      int    `json:"user_id"`
}

func AddMembershipCard(w http.ResponseWriter, r *http.Request) {

	var addCard Card
	err := json.NewDecoder(r.Body).Decode(&addCard)
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

	stmt, err := db.Prepare("INSERT INTO payment (first_name, last_name, card_num, payment_type, exp_date, street, city, state, zip_code, pay_default, user_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	eFirstName, err := Encrypt(addCard.FirstName, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eLastName, err := Encrypt(addCard.LastName, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eCardNum, err := Encrypt(addCard.CardNum, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	ePaymentType, err := Encrypt(addCard.PaymentType, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eExpDate, err := Encrypt(addCard.ExpDate, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eStreet, err := Encrypt(addCard.Street, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eCity, err := Encrypt(addCard.City, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eState, err := Encrypt(addCard.State, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eZip, err := Encrypt(addCard.State, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Execute the SQL statement
	_, err = stmt.Exec(eFirstName, eLastName, eCardNum, ePaymentType, eExpDate, eStreet, eCity, eState, eZip, addCard.PayDefault, addCard.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "card inserted successfully")
}

func AddCard(w http.ResponseWriter, r *http.Request) {

	var addCard Card
	err := json.NewDecoder(r.Body).Decode(&addCard)
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

	stmt, err := db.Prepare("INSERT INTO payment (first_name, last_name, card_num, payment_type, exp_date, street, city, state, zip_code, pay_default, user_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()
	eFirstName, err := Encrypt(addCard.FirstName, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eLastName, err := Encrypt(addCard.LastName, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eCardNum, err := Encrypt(addCard.CardNum, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	ePaymentType, err := Encrypt(addCard.PaymentType, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eExpDate, err := Encrypt(addCard.ExpDate, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eStreet, err := Encrypt(addCard.Street, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eCity, err := Encrypt(addCard.City, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eState, err := Encrypt(addCard.State, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eZip, err := Encrypt(addCard.State, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Execute the SQL statement
	_, err = stmt.Exec(eFirstName, eLastName, eCardNum, ePaymentType, eExpDate, eStreet, eCity, eState, eZip, addCard.PayDefault, addCard.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "card inserted successfully")
}

func GetCard(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Path[len("/billing/"):]

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var billingCard Card
	rows, err := db.Query("SELECT id, first_name, last_name, card_num, payment_type, exp_date, street, city, state, zip_code FROM payment WHERE user_id=$1 AND pay_default=$2", userId, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&billingCard.Id, &billingCard.FirstName, &billingCard.LastName, &billingCard.CardNum, &billingCard.PaymentType, &billingCard.ExpDate, &billingCard.Street, &billingCard.City, &billingCard.State, &billingCard.Zip); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	dFirstName, err := Decrypt(billingCard.FirstName, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dLastName, err := Decrypt(billingCard.LastName, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dCardNum, err := Decrypt(billingCard.CardNum, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	ending := lastFour(dCardNum)
	dPaymentType, err := Decrypt(billingCard.PaymentType, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dExpDate, err := Decrypt(billingCard.ExpDate, config.GetSecretKey())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	billingCard.CardNum = ending
	billingCard.FirstName = dFirstName
	billingCard.PaymentType = dPaymentType
	billingCard.ExpDate = dExpDate
	billingCard.LastName = dLastName
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(billingCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func GetAllCard(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Path[len("/allcard/"):]

	connectionString := config.GetConnectionString()
	db, err := config.OpenConnection(connectionString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var card []Card

	rows, err := db.Query("SELECT id, first_name, last_name, card_num, payment_type, exp_date FROM payment WHERE user_id=$1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cards Card
		if err := rows.Scan(&cards.Id, &cards.FirstName, &cards.LastName, &cards.CardNum, &cards.PaymentType, &cards.ExpDate); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dCardNum, err := Decrypt(cards.CardNum, config.GetSecretKey())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		endingIn := lastFour(dCardNum)

		dPaymentType, err := Decrypt(cards.PaymentType, config.GetSecretKey())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		dExpDate, err := Decrypt(cards.ExpDate, config.GetSecretKey())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		cards.CardNum = endingIn
		cards.PaymentType = dPaymentType
		cards.ExpDate = dExpDate
		card = append(card, cards)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func UpdateCardPayment(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Path[len("/updatecard/"):]

	var newCard Card
	err := json.NewDecoder(r.Body).Decode(&newCard)
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

	stmt2, err := db.Prepare("UPDATE payment SET pay_default=$1 WHERE user_id=$2 AND pay_default=$3")

	_, err = stmt2.Exec(false, userId, newCard.PayDefault)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("UPDATE payment SET pay_default=$1 WHERE id=$2 AND user_id=$3")

	_, err = stmt.Exec(newCard.PayDefault, newCard.Id, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Address shipping updated successfully")
}

func Encrypt(plainText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(cipherText, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", err
	}

	nonce, cipherText := data[:nonceSize], string(data[nonceSize:])
	plainText, err := gcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func lastFour(s string) string {
	return s[len(s)-4:]
}
