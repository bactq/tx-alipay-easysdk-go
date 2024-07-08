package facetoface

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

func (c *Client) Agent(appAuthToken string) *Client {
	c.InjectTextParam("app_auth_token", appAuthToken)
	return c
}
func (c *Client) Auth(authToken string) *Client {
	c.InjectTextParam("auth_token", authToken)
	return c
}

func (c *Client) AsyncNotify(url string) *Client {
	c.InjectTextParam("notify_url", url)
	return c
}

func (c *Client) Route(testUrl string) *Client {
	c.InjectTextParam("ws_service_url", testUrl)
	return c
}
func (c *Client) Optional(key, value string) *Client {
	c.InjectBizParam(key, value)
	return c
}

func (c *Client) BatchOptional(kv ...string) *Client {
	c.BizParams = append(c.BizParams, kv...)
	return c
}

func (c *Client) Pay(subject, outTradeNo, totalAmount, authCode string) (*AlipayTradePayResponse, error) {
	method := "alipay.trade.pay"
	systemParams := c.GetDefaultSystemParams(method)
	bizParams := append(c.BizParams,
		"subject", subject,
		"out_trade_no", outTradeNo,
		"total_amount", totalAmount,
		"auth_code", authCode,
		"scene", "bar_code",
	)
	textParams := c.TextParams
	bizContent := c.BizContent(bizParams)
	signParams, err := c.SignParams(systemParams, bizContent, textParams)
	if err != nil {
		return nil, err
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
	if nil != c.VerifyResp(string(body), "alipay_trade_pay_response") {
		return nil, err
	}
	ret := &AlipayTradePayResponse{}
	if nil != json.Unmarshal(body, &ret) {
		return nil, err
	}
	return ret, nil
}

func (c *Client) PreCreate(subject, outTradeNo, totalAmount string) (*AlipayTradePrecreateResponse, error) {
	method := "alipay.trade.precreate"
	systemParams := c.GetDefaultSystemParams(method)
	bizParams := append(c.BizParams,
		"subject", subject,
		"out_trade_no", outTradeNo,
		"total_amount", totalAmount,
	)
	textParams := c.TextParams
	bizContent := c.BizContent(bizParams)
	signParams, err := c.SignParams(systemParams, bizContent, textParams)
	if err != nil {
		return nil, err
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
	if nil != c.VerifyResp(string(body), "alipay_trade_precreate_response") {
		return nil, err
	}
	ret := &AlipayTradePrecreateResponse{}
	if nil != json.Unmarshal(body, &ret) {
		return nil, err
	}
	return ret, nil
}
