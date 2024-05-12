package controllers

import (
	coinService "coffeshop/services/coin"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

type TrackCoinRequestDTO struct {
	UserId int
	CoinId string
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

type CoinDTO struct {
	ID       string
	Name     string
	Supply   string
	PriceIdr float64
}

func GetAllCoins() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		listCoin, err := coinService.GetAllCoins()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": listCoin,
		})

	}

}

func TrackCoin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userCoin, err := coinService.TrackCoin(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": userCoin})
	}
}

func UntrackCoin(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := coinService.UntrackCoin(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetMyCoin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userCoins, _ := coinService.GetTrackedCoins(ctx)

		ctx.JSON(http.StatusOK, gin.H{"data": userCoins})
	}
}
