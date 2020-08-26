package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
)

type ObjectInfoStats struct {
	Timestamp                     int64   `json:"timestamp"`
	MarketPriceUsd                float64 `json:"market_price_usd"`
	HashRate                      float64 `json:"hash_rate"`
	TotalFeesBtc                  int64   `json:"total_fees_btc"`
	NBtcMined                     int64   `json:"n_btc_mined"`
	NTx                           int8     `json:"n_tx"`
	NBlocksMined                  int8     `json:"n_blocks_mined"`
	MinutesBetweenBlocks          float64 `json:"minutes_between_blocks"`
	Totalbc                       int64   `json:"totalbc"`
	NBlocksTotal                  int8     `json:"n_blocks_total"`
	EstimatedTransactionVolumeUsd float64 `json:"estimated_transaction_volume_usd"`
	BlocksSize                    int8     `json:"blocks_size"`
	MinersRevenueUsd              float64 `json:"miners_revenue_usd"`
	Nextretarget                  int8     `json:"nextretarget"`
	Difficulty                    int64   `json:"difficulty"`
	EstimatedBtcSent              int64   `json:"estimated_btc_sent"`
	MinersRevenueBtc              int8     `json:"miners_revenue_btc"`
	TotalBtcSent                  int64   `json:"total_btc_sent"`
	TradeVolumeBtc                float64 `json:"trade_volume_btc"`
	TradeVolumeUsd                float64 `json:"trade_volume_usd"`
}

func main() {
	for i := 0; i < 5; i++ {
	data := ObjectInfoStats {
		Timestamp: gofakeit.Int64(),
		MarketPriceUsd: gofakeit.Float64(),
		HashRate: gofakeit.Float64(),
		TotalFeesBtc: gofakeit.Int64(),
		NBtcMined: gofakeit.Int64(),
		NTx: gofakeit.Int8(),
		NBlocksMined: gofakeit.Int8(),
		MinutesBetweenBlocks: gofakeit.Float64(),
		Totalbc: gofakeit.Int64(),
		NBlocksTotal: gofakeit.Int8(),
		EstimatedTransactionVolumeUsd: gofakeit.Float64(),
		BlocksSize:  gofakeit.Int8(),
		MinersRevenueUsd: gofakeit.Float64(),
		Nextretarget: gofakeit.Int8(),
		Difficulty: gofakeit.Int64(),
		EstimatedBtcSent: gofakeit.Int64(),
		MinersRevenueBtc: gofakeit.Int8(),
		TotalBtcSent: gofakeit.Int64(),
		TradeVolumeBtc: gofakeit.Float64(),
		TradeVolumeUsd: gofakeit.Float64(),
	}
	file, _ := json.MarshalIndent(data, "", " ")

	concatenated := fmt.Sprintf("%d_Stats.json",i)
	_ = ioutil.WriteFile(concatenated, file, 0644)
	}
}
