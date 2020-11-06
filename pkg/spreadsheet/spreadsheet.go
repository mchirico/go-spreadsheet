package spreadsheet

import (
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Spreadsheet struct {
	CredentialsFile string
	TokenFile       string
	SpreadsheetID   string

}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func NewSP() *Spreadsheet {
	sp := &Spreadsheet{
		CredentialsFile: "../../credentials/credentials.json",
		TokenFile:       "../../credentials/token.json",
	}
	return sp
}

func (sp *Spreadsheet) GetClient() (*http.Client, error) {

	token, err := tokenFromFile(sp.TokenFile)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(sp.CredentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config.Client(context.Background(), token), err
}
