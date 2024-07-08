package common

type AlipayTradeCreateResponse struct {
	HttpBody   string `json:"http_body,omitempty"`
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}
type AlipayTradeQueryResponse struct {
	HttpBody   string `json:"http_body,omitempty"`
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}
