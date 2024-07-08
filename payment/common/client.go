package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/factory"
)

type Client struct {
	*factory.Client
}

func NewClient(client *factory.Client) *Client {
	return &Client{client}
}

func (c *Client) Create(subject, outTradeNo, totalAmount, buyerId string) (*AlipayTradeCreateResponse, error) {
	method := "alipay.trade.create"
	systemParams := c.GetDefaultSystemParams(method)
	bizParams := []string{
		"subject", subject,
		"out_trade_no", outTradeNo,
		"total_amount", totalAmount,
		"buyer_id", buyerId,
	}
	textParams := []string{
		"notify_url", c.NotifyUrl,
	}
	bizContent := []string{
		"biz_content", c.ToJson(bizParams),
	}
	sign, err := c.SignData(systemParams, bizContent, textParams)
	if err != nil {
		return nil, err
	}
	signParams := []string{
		"sign", sign,
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s?%s", c.BaseUrl, c.ToUrlEncoded(signParams, systemParams, textParams)),
		strings.NewReader(c.ToUrlEncoded(bizContent)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if nil != c.VerifyResp(string(body), "alipay_trade_create_response") {
		return nil, err
	}
	ret := &AlipayTradeCreateResponse{}
	if nil != json.Unmarshal(body, &ret) {
		return nil, err
	}
	return ret, nil
}

func (c *Client) Query(outTradeNo string) (*AlipayTradeQueryResponse, error) {
	method := "alipay.trade.query"
	systemParams := c.GetDefaultSystemParams(method)
	bizParams := []string{
		"out_trade_no", outTradeNo,
	}
	textParams := []string{
		"notify_url", c.NotifyUrl,
	}
	bizContent := []string{
		"biz_content", c.ToJson(bizParams),
	}
	sign, err := c.SignData(systemParams, bizContent, textParams)
	if err != nil {
		return nil, err
	}
	signParams := []string{
		"sign", sign,
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s?%s", c.BaseUrl, c.ToUrlEncoded(signParams, systemParams, textParams)),
		strings.NewReader(c.ToUrlEncoded(bizContent)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if nil != c.VerifyResp(string(body), "alipay_trade_query_response") {
		return nil, err
	}
	ret := &AlipayTradeQueryResponse{}
	if nil != json.Unmarshal(body, &ret) {
		return nil, err
	}
	return ret, nil
}
