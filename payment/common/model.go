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

type AlipayTradeCloseResponse struct {
	HttpBody   string `json:"http_body,omitempty"`
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	TradeNo    string `json:"tradeNo,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
}

type AlipayDataDataserviceBillDownloadurlQueryResponse struct {
	HttpBody        string `json:"http_body,omitempty"`
	Code            string `json:"code,omitempty"`
	Msg             string `json:"msg,omitempty"`
	SubCode         string `json:"sub_code,omitempty"`
	SubMsg          string `json:"sub_msg,omitempty"`
	TradeNo         string `json:"tradeNo,omitempty"`
	OutTradeNo      string `json:"out_trade_no,omitempty"`
	BillDownloadUrl string `json:"bill_download_url,omitempty"`
}

type AlipayTradeQueryResponse struct {
	HttpBody            string          `json:"http_body,omitempty"`
	Code                string          `json:"code,omitempty"`
	Msg                 string          `json:"msg,omitempty"`
	SubCode             string          `json:"sub_code,omitempty"`
	SubMsg              string          `json:"sub_msg,omitempty"`
	TradeNo             string          `json:"trade_no,omitempty"`
	OutTradeNo          string          `json:"out_trade_no,omitempty"`
	BuyerLogonId        string          `json:"buyer_logon_id,omitempty"`
	TradeStatus         string          `json:"trade_status,omitempty"`
	TotalAmount         string          `json:"total_amount,omitempty"`
	TransCurrency       string          `json:"trans_currency,omitempty"`
	SettleCurrency      string          `json:"settle_currency,omitempty"`
	SettleAmount        string          `json:"settle_amount,omitempty"`
	PayCurrency         string          `json:"pay_currency,omitempty"`
	PayAmount           string          `json:"pay_amount,omitempty"`
	SettleTransRate     string          `json:"settle_trans_rate,omitempty"`
	TransPayRate        string          `json:"trans_pay_rate,omitempty"`
	BuyerPayAmount      string          `json:"buyer_pay_amount,omitempty"`
	PointAmount         string          `json:"point_amount,omitempty"`
	InvoiceAmount       string          `json:"invoice_amount,omitempty"`
	SendPayDate         string          `json:"send_pay_date,omitempty"`
	ReceiptAmount       string          `json:"receipt_amount,omitempty"`
	StoreId             string          `json:"store_id,omitempty"`
	TerminalId          string          `json:"terminal_id,omitempty"`
	FundBillList        []TradeFundBill `json:"fund_bill_list,omitempty"`
	StoreName           string          `json:"store_name,omitempty"`
	BuyerUserId         string          `json:"buyer_user_id,omitempty"`
	ChargeAmount        string          `json:"charge_amount,omitempty"`
	ChargeFlags         string          `json:"charge_flags,omitempty"`
	SettlementId        string          `json:"settlement_id,omitempty"`
	TradeSettleInfo     TradeSettleInfo `json:"trade_settle_info,omitempty"`
	AuthTradePayMode    string          `json:"auth_trade_pay_mode,omitempty"`
	BuyerUserType       string          `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string          `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string          `json:"discount_amount,omitempty"`
	BuyerUserName       string          `json:"buyer_user_name,omitempty"`
	Subject             string          `json:"subject,omitempty"`
	Body                string          `json:"body,omitempty"`
	AlipaySubMerchantId string          `json:"alipay_sub_merchant_id,omitempty"`
	ExtInfos            string          `json:"ext_infos,omitempty"`
}
type AlipayTradeCancelResponse struct {
	HttpBody           string `json:"http_body,omitempty"`
	Code               string `json:"code,omitempty"`
	Msg                string `json:"msg,omitempty"`
	SubCode            string `json:"sub_code,omitempty"`
	SubMsg             string `json:"sub_msg,omitempty"`
	TradeNo            string `json:"trade_no,omitempty"`
	OutTradeNo         string `json:"out_trade_no,omitempty"`
	RetryFlag          string `json:"retry_flag,omitempty"`
	Action             string `json:"action,omitempty"`
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId string `json:"refund_settlement_id,omitempty"`
}

type AlipayTradeFastpayRefundQueryResponse struct {
	HttpBody                     string                `json:"http_body,omitempty"`
	Code                         string                `json:"code,omitempty"`
	Msg                          string                `json:"msg,omitempty"`
	SubCode                      string                `json:"sub_code,omitempty"`
	SubMsg                       string                `json:"sub_msg,omitempty"`
	ErrorCode                    string                `json:"error_code,omitempty"`
	GmtRefundPay                 string                `json:"gmt_refund_pay,omitempty"`
	IndustrySepcDetail           string                `json:"industry_sepc_detail,omitempty"`
	OutRequestNo                 string                `json:"out_request_no,omitempty"`
	OutTradeNo                   string                `json:"out_trade_no,omitempty"`
	PresentRefundBuyerAmount     string                `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string                `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string                `json:"present_refund_mdiscount_amount,omitempty"`
	RefundAmount                 string                `json:"refund_amount,omitempty"`
	RefundChargeAmount           string                `json:"refund_charge_amount,omitempty"`
	RefundDetailItemList         []TradeFundBill       `json:"refund_detail_item_list,omitempty"`
	RefundReason                 string                `json:"refund_reason,omitempty"`
	RefundRoyaltys               []RefundRoyaltyResult `json:"refund_royaltys,omitempty"`
	RefundSettlementId           string                `json:"refund_settlement_id,omitempty"`
	RefundStatus                 string                `json:"refund_status,omitempty"`
	SendBackFee                  string                `json:"send_back_fee,omitempty"`
	TotalAmount                  string                `json:"total_amount,omitempty"`
	TradeNo                      string                `json:"trade_no,omitempty"`
}

type AlipayTradeRefundResponse struct {
	HttpBody                     string              `json:"http_body,omitempty"`
	Code                         string              `json:"code,omitempty"`
	Msg                          string              `json:"msg,omitempty"`
	SubCode                      string              `json:"sub_code,omitempty"`
	SubMsg                       string              `json:"sub_msg,omitempty"`
	TradeNo                      string              `json:"trade_no,omitempty"`
	OutTradeNo                   string              `json:"out_trade_no,omitempty"`
	BuyerLogonId                 string              `json:"buyer_logon_id,omitempty"`
	FundChange                   string              `json:"fund_change,omitempty"`
	RefundFee                    string              `json:"refund_fee,omitempty"`
	RefundCurrency               string              `json:"refund_currency,omitempty"`
	GmtRefundPay                 string              `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList         []TradeFundBill     `json:"refund_detail_item_list,omitempty"`
	StoreName                    string              `json:"store_name,omitempty"`
	BuyerUserId                  string              `json:"buyer_user_id,omitempty"`
	RefundPresetPaytoolList      []PresetPayToolInfo `json:"refund_preset_paytool_list,omitempty"`
	RefundSettlementId           string              `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string              `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string              `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string              `json:"present_refund_mdiscount_amount,omitempty"`
}

type PresetPayToolInfo struct {
	Amount         string `json:"amount,omitempty"`
	AssertTypeCode string `json:"assert_type_code,omitempty"`
}

type RefundRoyaltyResult struct {
	RefundAmount  string `json:"refund_amount,omitempty"`
	RoyaltyType   string `json:"royalty_type,omitempty"`
	ResultCode    string `json:"result_code,omitempty"`
	TransOut      string `json:"trans_out,omitempty"`
	TransOutEmail string `json:"trans_out_email,omitempty"`
	TransIn       string `json:"trans_in,omitempty"`
	TransInEmail  string `json:"trans_in_email,omitempty"`
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"`
	BankCode    string `json:"bank_code,omitempty"`
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
	FundType    string `json:"fund_type,omitempty"`
}

type TradeSettleDetail struct {
	OperationType      string `json:"operation_type,omitempty"`
	OperationSerial_no string `json:"operation_serial_no,omitempty"`
	OperationDt        string `json:"operation_dt,omitempty"`
	TransOut           string `json:"trans_out,omitempty"`
	TransIn            string `json:"trans_in,omitempty"`
	Amount             string `json:"amount,omitempty"`
}

type TradeSettleInfo struct {
	TradeSettleDetailList []TradeSettleDetail `json:"trade_settle_detail_list,omitempty"`
}
