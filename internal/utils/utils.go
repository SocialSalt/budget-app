package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/socialsalt/budget-app/internal/model"
)

func ParseDate(rawDate string) (time.Time, error) {
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

func ParseTransactionCSV(data [][]string) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, len(data))
	headers := data[0]
	for i, line := range data {
		if i == 0 {
			continue
		}
		t := model.Transaction{}
		for j, header := range headers {
			switch header {
			case "Date":
				d, err := ParseDate(line[j])
				if err != nil {
					return nil, err
				}
				t.Date = d
			case "Company":
				t.Company = line[j]
			case "Category":
				t.Category = line[j]
			case "Amount":
				s := strings.ReplaceAll(line[j], ",", "")
				s = strings.ReplaceAll(s, "$", "")
				s = strings.ReplaceAll(s, ".", "")
				a, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return nil, err
				}
				t.Amount = a
			case "AccountNumber":
				t.AccountNumber = line[j]
			case "Institution":
				t.Institution = line[j]
			case "FullDescription":
				t.FullDescription = line[j]
			}
			t.DateAdded = time.Now()
		}
		transactions[i] = t
	}
	return transactions, nil
}
