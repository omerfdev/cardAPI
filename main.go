package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CardTypeResponse struct {
	CardType string `json:"cardType"`
}

func checkCardTypeHandler(w http.ResponseWriter, r *http.Request) {
	cardNumberStr := r.URL.Query().Get("cardNumber")
	if cardNumberStr == "" {
		http.Error(w, "Card number is required", http.StatusBadRequest)
		return
	}

	if len(cardNumberStr) != 16 {
		http.Error(w, "Invalid card number length", http.StatusBadRequest)
		return
	}

	firstDigit, err := strconv.Atoi(string(cardNumberStr[0]))
	if err != nil {
		http.Error(w, "Invalid card number", http.StatusBadRequest)
		return
	}

	var cardType string
	if firstDigit == 4 {
		cardType = "Visa"
	} else if firstDigit == 5 {
		cardType = "Mastercard"
	} else {
		cardType = "Not found"
	}

	response := CardTypeResponse{CardType: cardType}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/checkCardType", checkCardTypeHandler)
	http.ListenAndServe(":8080", nil)
}
