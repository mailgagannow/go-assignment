package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"points-interview/packages/util"
)

type Error struct {
	Message string `json:"message"`
}

/*
Handler function for the /tax-calculator route.
Accepts a GET request.
Returns the tax brackets for 2022 year.
*/
func calculator(w http.ResponseWriter, r *http.Request) {
	results := make(map[string]interface{})

	brackets, err := util.GetTaxBrackets(2022)
	if err != nil {
		results["error"] = Error{Message: err.Error()}
	} else {
		results["tax_brackets"] = brackets
	}

	util.SendResponse(w, results)
}

/*
Handler function for the /tax-calculator/tax-year/ route.
Accepts a POST request with the tax-year as a url parameter and optional post parameter income to calculate the tax.
Returns the tax brackets for the tax year passed and calculate the tax if income is passed.
*/
func tax(w http.ResponseWriter, r *http.Request) {
	results := make(map[string]interface{})
	
	vars := mux.Vars(r)
	//Convert string tax-year to integer tax-year
	taxYear, err := strconv.Atoi(vars["tax-year"])
	if err != nil {
		// Handle the error
		results["error"] = Error{Message: err.Error()}
		util.SendResponse(w, results)
		return

	}

	if taxYear > 2022 || taxYear < 2019 {
		results["error"] = Error{Message: "Please enter year between 2019 and 2022"}
		util.SendResponse(w, results)
		return
	}

	//calculate the tax bracket for the year passed
	brackets, err := util.GetTaxBrackets(taxYear)

	if err != nil {
		results["error"] = Error{Message: err.Error()}
		util.SendResponse(w, results)
		return
	} else {
		results["tax_brackets"] = brackets
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			results["error"] = Error{Message: err.Error()}
			util.SendResponse(w, results)
			return
		}

		type RequestBody struct {
			Income string `json:"income"`
		}

		if len(body) != 0 {
			var reqBody RequestBody
			err = json.Unmarshal(body, &reqBody)
			if err != nil {
				results["error"] = Error{Message: err.Error()}
				util.SendResponse(w, results)
				return

			}

			//get income from request body and convert into float
			currentIncome, err := strconv.ParseFloat(reqBody.Income, 64)
			if err != nil {
				results["error"] = Error{Message: err.Error()}
				util.SendResponse(w, results)
				return

			} else {
				//calculate the tax based on the income passed
				results = util.CalculateTax(results, currentIncome)

			}
		}

	}
	//Send Response
	util.SendResponse(w, results)
}


/*
This is a function that handles HTTP requests using a mux router.
*/
func handleRequests() {
	//creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	//replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/tax-calculator", calculator).Methods(http.MethodGet)
	myRouter.HandleFunc("/tax-calculator/tax-year/{tax-year}",tax).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

/*
The main function is the entry point of the Go program. Here, we are calling the handleRequests() function to start the web server and listen for incoming requests.
*/
func main() {

	handleRequests()
}
