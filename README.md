# My Tax Calculator API

My Tax Calculator API calculates tax based on the income and tax year.

## Installation

To install Tax Calculator API, you will need to have Go installed on your machine. You can download and install Go from the official website: https://golang.org/dl/.

Once you have Go installed, you can clone the repository:

```
git clone https://github.com/mailgagannow/go-assignment.git
```

## Usage

To start the server, navigate to the project directory and run the following command:

```
go run main.go
```

This will start the server on port 10000. You can access the API by navigating to http://localhost:10000 in your web browser.

### Endpoints

The following endpoints are available:

- `GET/tax-calculator`: Returns Tax Brackets For Year 2022.
- `POST/tax-calculator/tax-year/{tax-year}`: Return the tax brackets for the year passed/calculate tax if the income in passed the post parameter.

### Sample Request

1-Please find a sample request at [/tax-calculator/tax-year](http://localhost:10000/tax-calculator/tax-year/2022), 

Parameters:
tax-year (required): The tax year to return the year brackets and use for the tax calculation.

It returns the following JSON response: 

```json
{
    "tax_brackets": [
        {
            "min": 0,
            "max": 50197,
            "rate": 0.15
        },
        {
            "min": 50197,
            "max": 100392,
            "rate": 0.205
        },
        {
            "min": 100392,
            "max": 155625,
            "rate": 0.26
        },
        {
            "min": 155625,
            "max": 221708,
            "rate": 0.29
        },
        {
            "min": 221708,
            "rate": 0.33
        }
    ]
}

```
Request body (income:optional):

```json
{
    "income":"75000"
}
```
Response

```json
{
    "effectiveTaxRate": 17.01,
    "taxAmount": 12755.35,
    "taxPerSlab": [
        {
            "max": 47630,
            "min": 0,
            "rate": 0.15,
            "tax": 7144.5,
            "taxable_income": 47630
        },
        {
            "max": 95259,
            "min": 47630,
            "rate": 0.205,
            "tax": 5610.85,
            "taxable_income": 27370
        }
    ],
    "tax_brackets": [
        {
            "min": 0,
            "max": 47630,
            "rate": 0.15
        },
        {
            "min": 47630,
            "max": 95259,
            "rate": 0.205
        },
        {
            "min": 95259,
            "max": 147667,
            "rate": 0.26
        },
        {
            "min": 147667,
            "max": 210371,
            "rate": 0.29
        },
        {
            "min": 210371,
            "rate": 0.33
        }
    ],
    "totalIncome": 75000
}

```

## Testing

To run the tests for My Go API, navigate to the project directory and run the following command:

```
go test ./...
```

This will run all of the tests in the project.
