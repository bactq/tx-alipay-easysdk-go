package common

type AlipayDataDataserviceBillDownloadurlQueryResponse struct {
	HttpBody        string `json:"http_body"`
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	SubCode         string `json:"sub_code"`
	SubMsg          string `json:"sub_msg"`
	BillDownloadUrl string `json:"bill_download_url"`
}

type AlipayTradeCancelResponse struct {
	HttpBody           string `json:"http_body"`
	Code               string `json:"code"`
	Msg                string `json:"msg"`
	SubCode            string `json:"sub_code"`
	SubMsg             string `json:"sub_msg"`
	TradeNo            string `json:"trade_no"`
	OutTradeNo         string `json:"out_trade_no"`
	RetryFlag          string `json:"retry_flag"`
	Action             string `json:"action"`
	GmtRefundPay       string `json:"gmt_refund_pay"`
	RefundSettlementId string `json:"refund_settlement_id"`
}

type AlipayTradeCloseResponse struct {
	HttpBody   string `json:"http_body"`
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"sub_code"`
	SubMsg     string `json:"sub_msg"`
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}

type AlipayTradeCreateResponse struct {
	HttpBody   string `json:"http_body"`
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"sub_code"`
	SubMsg     string `json:"sub_msg"`
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}
type TradeFundBill struct {
	FundChannel string `json:"fund_channel"`
	BankCode    string `json:"bank_code"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
	FundType    string `json:"fund_type"`
}

type RefundRoyaltyResult struct {
	RefundAmount  string `json:"refund_amount"`
	RoyaltyType   string `json:"royalty_type"`
	ResultCode    string `json:"result_code"`
	TransOut      string `json:"trans_out"`
	TransOutEmail string `json:"trans_out_email"`
	TransIn       string `json:"trans_in"`
	TransInEmail  string `json:"trans_in_email"`
}
type AlipayTradeFastpayRefundQueryResponse struct {
	HttpBody                     string                 `json:"http_body"`
	Code                         string                 `json:"code"`
	Msg                          string                 `json:"msg"`
	SubCode                      string                 `json:"sub_code"`
	SubMsg                       string                 `json:"sub_msg"`
	ErrorCode                    string                 `json:"error_code"`
	GmtRefundPay                 string                 `json:"gmt_refund_pay"`
	IndustrySepcDetail           string                 `json:"industry_sepc_detail"`
	OutRequestNo                 string                 `json:"out_request_no"`
	OutTradeNo                   string                 `json:"out_trade_no"`
	PresentRefundBuyerAmount     string                 `json:"present_refund_buyer_amount"`
	PresentRefundDiscountAmount  string                 `json:"present_refund_discount_amount"`
	PresentRefundMdiscountAmount string                 `json:"present_refund_mdiscount_amount"`
	RefundAmount                 string                 `json:"refund_amount"`
	RefundChargeAmount           string                 `json:"refund_charge_amount"`
	RefundDetailItemList         []*TradeFundBill       `json:"refund_detail_item_list"`
	RefundReason                 string                 `json:"refund_reason"`
	RefundRoyaltys               []*RefundRoyaltyResult `json:"refund_royaltys"`
	RefundSettlementId           string                 `json:"refund_settlement_id"`
	RefundStatus                 string                 `json:"refund_status"`
	SendBackFee                  string                 `json:"send_back_fee"`
	TotalAmount                  string                 `json:"total_amount"`
	TradeNo                      string                 `json:"trade_no"`
}
type TradeSettleDetail struct {
	OperationType string `json:"operation_type"`

	OperationSerial_no string `json:"operation_serial_no"`

	OperationDt string `json:"operation_dt"`

	TransOut string `json:"trans_out"`

	TransIn string `json:"trans_in"`

	Amount string `json:"amount"`
}
type TradeSettleInfo struct {
	TradeSettleDetailList []*TradeSettleDetail `json:"trade_settle_detail_list"`
}
type AlipayTradeQueryResponse struct {
	// 响应原始字符串

	HttpBody string `json:"http_body"`

	Code string `json:"code"`

	Msg string `json:"msg"`

	SubCode string `json:"sub_code"`

	SubMsg string `json:"sub_msg"`

	TradeNo string `json:"trade_no"`

	OutTradeNo string `json:"out_trade_no"`

	BuyerLogonId string `json:"buyer_logon_id"`

	TradeStatus string `json:"trade_status"`

	TotalAmount string `json:"total_amount"`

	TransCurrency string `json:"trans_currency"`

	SettleCurrency string `json:"settle_currency"`

	SettleAmount string `json:"settle_amount"`

	PayCurrency string `json:"pay_currency"`

	PayAmount string `json:"pay_amount"`

	SettleTransRate string `json:"settle_trans_rate"`

	TransPayRate string `json:"trans_pay_rate"`

	BuyerPayAmount string `json:"buyer_pay_amount"`

	PointAmount string `json:"point_amount"`

	InvoiceAmount string `json:"invoice_amount"`

	SendPayDate string `json:"send_pay_date"`

	ReceiptAmount string `json:"receipt_amount"`

	StoreId string `json:"store_id"`

	TerminalId string `json:"terminal_id"`

	FundBillList []*TradeFundBill `json:"fund_bill_list"`

	StoreName string `json:"store_name"`

	BuyerUserId string `json:"buyer_user_id"`

	ChargeAmount string `json:"charge_amount"`

	ChargeFlags string `json:"charge_flags"`

	SettlementId string `json:"settlement_id"`

	TradeSettleInfo []*TradeSettleInfo `json:"trade_settle_info"`

	AuthTradePayMode string `json:"auth_trade_pay_mode"`

	BuyerUserType string `json:"buyer_user_type"`

	MdiscountAmount string `json:"mdiscount_amount"`

	DiscountAmount string `json:"discount_amount"`

	BuyerUserName string `json:"buyer_user_name"`

	Subject string `json:"subject"`

	Body string `json:"body"`

	AlipaySubMerchantId string `json:"alipay_sub_merchant_id"`

	ExtInfos string `json:"ext_infos"`
}
type PresetPayToolInfo struct {
	Amount []string `json:"amount"`

	AssertTypeCode string `json:"assert_type_code"`
}
type AlipayTradeRefundResponse struct {
	// 响应原始字符串
	HttpBody string `json:"http_body"`

	Code string `json:"code"`

	Msg string `json:"msg"`

	SubCode string `json:"sub_code"`

	SubMsg string `json:"sub_msg"`

	TradeNo string `json:"trade_no"`

	OutTradeNo string `json:"out_trade_no"`

	BuyerLogonId string `json:"buyer_logon_id"`

	FundChange string `json:"fund_change"`

	RefundFee string `json:"refund_fee"`

	RefundCurrency string `json:"refund_currency"`

	GmtRefundPay string `json:"gmt_refund_pay"`

	RefundDetailItemList []*TradeFundBill `json:"refund_detail_item_list"`

	StoreName string `json:"store_name"`

	BuyerUserId string `json:"buyer_user_id"`

	RefundPresetPaytoolList []*PresetPayToolInfo `json:"refund_preset_paytool_list"`

	RefundSettlementId string `json:"refund_settlement_id"`

	PresentRefundBuyerAmount string `json:"present_refund_buyer_amount"`

	PresentRefundDiscountAmount string `json:"present_refund_discount_amount"`

	PresentRefundMdiscountAmount string `json:"present_refund_mdiscount_amount"`
}
