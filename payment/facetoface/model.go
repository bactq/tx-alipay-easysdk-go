package facetoface

type VoucherDetail struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Amount string `json:"amount"`

	MerchantContribute string `json:"merchant_contribute"`

	OtherContribute string `json:"other_contribute"`

	Memo string `json:"memo"`

	TemplateId string `json:"template_id"`

	PurchaseBuyerContribute string `json:"purchase_buyer_contribute"`

	PurchaseMerchantContribute string `json:"purchase_merchant_contribute"`

	PurchaseAntContribute string `json:"purchase_ant_contribute"`
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel"`

	BankCode string `json:"bank_code"`

	Amount string `json:"amount"`

	RealAmount string `json:"real_amount"`
}
type AlipayTradePrecreateResponse struct {
	HttpBody string `json:"http_body"`

	Code string `json:"code"`

	Msg string `json:"msg"`

	SubCode string `json:"sub_code"`

	SubMsg string `json:"sub_msg"`

	OutTradeNo string `json:"out_trade_no"`

	QrCode string `json:"qr_code"`
}
type AlipayTradePayResponse struct {
	// 响应原始字符串

	HttpBody string `json:"http_body"`

	Code string `json:"code"`

	Msg string `json:"msg"`

	SubCode string `json:"sub_code"`

	SubMsg string `json:"sub_msg"`

	TradeNo string `json:"trade_no"`

	OutTradeNo string `json:"out_trade_no"`

	BuyerLogonId string `json:"buyer_logon_id"`

	SettleAmount string `json:"settle_amount"`

	PayCurrency string `json:"pay_currency"`

	PayAmount string `json:"pay_amount"`

	SettleTransRate string `json:"settle_trans_rate"`

	TransPayRate string `json:"trans_pay_rate"`

	TotalAmount string `json:"total_amount"`

	TransCurrency string `json:"trans_currency"`

	SettleCurrency string `json:"settle_currency"`

	ReceiptAmount string `json:"receipt_amount"`

	BuyerPayAmount string `json:"buyer_pay_amount"`

	PointAmount string `json:"point_amount"`

	InvoiceAmount string `json:"invoice_amount"`

	GmtPayment string `json:"gmt_payment"`

	FundBillList []*TradeFundBill `json:"fund_bill_list"`

	CardBalance string `json:"card_balance"`

	StoreName string `json:"store_name"`

	BuyerUserId string `json:"buyer_user_id"`

	DiscountGoodsDetail string `json:"discount_goods_detail"`

	VoucherDetailList []*VoucherDetail `json:"voucher_detail_list"`

	AdvanceAmount string `json:"advance_amount"`

	AuthTradePayMode string `json:"auth_trade_pay_mode"`

	ChargeAmount string `json:"charge_amount"`

	ChargeFlags string `json:"charge_flags"`

	SettlementId string `json:"settlement_id"`

	BusinessParams string `json:"business_params"`

	BuyerUserType string `json:"buyer_user_type"`

	MdiscountAmount string `json:"mdiscount_amount"`

	DiscountAmount string `json:"discount_amount"`

	BuyerUserName string `json:"buyer_user_name"`
}
