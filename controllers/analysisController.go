package controllers

import (
	. "golang-dev-logic-challenge-devphilip/model"
	"strings"
)

// Range of prices [80, 90, 100, 110, 120]
var prices = []float64{80, 90, 100, 110, 120}

// AnalysisHandler returns the analysis of a list of options contracts
// @Summary returns options contracts analysis
// @Description returns the analysis of a list of options contracts
// @Tags Options Analysis
// @Accept json
// @Produce json
// @Param contracts body []OptionsContract true "List of options contracts"
// @Success 200 {object} AnalysisResponse
// @Failure 400 {object} ErrorResponse
// @Router /analyze [post]
func AnalysisHandler(contracts []OptionsContract) AnalysisResponse {

	// Perform analysis
	xyValues := calculateXYValues(contracts)
	maxProfit := calculateMaxProfit(contracts)
	maxLoss := calculateMaxLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	// Create and return response
	return AnalysisResponse{
		XYValues:        xyValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}

}

func calculateXYValues(contracts []OptionsContract) []XYValue {
	var xyValues []XYValue

	// Calculate total profit or loss for each contract using the price range
	for _, price := range prices {
		totalProfitOrLoss := 0.0
		for _, contract := range contracts {
			totalProfitOrLoss += calculateProfitOrLoss(contract, price)
		}
		xyValues = append(xyValues, XYValue{X: price, Y: totalProfitOrLoss})
	}

	return xyValues
}

func calculateMaxProfit(contracts []OptionsContract) float64 {
	maxProfit, _ := calculateMetrics(contracts, prices)
	return maxProfit
}

func calculateMaxLoss(contracts []OptionsContract) float64 {
	_, maxLoss := calculateMetrics(contracts, prices)
	return maxLoss
}

func calculateBreakEvenPoints(contracts []OptionsContract) []float64 {
	var breakEvenPoints []float64

	for _, contract := range contracts {
		contractType := strings.ToLower(contract.Type)
		contractLongShort := strings.ToLower(contract.LongShort)
		if contractLongShort == "long" {
			if contractType == "call" {
				//Call: strikePrice+Premium // Premium = Ask
				breakEvenPoints = append(breakEvenPoints, contract.StrikePrice+contract.Ask)
			} else if contractType == "put" {
				//Put: strikePrice−Premium // Premium = Ask
				breakEvenPoints = append(breakEvenPoints, contract.StrikePrice-contract.Ask)
			}
		} else if contractLongShort == "short" {
			if contractType == "call" {
				//Call: strikePrice+Premium // Premium = Bid
				breakEvenPoints = append(breakEvenPoints, contract.StrikePrice+contract.Bid)
			} else if contractType == "put" {
				//Put: strikePrice−Premium // Premium = Bid
				breakEvenPoints = append(breakEvenPoints, contract.StrikePrice-contract.Bid)
			}
		}
	}

	return breakEvenPoints
}

// Calculate the profit or loss for a single option at a given price
func calculateProfitOrLoss(contract OptionsContract, priceAtExpiry float64) float64 {

	var profitOrLoss float64
	contractType := strings.ToLower(contract.Type)
	contractLongShort := strings.ToLower(contract.LongShort)
	if contractLongShort == "long" {
		if contractType == "call" {
			// Long Call: Profit/Loss=max(0,priceAtExpiry−strikePrice)−Ask
			profitOrLoss = max(0, priceAtExpiry-contract.StrikePrice) - contract.Ask
		} else if contractType == "put" {
			//Long Put: Profit/Loss=max(0,strikePrice−priceAtExpiry)−Ask
			profitOrLoss = max(0, contract.StrikePrice-priceAtExpiry) - contract.Ask
		}
	} else if contractLongShort == "short" {
		if contractType == "call" {
			// Short Call: Profit/Loss=Bid−max(0, priceAtExpiry−strikePrice)
			profitOrLoss = contract.Bid - max(0, priceAtExpiry-contract.StrikePrice)
		} else if contractType == "put" {
			//Short Put: Profit/Loss=Bid−max(0, strikePrice−priceAtExpiry)
			profitOrLoss = contract.Bid - max(0, contract.StrikePrice-priceAtExpiry)
		}
	}

	return profitOrLoss
}

// Calculate the maximum profit, maximum loss, and break-even points
func calculateMetrics(contracts []OptionsContract, prices []float64) (maxProfit, maxLoss float64) {
	maxProfit = -1e9
	maxLoss = 1e9

	for _, price := range prices {
		totalProfitOrLoss := 0.0
		for _, contract := range contracts {
			totalProfitOrLoss += calculateProfitOrLoss(contract, price)
		}
		if totalProfitOrLoss > maxProfit {
			maxProfit = totalProfitOrLoss
		}
		if totalProfitOrLoss < maxLoss {
			maxLoss = totalProfitOrLoss
		}
	}

	return
}
