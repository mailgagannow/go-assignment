package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

type TaxBrackets struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max,omitempty"`
	Rate float64 `json:"rate"`
}

/*
GetTaxBrackets retrieves tax brackets data for a specific year based on the given year.

@param year int - the year is used to find the JSON file
@return []TaxBrackets - a slice of TaxBrackets structs
@return error - the error message if an error occurred during file reading or unmarshalling
*/
func GetTaxBrackets(year int) ([]TaxBrackets, error) {
	fmt.Println(year)

	filename := fmt.Sprintf("tax-files/tax-brackets--%d.json", year)
	fmt.Println(filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {

		return nil, err
	}

	var brackets []TaxBrackets
	err = json.Unmarshal(data, &brackets)
	if err != nil {
		return nil, err
	}

	return brackets, nil
}

/*
SendResponse function accepts an http.ResponseWriter and a map[string]interface{} as input parameters and used to send json response.

@param w http.ResponseWriter
@param results map[string]interface{}
@return None
*/
func SendResponse(w http.ResponseWriter, results map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if message, ok := results["error"]; ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}


/*
CalculateTax function takes a map of tax brackets and income, and calculates the federal tax amount, effective tax rate, and tax slabs.

@param income float to calculate the tax.
@return results map[string]interface{}
*/
func CalculateTax(results map[string]interface{}, income float64) map[string]interface{} {

	currentIncome := income
	remainingIncome := currentIncome

	if brackets, ok := results["tax_brackets"].([]TaxBrackets); ok {
		federalTax := 0.0
		var slabs []map[string]interface{}

		// loop over the brackets slice
		for _, bracket := range brackets {
			currentSlabTax := 0.0
			diff := 0.0

			if remainingIncome <= 0 {
				break
			}

			maxValue := bracket.Max
			minValue := bracket.Min

			if maxValue == 0 {
				diff = minValue
			} else {
				diff = maxValue - minValue
			}

			taxableIncome,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Min(diff, remainingIncome)), 64)
			currentSlabTax, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", taxableIncome*bracket.Rate), 64)
			federalTax += currentSlabTax
			remainingIncome -= taxableIncome

			slab := map[string]interface{}{
				"min":            bracket.Min,
				"max":            bracket.Max,
				"rate":           bracket.Rate,
				"tax":            currentSlabTax,
				"taxable_income": taxableIncome,
			}
			
			slabs = append(slabs, slab)

		}

		effectiveTaxRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", (federalTax/currentIncome)*100), 64)

		fmt.Println(federalTax)
		fmt.Println(effectiveTaxRate)
		results["totalIncome"] = currentIncome
		results["taxAmount"] = federalTax
		results["effectiveTaxRate"] = effectiveTaxRate
		results["taxPerSlab"] = slabs

		//fmt.Println("%s results \n", results)
	}

	return results
}