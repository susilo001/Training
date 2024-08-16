package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/xuri/excelize/v2"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func main() {
	// Set up logging to a .log file
	logFile, err := os.OpenFile("process.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting process to fetch user data...")

	url := "https://jsonplaceholder.typicode.com/users"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var users []User
	err = json.Unmarshal(responseData, &users)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Create CSV file
	log.Println("Creating users.csv...")
	file, err := os.Create("users.csv")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"ID", "Name", "Username", "Email", "Address", "Phone", "Website", "Company"}
	err = writer.Write(header)
	if err != nil {
		log.Fatalf("Error writing header: %v", err)
	}

	// Create an Excel file
	log.Println("Creating users.xlsx...")
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.NewSheet(sheet)

	// Write Excel header
	for i, h := range header {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Write CSV and Excel rows
	for rowIndex, user := range users {
		// Format the address as a single string
		address := fmt.Sprintf("%s, %s, %s, %s", user.Address.Street, user.Address.Suite, user.Address.City, user.Address.Zipcode)

		row := []string{
			fmt.Sprintf("%d", user.ID),
			user.Name,
			user.Username,
			user.Email,
			address,
			user.Phone,
			user.Website,
			user.Company.Name,
		}

		// Write to CSV
		err = writer.Write(row)
		if err != nil {
			log.Fatalf("Error writing row to CSV: %v", err)
		}

		// Write to Excel
		for colIndex, cellValue := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2) // Excel rows start at 1, headers occupy the first row
			f.SetCellValue(sheet, cell, cellValue)
		}
	}

	// Save the Excel file
	err = f.SaveAs("users.xlsx")
	if err != nil {
		log.Fatalf("Error saving Excel file: %v", err)
	}

	log.Println("Process completed successfully!")
	fmt.Println("Files created: users.csv, users.xlsx, and process.log")
}
