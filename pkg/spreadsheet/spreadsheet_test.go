package spreadsheet

import (
	"fmt"
	"google.golang.org/api/sheets/v4"
	"log"
	"testing"
)

func TestSpreadsheet_GetClient(t *testing.T) {

	sp := NewSP()
	client, err := sp.GetClient()
	if err != nil {
		t.Fatalf("error GetClient")
	}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetID := "1zghcqhde3w15UmAgiXasoCT30-Z3s16_TlN7ri0Fer0"
	readRange := "Sheet1!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}

	writeRange := "A1"

	var vr sheets.ValueRange

	myval := []interface{}{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	vr.Values = append(vr.Values, myval)

	_, err = srv.Spreadsheets.Values.Update(spreadsheetID, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

}
