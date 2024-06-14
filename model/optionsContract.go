package model

import "time"

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	// type (call or put), strike_price, bid, ask, expiration_date, long_short
	Type           string    `json:"type"`
	StrikePrice    float64   `json:"strike_price"`
	Bid            float64   `json:"bid"`
	Ask            float64   `json:"ask"`
	ExpirationDate time.Time `json:"expiration_date"`
	LongShort      string    `json:"long_short"`
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
