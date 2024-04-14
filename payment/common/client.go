package common

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

func (c *Client) Create(subject, outTradeNo, totalAmount, buyerId string) (*AlipayTradeCreateResponse, error) {
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
			tea.NewTeaPair("buyer_id", buyerId),
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
			ret := &AlipayTradeCreateResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) Query(outTradeNo string) (*AlipayTradeQueryResponse, error) {
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
			tea.NewTeaPair("method", "alipay.trade.query"),
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
			tea.NewTeaPair("out_trade_no", outTradeNo),
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
			ret := &AlipayTradeQueryResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) Refund(outTradeNo, refundAmount string) (*AlipayTradeRefundResponse, error) {
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
			tea.NewTeaPair("method", "alipay.trade.query"),
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
			tea.NewTeaPair("out_trade_no", outTradeNo),
			tea.NewTeaPair("refund_amount", refundAmount),
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
			ret := &AlipayTradeRefundResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) Close(outTradeNo string) (*AlipayTradeCloseResponse, error) {
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
			tea.NewTeaPair("method", "alipay.trade.query"),
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
			tea.NewTeaPair("out_trade_no", outTradeNo),
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
			ret := &AlipayTradeCloseResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) Cancel(outTradeNo string) (*AlipayTradeCancelResponse, error) {
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
			tea.NewTeaPair("method", "alipay.trade.query"),
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
			tea.NewTeaPair("out_trade_no", outTradeNo),
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
			ret := &AlipayTradeCancelResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) QueryRefund(outTradeNo, outRequestNo string) (*AlipayTradeFastpayRefundQueryResponse, error) {
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
			tea.NewTeaPair("method", "alipay.trade.query"),
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
			tea.NewTeaPair("out_trade_no", outTradeNo),
			tea.NewTeaPair("out_request_no", outRequestNo),
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
			ret := &AlipayTradeFastpayRefundQueryResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}

func (c *Client) DownloadBill(billType, billDate string) (*AlipayDataDataserviceBillDownloadurlQueryResponse, error) {
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
			tea.NewTeaPair("method", "alipay.trade.query"),
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
			tea.NewTeaPair("bill_type", billType),
			tea.NewTeaPair("bill_date", billDate),
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
			ret := &AlipayDataDataserviceBillDownloadurlQueryResponse{}
			json.Unmarshal([]byte(response_), &ret)
			return ret, nil
		} else {
			return nil, errors.New("验签失败，请检查支付宝公钥设置是否正确。")
		}
	}
	return nil, lastErr
}
