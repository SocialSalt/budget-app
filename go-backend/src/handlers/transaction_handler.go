package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-backend/src/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTransaction(w, r)
	case http.MethodPost:
		PostTransaction(w, r)
	default:
		http.Error(w, "Unknown Method", http.StatusBadRequest)
	}
}

func parseDate(rawDate string) (time.Time, error) {
	splitDate := strings.Split(rawDate, "/")
	for k, x := range splitDate {
		if len(x) == 1 {
			splitDate[k] = "0" + x
		}
	}
	rawDate = strings.Join(splitDate, "/")

	form := "01/02/2006"
	d, err := time.Parse(form, rawDate)
	if err != nil {
		return time.Time{}, err
	}
	return d, nil
}

func csvToTransactions(data [][]string) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, len(data))
	headers := data[0]
	for i, line := range data {
		if i == 0 {
			continue
		}
		t := models.Transaction{}
		for j, header := range headers {
			switch header {
			case "Date":
				d, err := parseDate(line[j])
				if err != nil {
					return nil, err
				}
				t.Date = d
				t.Month = d.Month()
				_, t.Week = d.ISOWeek()
			case "Company":
				t.Company = line[j]
			case "Category":
				t.Category = line[j]
			case "Amount":
				s := strings.ReplaceAll(line[j], ",", "")
				s = strings.ReplaceAll(s, "$", "")
				a, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return nil, err
				}
				t.Amount = a
			case "AccountNumber":
				t.AccountNumber = line[j]
			case "Institution":
				t.Institution = line[j]
			case "CheckNumber":
				t.CheckNumber = line[j]
			case "FullDescription":
				t.FullDescription = line[j]
			}
			t.DateAdded = time.Now()
		}
		transactions[i] = t
	}
	return transactions, nil
}

func UploadTransaction(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)     // limit max input length
	file, _, err := r.FormFile("file") // "file" is the key set in the request
	if err != nil {
		log.Print("Failed to read file from request: ", err)
		http.Error(w, "Failed to read file from request", http.StatusInternalServerError)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Print("Failed to parse csv file: ", err)
		http.Error(w, "Failed to parse csv file", http.StatusBadRequest)
	}

	transactions, err := csvToTransactions(data)
	if err != nil {
		log.Print("Failed to parse csv file: ", err)
		http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
	}
	// fmt.Print(transactions)
	models.DB.Create(&transactions)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	models.DB.Find(&transactions)
	b, err := json.Marshal(transactions)
	if err != nil {
		log.Print("Failed to convert transactions: ", err)
		http.Error(w, "Failed to convert transactions", http.StatusInternalServerError)
	}
	w.Write(b)
}

func PostTransaction(http.ResponseWriter, *http.Request) {

}
