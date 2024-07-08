package common

import (
	"encoding/json"
	"errors"
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s?%s", c.BaseUrl, c.ToUrlEncodedRequestBody(signParams, systemParams, textParams)),
		strings.NewReader(c.ToUrlEncodedRequestBody(bizContent)))
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
	ret := &AlipayTradeCreateResponse{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	//验签
	if c.VerifyResp(string(body), "alipay_trade_create_response") {
		return ret, nil
	} else {
		// 验签失败
		return nil, errors.New("alipay_trade_create_response:验签失败")
	}
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s?%s", c.BaseUrl, c.ToUrlEncodedRequestBody(signParams, systemParams, textParams)),
		strings.NewReader(c.ToUrlEncodedRequestBody(bizContent)))
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
	ret := &AlipayTradeQueryResponse{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	//验签
	if c.VerifyResp(string(body), "alipay_trade_query_response") {
		return ret, nil
	} else {
		// 验签失败
		return nil, errors.New("alipay_trade_query_response:验签失败")
	}
}
