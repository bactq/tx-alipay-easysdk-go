package facetoface

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel"
	"github.com/tianxinzizhen/tx-alipay-easysdk-go/tea"
	"github.com/tianxinzizhen/tx-alipay-easysdk-go/tea/TeaConverter"
)

type Client struct {
	_kernel *kernel.KernelClient
}

func NewClient(kernel *kernel.KernelClient) *Client {
	return &Client{_kernel: kernel}
}

func (c *Client) Pay(subject, outTradeNo, totalAmount, authCode string) (*AlipayTradePayResponse, error) {
	_kernel := c._kernel
	runtime_ := TeaConverter.BuildMap[any](
		tea.NewTeaPair("ignoreSSL", _kernel.GetConfig("ignoreSSL")),
		tea.NewTeaPair("httpProxy", _kernel.GetConfig("httpProxy")),
		tea.NewTeaPair("connectTimeout", 15000),
		tea.NewTeaPair("readTimeout", 15000),
		tea.NewTeaPair("retry", TeaConverter.BuildMap[any](
			tea.NewTeaPair("maxAttempts", 0),
		)),
	)
	_now := time.Now().UnixMilli()
	_retryTimes := 0
	var lastErr error
	for tea.AllowRetry(runtime_["retry"].(map[string]any), _retryTimes, _now) {
		if _retryTimes > 0 {
			backoffTime := tea.GetBackoffTime(runtime_["backoff"], _retryTimes)
			if backoffTime > 0 {
				tea.Sleep(backoffTime)
			}
		}
		_retryTimes = _retryTimes + 1
		request_ := tea.NewTeaRequest()
		systemParams := TeaConverter.BuildMap[string](
			tea.NewTeaPair("method", "alipay.trade.create"),
			tea.NewTeaPair("app_id", _kernel.GetConfig("appId")),
			tea.NewTeaPair("timestamp", _kernel.GetTimestamp()),
			tea.NewTeaPair("format", "json"),
			tea.NewTeaPair("version", "1.0"),
			tea.NewTeaPair("alipay_sdk", _kernel.GetSdkVersion()),
			tea.NewTeaPair("charset", "UTF-8"),
			tea.NewTeaPair("sign_type", _kernel.GetConfig("signType")),
			tea.NewTeaPair("app_cert_sn", _kernel.GetMerchantCertSN()),
			tea.NewTeaPair("alipay_root_cert_sn", _kernel.GetAlipayRootCertSN()),
		)
		bizParams := TeaConverter.BuildMap[any](
			tea.NewTeaPair("subject", subject),
			tea.NewTeaPair("out_trade_no", outTradeNo),
			tea.NewTeaPair("total_amount", totalAmount),
			tea.NewTeaPair("auth_code", authCode),
			tea.NewTeaPair("scene", "bar_code"),
		)
		textParams := map[string]string{}
		request_.Protocol = _kernel.GetConfig("protocol")
		request_.Method = "POST"
		request_.Pathname = "/gateway.do"
		request_.Headers = TeaConverter.BuildMap[string](
			tea.NewTeaPair("host", _kernel.GetConfig("gatewayHost")),
			tea.NewTeaPair("content-type", "application/x-www-form-urlencoded;charset=utf-8"),
		)
		request_.Query = _kernel.SortMap(TeaConverter.MergeString(TeaConverter.BuildMap[string](
			tea.NewTeaPair("sign", _kernel.Sign(systemParams, bizParams, textParams))),
			systemParams,
			textParams,
		))
		request_.Body = tea.ToReadable(_kernel.ToUrlEncodedRequestBody(bizParams))
		response_, err := tea.DoAction(request_, runtime_)
		if err != nil {
			lastErr = err
			continue
		}
		respMap := _kernel.ReadAsJson(response_, "alipay.trade.create")
		if _kernel.Verify(respMap) {
			ret := &AlipayTradePayResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) PreCreate(subject, outTradeNo, totalAmount string) (*AlipayTradePrecreateResponse, error) {
	_kernel := c._kernel
	runtime_ := TeaConverter.BuildMap[any](
		tea.NewTeaPair("ignoreSSL", _kernel.GetConfig("ignoreSSL")),
		tea.NewTeaPair("httpProxy", _kernel.GetConfig("httpProxy")),
		tea.NewTeaPair("connectTimeout", 15000),
		tea.NewTeaPair("readTimeout", 15000),
		tea.NewTeaPair("retry", TeaConverter.BuildMap[any](
			tea.NewTeaPair("maxAttempts", 0),
		)),
	)
	_now := time.Now().UnixMilli()
	_retryTimes := 0
	var lastErr error
	for tea.AllowRetry(runtime_["retry"].(map[string]any), _retryTimes, _now) {
		if _retryTimes > 0 {
			backoffTime := tea.GetBackoffTime(runtime_["backoff"], _retryTimes)
			if backoffTime > 0 {
				tea.Sleep(backoffTime)
			}
		}
		_retryTimes = _retryTimes + 1
		request_ := tea.NewTeaRequest()
		systemParams := TeaConverter.BuildMap[string](
			tea.NewTeaPair("method", "alipay.trade.create"),
			tea.NewTeaPair("app_id", _kernel.GetConfig("appId")),
			tea.NewTeaPair("timestamp", _kernel.GetTimestamp()),
			tea.NewTeaPair("format", "json"),
			tea.NewTeaPair("version", "1.0"),
			tea.NewTeaPair("alipay_sdk", _kernel.GetSdkVersion()),
			tea.NewTeaPair("charset", "UTF-8"),
			tea.NewTeaPair("sign_type", _kernel.GetConfig("signType")),
			tea.NewTeaPair("app_cert_sn", _kernel.GetMerchantCertSN()),
			tea.NewTeaPair("alipay_root_cert_sn", _kernel.GetAlipayRootCertSN()),
		)
		bizParams := TeaConverter.BuildMap[any](
			tea.NewTeaPair("subject", subject),
			tea.NewTeaPair("out_trade_no", outTradeNo),
			tea.NewTeaPair("total_amount", totalAmount),
		)
		textParams := map[string]string{}
		request_.Protocol = _kernel.GetConfig("protocol")
		request_.Method = "POST"
		request_.Pathname = "/gateway.do"
		request_.Headers = TeaConverter.BuildMap[string](
			tea.NewTeaPair("host", _kernel.GetConfig("gatewayHost")),
			tea.NewTeaPair("content-type", "application/x-www-form-urlencoded;charset=utf-8"),
		)
		request_.Query = _kernel.SortMap(TeaConverter.MergeString(TeaConverter.BuildMap[string](
			tea.NewTeaPair("sign", _kernel.Sign(systemParams, bizParams, textParams))),
			systemParams,
			textParams,
		))
		request_.Body = tea.ToReadable(_kernel.ToUrlEncodedRequestBody(bizParams))
		response_, err := tea.DoAction(request_, runtime_)
		if err != nil {
			lastErr = err
			continue
		}
		respMap := _kernel.ReadAsJson(response_, "alipay.trade.create")
		if _kernel.Verify(respMap) {
			ret := &AlipayTradePrecreateResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}
