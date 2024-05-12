package coinService

import (
	"coffeshop/models"
	coinRepository "coffeshop/repositories/CoinRepository"
	"errors"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type CoinDTO struct {
	ID       string
	Name     string
	Supply   string
	PriceIdr float64
}

type TrackCoinRequestDTO struct {
	UserId int
	CoinId string
}

func GetAllCoins() ([]CoinDTO, error) {

	allCoins, err := coinRepository.GetAllCoins()
	var rupiahRate float64

	if err != nil {
		return []CoinDTO{}, err
	}

	coinRates, err := coinRepository.GetAllCoinRates()

	if err != nil {
		return []CoinDTO{}, err
	}

	for _, v := range coinRates.Data {
		if v.ID == "indonesian-rupiah" {
			rupiahRate, _ = strconv.ParseFloat(v.RateUsd, 64)
			break
		}
	}

	var listCoin []CoinDTO

	for _, v := range allCoins.Data {
		usdRate, _ := strconv.ParseFloat(v.PriceUsd, 64)
		coin := CoinDTO{
			ID:       v.ID,
			Name:     v.Name,
			PriceIdr: usdRate / rupiahRate,
		}
		listCoin = append(listCoin, coin)
	}

	return listCoin, nil
}

func TrackCoin(ctx *gin.Context) (models.UserCoin, error) {
	var trackCointRequest TrackCoinRequestDTO

	if err := ctx.BindJSON(&trackCointRequest); err != nil {
		return models.UserCoin{}, errors.New("error mapping json")
	}

	retrievedUserCoin, err := coinRepository.GetTrackedCoinByUserIdAndCoinId(trackCointRequest.UserId, trackCointRequest.CoinId)

	if err != nil {
		return models.UserCoin{}, err
	}

	if retrievedUserCoin != (models.UserCoin{}) {
		return retrievedUserCoin, nil
	}

	coin, err := coinRepository.GetCoinById(trackCointRequest.CoinId)

	if err != nil {
		return models.UserCoin{}, err
	}

	if coin == (coinRepository.CoinResponse{}) {
		return models.UserCoin{}, errors.New("coin not exists")
	}

	userCoin, err := coinRepository.AddTrackedCoin(trackCointRequest.UserId, trackCointRequest.CoinId)
	if err != nil {
		return models.UserCoin{}, err
	}
	return userCoin, nil

}

func UntrackCoin(ctx *gin.Context) (bool, error) {
	var trackCointRequest TrackCoinRequestDTO

	if err := ctx.BindJSON(&trackCointRequest); err != nil {
		return false, errors.New("error bind trackcoin request")
	}

	if err := coinRepository.DeleteUserCoin(trackCointRequest.UserId, trackCointRequest.CoinId); err != nil {
		return false, err
	}

	return true, nil

}

func GetTrackedCoins(ctx *gin.Context) ([]CoinDTO, error) {

	userId, err := strconv.Atoi(ctx.Params.ByName("userId"))
	var rupiahRate float64

	if err != nil {
		return []CoinDTO{}, errors.New("error parse user id")
	}

	coinIds, err := coinRepository.GetTrackedCoinId(userId)

	if err != nil {
		return []CoinDTO{}, err
	}

	rates, err := coinRepository.GetAllCoinRates()

	if err != nil {
		return []CoinDTO{}, err
	}

	for _, v := range rates.Data {
		if v.ID == "indonesian-rupiah" {
			rupiahRate, _ = strconv.ParseFloat(v.RateUsd, 64)
			break
		}
	}

	var userCoins []CoinDTO

	var wg sync.WaitGroup

	for _, v := range coinIds {

		wg.Add(1)
		go func(coinId string) {
			defer wg.Done()

			coinResponse, _ := coinRepository.GetCoinById(coinId)

			usdRate, _ := strconv.ParseFloat(coinResponse.Data.PriceUsd, 64)

			coinDTO := CoinDTO{
				ID:       coinResponse.Data.ID,
				PriceIdr: usdRate / rupiahRate,
				Name:     coinResponse.Data.Name,
			}

			userCoins = append(userCoins, coinDTO)

		}(v)

	}
	wg.Wait()

	return userCoins, nil

}
