package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

type Bitcoin struct {
	gorm.Model
	Value      float64   `gorm:"not null"`
	LastUpdate time.Time `gorm:"not null;type:timestamp"`
}

// CoinMarket struct

type Status struct {
	Timestamp    time.Time   `json:"timestamp"`
	ErrorCode    int         `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int         `json:"elapsed"`
	CreditCount  int         `json:"credit_count"`
}

type Data struct {
	ID          int       `json:"id"`
	Symbol      string    `json:"symbol"`
	Name        string    `json:"name"`
	Amount      int       `json:"amount"`
	LastUpdated time.Time `json:"last_updated"`
	Quote       *Quote    `json:"quote"`
}

type Quote struct {
	Btc *BTC `json:"BTC"`
}

type BTC struct {
	Price       float64   `json:"price"`
	LastUpdated time.Time `json:"last_updated"`
}

type CoinMarket struct {
	Status *Status `json:"status"`
	Data   *Data   `json:"data"`
}

func (bitcoin *Bitcoin) GetValueBTC(option int) *Bitcoin {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/tools/price-conversion", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("id", "2783")
	q.Add("amount", "1")
	q.Add("convert", "BTC")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("TOKEN_COINMARKET"))

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao atualizar valor")
		return nil
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	response := CoinMarket{}
	err = json.Unmarshal([]byte(respBody), &response)
	if err != nil {
		fmt.Println("Erro ao atualizar valor")
		return nil
	}

	bitcoin.Value = response.Data.Quote.Btc.Price
	bitcoin.LastUpdate = time.Now()

	if option == 1 {
		fmt.Println("Creating bitcoin register")
		GetDB().Create(bitcoin)
		return bitcoin
	}

	fmt.Println("Updating bitcoin register")
	GetDB().Save(bitcoin)
	return bitcoin
}
