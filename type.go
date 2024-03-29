package alpacaio

import "time"

type Bar struct {
	Timestamp time.Time `json:"t"`
	Open      float32   `json:"o"`
	High      float32   `json:"h"`
	Low       float32   `json:"l"`
	Close     float32   `json:"c"`
	Volume    int64     `json:"v"`
}

type Bars []Bar

type StockBarsResponse struct {
	Symbol    string `json:"symbol"`
	PageToken string `json:"next_page_token"`
	Bars      Bars   `json:"bars"`
}

type Timespan string

const (
	Minute Timespan = "1Min"
	Hour   Timespan = "1Hour"
	Day    Timespan = "1Day"
)

type RequestOptions struct {
	PageToken string `url:"page_token,omitempty"`
	Limit     int32  `url:"limit,omitempty"`
	Start     string `url:"start,omitempty"`
	End       string `url:"end,omitempty"`
	Timeframe string `url:"timeframe,omitempty"`
}

type StreamingServerMsg struct {
	Event     string    `json:"T"`
	Symbol    string    `json:"S"`
	TradeID   int64     `json:"i"`
	Exchange  string    `json:"x"`
	Price     float32   `json:"p"`
	Size      int32     `json:"s"`
	Timestamp time.Time `json:"t"`
	Tape      string    `json:"z"`
	// Quote
	BidExchange string  `json:"bx"`
	AskExchange string  `json:"ax"`
	BidPrice    float32 `json:"bp"`
	AskPrice    float32 `json:"ap"`
	BidSize     int32   `json:"bs"`
	AskSize     int32   `json:"as"`
	// Bar
	Volume int32       `json:"v"`
	Open   float32     `json:"o"`
	High   float32     `json:"h"`
	Low    float32     `json:"l"`
	C      interface{} `json:"c"`
	// Error
	Message string `json:"msg"`
	Code    int32  `json:"code"`
}

//easyjson:json
type StreamingServerMsges []StreamingServerMsg

type Trade struct {
	Event      string    `json:"T"`
	Symbol     string    `json:"S"`
	TradeID    int64     `json:"i"`
	Exchange   string    `json:"x"`
	Price      float32   `json:"p"`
	Size       int32     `json:"s"`
	Timestamp  time.Time `json:"t"`
	Conditions []string  `json:"c"`
	Tape       string    `json:"z"`
}

//easyjson:json
type Trades []Trade

type StockTradesResponse struct {
	Symbol    string `json:"symbol"`
	PageToken string `json:"next_page_token"`
	Trades    Trades `json:"trades"`
}

type Quote struct {
	Event       string    `json:"T"`
	Symbol      string    `json:"S"`
	Condition   []string  `json:"c"`
	BidExchange string    `json:"bx"`
	AskExchange string    `json:"ax"`
	BidPrice    float32   `json:"bp"`
	AskPrice    float32   `json:"ap"`
	BidSize     int32     `json:"bs"`
	AskSize     int32     `json:"as"`
	Timestamp   time.Time `json:"t"`
	Tape        string    `json:"z"`
}

//easyjson:json
type Quotes []Quote

type StockQuotesResponse struct {
	Symbol    string `json:"symbol"`
	PageToken string `json:"next_page_token"`
	Quotes    Quotes `json:"quotes"`
}

type AggregatedBar struct {
	Event     string    `json:"T"`
	Symbol    string    `json:"S"`
	Volume    int32     `json:"v"`
	Open      float32   `json:"o"`
	High      float32   `json:"h"`
	Low       float32   `json:"l"`
	Close     float32   `json:"c"`
	Timestamp time.Time `json:"s"`
}

//easyjson:json
type AggregatedBars []AggregatedBar
