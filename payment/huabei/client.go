package huabei

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

func (c *Client) Create(subject, outTradeNo, totalAmount, buyerId string, extendParams HuabeiConfig) (*AlipayTradeCreateResponse, error) {
	method := "alipay.trade.create"
	systemParams := c.GetDefaultSystemParams(method)
	// 获取扩展参数的json
	extendParamsJson, err := json.Marshal(extendParams)
	if err != nil {
		return nil, err
	}
	bizParams := append(c.BizParams,
		"subject", subject,
		"out_trade_no", outTradeNo,
		"total_amount", totalAmount,
		"buyer_id", buyerId,
		"extend_params", string(extendParamsJson),
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
	err = c.VerifyResp(string(body), "alipay_trade_create_response")
	if err != nil {
		return nil, err
	}
	ret := &AlipayTradeCreateResponse{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
