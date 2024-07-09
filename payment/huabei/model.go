package huabei

type AlipayTradeCreateResponse struct {
	HttpBody   string `json:"http_body,omitempty"`
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

type HuabeiConfig struct {
	HbFqNum           string `json:"hb_fq_num,omitempty"`
	HbFqSellerPercent string `json:"hb_fq_seller_percent,omitempty"`
}
