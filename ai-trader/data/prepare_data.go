package data

type TrainingMinute struct {
	HighDiff                     float64
	LowDiff                      float64
	Volume                       float64
	Trades                       float64
	TradesDiff                   float64
	TakerBuyBaseAssetVolume      float64
	TakerBuyBaseAssetVolumeDiff  float64
	TakerBuyQuoteAssetVolume     float64
	TakerBuyQuoteAssetVolumeDiff float64
}

// 1 minute will have 9 data-points. Diffs are calculated compared to previous minute.
// 1hour will thus have 540 data-points.
// This seems like a good first-start for training.
// This does however mean our data-model's input layer has 540 nodes.
// We will use 2/3's nodes for our hidden layer. Being 360 nodes.
