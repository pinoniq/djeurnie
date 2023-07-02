package data

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type bitPandaCandlestick struct {
	High        float64 `json:"high,string"`
	Low         float64 `json:"low,string"`
	Open        float64 `json:"open,string"`
	Close       float64 `json:"close,string"`
	TotalAmount float64 `json:"total_amount,string"`
	Volume      float64 `json:"volume,string"`
}

type binanceCandlestick struct {
	OpenTime                 int64
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	Trades                   float64
	QuoteAssetVolume         float64
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
}

func (c *binanceCandlestick) ToSlice() []string {
	return []string{
		strconv.FormatInt(c.OpenTime, 10),
		strconv.FormatFloat(c.Open, 'f', 6, 64),
		strconv.FormatFloat(c.High, 'f', 6, 64),
		strconv.FormatFloat(c.Low, 'f', 6, 64),
		strconv.FormatFloat(c.Close, 'f', 6, 64),
		strconv.FormatFloat(c.Volume, 'f', 6, 64),
		strconv.FormatFloat(c.Trades, 'f', 6, 64),
		strconv.FormatFloat(c.QuoteAssetVolume, 'f', 6, 64),
		strconv.FormatFloat(c.TakerBuyBaseAssetVolume, 'f', 6, 64),
		strconv.FormatFloat(c.TakerBuyQuoteAssetVolume, 'f', 6, 64),
	}
}

func (c *binanceCandlestick) UnmarshalJSON(data []byte) error {
	var tmp []interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	c.OpenTime = int64(tmp[0].(float64))
	c.Open, _ = strconv.ParseFloat(tmp[1].(string), 64)
	c.High, _ = strconv.ParseFloat(tmp[2].(string), 64)
	c.Low, _ = strconv.ParseFloat(tmp[3].(string), 64)
	c.Close, _ = strconv.ParseFloat(tmp[4].(string), 64)
	c.Volume, _ = strconv.ParseFloat(tmp[5].(string), 64)
	c.Trades = tmp[8].(float64)
	c.QuoteAssetVolume, _ = strconv.ParseFloat(tmp[7].(string), 64)
	c.TakerBuyBaseAssetVolume, _ = strconv.ParseFloat(tmp[9].(string), 64)
	c.TakerBuyQuoteAssetVolume, _ = strconv.ParseFloat(tmp[10].(string), 64)

	return nil
}

func CreateData(from time.Time, to time.Time) {
	fmt.Println("Retrieving data from " + from.Format(time.DateTime) + " to " + to.Format(time.DateTime))
	queryParams := url.Values{
		"symbol":    {"BTCEUR"},
		"interval":  {"1m"},
		"startTime": {strconv.FormatInt(from.UnixMilli(), 10)},
		"endTime":   {strconv.FormatInt(to.UnixMilli(), 10)},
	}
	endpoint := url.URL{
		Scheme:   "https",
		Host:     "data-api.binance.vision",
		Path:     "/api/v3/klines",
		RawQuery: queryParams.Encode(),
	}

	rawJson, err := retrieveJsonFromUrl(endpoint.String())
	if err != nil {
		panic(err)
	}

	var sticks []binanceCandlestick
	err = json.Unmarshal(rawJson, &sticks)
	if err != nil {
		fmt.Println(string(rawJson))
		panic(err)
	}

	fmt.Println("Writing data to data/train_minutes.csv")
	fileName := "data/train_minutes.csv"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := range sticks {
		err := writer.Write(sticks[i].ToSlice())
		if err != nil {
			panic(err)
		}
	}

}

func retrieveJsonFromUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
