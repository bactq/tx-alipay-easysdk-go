package facetoface

type AlipayTradePayResponse struct {
	HttpBody            string          `json:"http_body,omitempty"`
	Code                string          `json:"code,omitempty"`
	Msg                 string          `json:"msg,omitempty"`
	SubCode             string          `json:"sub_code,omitempty"`
	SubMsg              string          `json:"sub_msg,omitempty"`
	TradeNo             string          `json:"trade_no,omitempty"`
	OutTradeNo          string          `json:"out_trade_no,omitempty"`
	BuyerLogonId        string          `json:"buyer_logon_id,omitempty"`
	SettleAmount        string          `json:"settle_amount,omitempty"`
	PayCurrency         string          `json:"pay_currency,omitempty"`
	PayAmount           string          `json:"pay_amount,omitempty"`
	SettleTransRate     string          `json:"settle_trans_rate,omitempty"`
	TransPayRate        string          `json:"trans_pay_rate,omitempty"`
	TotalAmount         string          `json:"total_amount,omitempty"`
	TransCurrency       string          `json:"trans_currency,omitempty"`
	SettleCurrency      string          `json:"settle_currency,omitempty"`
	ReceiptAmount       string          `json:"receipt_amount,omitempty"`
	BuyerPayAmount      string          `json:"buyer_pay_amount,omitempty"`
	PointAmount         string          `json:"point_amount,omitempty"`
	InvoiceAmount       string          `json:"invoice_amount,omitempty"`
	GmtPayment          string          `json:"gmt_payment,omitempty"`
	FundBillList        []TradeFundBill `json:"fund_bill_list,omitempty"`
	CardBalance         string          `json:"card_balance,omitempty"`
	StoreName           string          `json:"store_name,omitempty"`
	BuyerUserId         string          `json:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string          `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   []VoucherDetail `json:"voucher_detail_list,omitempty"`
	AdvanceAmount       string          `json:"advance_amount,omitempty"`
	AuthTradePayMode    string          `json:"auth_trade_pay_mode,omitempty"`
	ChargeAmount        string          `json:"charge_amount,omitempty"`
	ChargeFlags         string          `json:"charge_flags,omitempty"`
	SettlementId        string          `json:"settlement_id,omitempty"`
	BusinessParams      string          `json:"business_params,omitempty"`
	BuyerUserType       string          `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string          `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string          `json:"discount_amount,omitempty"`
	BuyerUserName       string          `json:"buyer_user_name,omitempty"`
}

type AlipayTradePrecreateResponse struct {
	HttpBody   string `json:"http_body,omitempty"`
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	SubCode    string `json:"sub_code,omitempty"`
	SubMsg     string `json:"sub_msg,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"`
	BankCode    string `json:"bank_code,omitempty"`
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Amount                     string `json:"amount,omitempty"`
	MerchantContribute         string `json:"merchant_contribute,omitempty"`
	OtherContribute            string `json:"other_contribute,omitempty"`
	Memo                       string `json:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`
}
