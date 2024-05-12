package coinRepository

import (
	"coffeshop/database"
	"coffeshop/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CoinResponse struct {
	Data struct {
		ID                string `json:"id"`
		Rank              string `json:"rank"`
		Symbol            string `json:"symbol"`
		Name              string `json:"name"`
		Supply            string `json:"supply"`
		MaxSupply         string `json:"maxSupply"`
		MarketCapUsd      string `json:"marketCapUsd"`
		VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
		PriceUsd          string `json:"priceUsd"`
		ChangePercent24Hr string `json:"changePercent24Hr"`
		Vwap24Hr          string `json:"vwap24Hr"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type ListCoinResponse struct {
	Data []struct {
		ID                string `json:"id"`
		Rank              string `json:"rank"`
		Symbol            string `json:"symbol"`
		Name              string `json:"name"`
		Supply            string `json:"supply"`
		MaxSupply         string `json:"maxSupply"`
		MarketCapUsd      string `json:"marketCapUsd"`
		VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
		PriceUsd          string `json:"priceUsd"`
		ChangePercent24Hr string `json:"changePercent24Hr"`
		Vwap24Hr          string `json:"vwap24Hr"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type RatesResponse struct {
	Data []struct {
		ID             string `json:"id"`
		Symbol         string `json:"symbol"`
		CurrencySymbol string `json:"currencySymbol"`
		Type           string `json:"type"`
		RateUsd        string `json:"rateUsd"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func GetAllCoins() (ListCoinResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get("https://api.coincap.io/v2/assets")

	if err != nil {
		return ListCoinResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ListCoinResponse{}, errors.New("error retrieve coin from coincap")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return ListCoinResponse{}, err
	}

	var responseCoin ListCoinResponse

	err = json.Unmarshal(body, &responseCoin)

	return responseCoin, err
}

func GetAllCoinRates() (RatesResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get("https://api.coincap.io/v2/rates")

	if err != nil || resp.StatusCode != http.StatusOK {
		return RatesResponse{}, errors.New("fail get rate")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return RatesResponse{}, errors.New("fail read rate response")
	}

	var responseRate RatesResponse

	err = json.Unmarshal(body, &responseRate)

	return responseRate, err
}

func GetTrackedCoinByUserIdAndCoinId(userId int, coinId string) (models.UserCoin, error) {

	db := database.Client

	row := db.QueryRow("SELECT * FROM USERCOINS WHERE userId=? and coinId=?", userId, coinId)

	var retrievedUserCoin models.UserCoin
	if err := row.Scan(&retrievedUserCoin.UserId, &retrievedUserCoin.CoinId); err != nil {
		return retrievedUserCoin, errors.New("error retrieve user coin")
	}

	return retrievedUserCoin, nil
}

func GetCoinById(coinId string) (CoinResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	var coinResponse CoinResponse

	resp, err := client.Get(fmt.Sprintf("https://api.coincap.io/v2/assets/%s", coinId))

	if err != nil || resp.StatusCode != http.StatusOK {
		return CoinResponse{}, errors.New("error get coin by id")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return CoinResponse{}, errors.New("error read all")
	}

	err = json.Unmarshal(body, &coinResponse)
	if err != nil {
		return CoinResponse{}, errors.New("error map to coinresponse")
	}

	return coinResponse, nil
}

func AddTrackedCoin(userId int, coinId string) (models.UserCoin, error) {

	db := database.Client
	tx, err := db.Begin()

	if err != nil {
		return models.UserCoin{}, errors.New("error begin transaction")
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO USERCOINS (id, userId, coinId) VALUES (?, ?, ?)")
	stmt.Exec(nil, userId, coinId)

	err = tx.Commit()

	if err != nil {
		return models.UserCoin{}, errors.New("error insert user coin")
	}

	defer stmt.Close()

	return models.UserCoin{UserId: userId, CoinId: coinId}, nil

}

func DeleteUserCoin(userId int, coinId string) error {
	db := database.Client
	tx, err := db.Begin()

	if err != nil {
		return errors.New("error begin db transacition")
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("DELETE FROM USERCOINS WHERE userid=? and coinId=?")
	stmt.Exec(userId, coinId)

	err = tx.Commit()

	if err != nil {
		return errors.New("error delete from db")
	}

	defer stmt.Close()

	return nil
}

func GetTrackedCoinId(userId int) ([]string, error) {
	db := database.Client
	rows, err := db.Query("SELECT coinId FROM USERCOINS WHERE userId=?", userId)
	coinIds := make([]string, 0)
	if err != nil {
		return []string{}, errors.New("error get coin id from db")
	}

	for rows.Next() {
		var coinId string
		err = rows.Scan(&coinId)

		if err != nil {
			return []string{}, errors.New("error scan query result")
		}

		coinIds = append(coinIds, coinId)

	}

	return coinIds, nil
}
