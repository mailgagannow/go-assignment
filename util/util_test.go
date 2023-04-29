package util

import (
	"reflect"
	"testing"

)

func TestGetTaxBrackets(t *testing.T) {
	year := 2022
	expected := []TaxBrackets{
		{0, 50197, 0.15},
		{50197, 100392, 0.205},
		{100392, 155625, 0.26},
		{155625, 221708, 0.29},
		{221708, 0, 0.33},
	}

	brackets, err := GetTaxBrackets(year)

	if err != nil {
		t.Errorf("GetTaxBrackets(%d) returned an error: %v", year, err)
	}

	if !reflect.DeepEqual(brackets, expected) {
		t.Errorf("GetTaxBrackets(%d) returned %v, expected %v", year, brackets, expected)
	}
}
