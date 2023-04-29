package util

import (
	"reflect"
	"testing"
)

/*
TestGetTaxBrackets tests the GetTaxBrackets function to ensure that it returns the expected tax brackets for a given year.
*/
func TestGetTaxBrackets(t *testing.T) {
	year := 2022
	expected := []TaxBrackets{
		{0, 50197, 0.15},
		{50197, 100392, 0.205},
		{100392, 155625, 0.26},
		{155625, 221708, 0.29},
		{221708, 0, 0.33},
	}

	// Call the function
	brackets, err := GetTaxBrackets(year)

	if err != nil {
		t.Errorf("GetTaxBrackets(%d) returned an error: %v", year, err)
	}

	if !reflect.DeepEqual(brackets, expected) {
		t.Errorf("GetTaxBrackets(%d) returned %v, expected %v", year, brackets, expected)
	}
}

/*
TestCalculateTax is a unit test function to test the CalculateTax function
It sets up the required test data using the setup function and calls the CalculateTax function with it.
It then verifies the expected output by checking the existence of the required fields in the result.
*/
func TestCalculateTax(t *testing.T) {

	income := 50000.0
	brackets := setup()

	// Call the function
	result := CalculateTax(brackets, income)

	// Verify expected output
	if _, ok := result["totalIncome"]; !ok {
		t.Error("Expected totalIncome field in the result")
	}
	if _, ok := result["taxAmount"]; !ok {
		t.Error("Expected taxAmount field in the result")
	}
	if _, ok := result["effectiveTaxRate"]; !ok {
		t.Error("Expected effectiveTaxRate field in the result")
	}
	if _, ok := result["taxPerSlab"]; !ok {
		t.Error("Expected taxPerSlab field in the result")
	}
}

func setup() map[string]interface{} {
	// Perform setup tasks, such as loading test data or setting up a database connection
	results := make(map[string]interface{})
	brackets, _ := GetTaxBrackets(2022)
	results["tax_brackets"] = brackets

	return results
}
