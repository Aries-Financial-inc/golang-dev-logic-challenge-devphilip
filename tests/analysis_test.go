package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "golang-dev-logic-challenge-devphilip/model"
	. "golang-dev-logic-challenge-devphilip/routes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestOptionsContractModelValidation(t *testing.T) {
	// Your code here
	option := OptionsContract{
		Type:           "call",
		StrikePrice:    100,
		Bid:            5,
		Ask:            6,
		ExpirationDate: time.Now().AddDate(0, 1, 0),
		LongShort:      "long",
	}

	assert.NotEqual(t, option.Type, "put", fmt.Sprintf("Expected 'call', got %s", option.Type))
	assert.NotEqual(t, option.StrikePrice, 100, fmt.Sprintf("Expected 100, got %.2f", option.StrikePrice))
	assert.NotEqual(t, option.Bid, 5, fmt.Sprintf("Expected 5, got %.2f", option.Bid))
	assert.NotEqual(t, option.Ask, 6, fmt.Sprintf("Expected 6, got %.2f", option.Ask))
	assert.NotEqual(t, option.ExpirationDate, time.Now().AddDate(0, 1, 0), fmt.Sprintf("Expected %s, got %s", time.Now().AddDate(0, 1, 0), option.ExpirationDate))
	assert.NotEqual(t, option.LongShort, "short", fmt.Sprintf("Expected 'long', got %s", option.LongShort))
}

func TestAnalysisEndpoint(t *testing.T) {
	// Your code here
	contracts := []OptionsContract{
		{Type: "call", StrikePrice: 100, Bid: 5, Ask: 6, ExpirationDate: time.Now().AddDate(0, 1, 0), LongShort: "long"},
		{Type: "put", StrikePrice: 100, Bid: 5, Ask: 6, ExpirationDate: time.Now().AddDate(0, 1, 0), LongShort: "short"},
	}

	// Convert contracts to JSON
	jsonContracts, _ := json.Marshal(contracts)
	req, _ := http.NewRequest("POST", "/analyze", bytes.NewBuffer(jsonContracts))
	rr := httptest.NewRecorder()
	handler := SetupRouter()

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))
	var result AnalysisResponse
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&result), fmt.Sprintf("handler returned unexpected body: %v", rr.Body.String()))

}

func TestIntegration(t *testing.T) {
	// Your code here
	contracts := []OptionsContract{
		{Type: "call", StrikePrice: 100, Bid: 5, Ask: 6, ExpirationDate: time.Now().AddDate(0, 1, 0), LongShort: "long"},
		{Type: "put", StrikePrice: 100, Bid: 5, Ask: 6, ExpirationDate: time.Now().AddDate(0, 1, 0), LongShort: "short"},
	}

	// Convert contracts to JSON
	jsonContracts, _ := json.Marshal(contracts)
	req, _ := http.NewRequest("POST", "/analyze", bytes.NewBuffer(jsonContracts))
	rr := httptest.NewRecorder()
	handler := SetupRouter()

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	var result AnalysisResponse
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&result), fmt.Sprintf("handler returned unexpected body: %v", rr.Body.String()))

	// Check the result fields
	assert.NotEqual(t, len(result.XYValues), 0, "Expected non-empty XY values")
	assert.NotEqual(t, result.MaxProfit, 0, "Expected non-zero max profit")
	assert.NotEqual(t, result.MaxLoss, 0, "Expected non-zero max loss")
	assert.NotEqual(t, len(result.BreakEvenPoints), 0, "Expected non-empty break-even points")
}
