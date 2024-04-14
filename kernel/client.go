package kernel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/AlipayConstants"
	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/util"
)

type KernelClient struct {
	context            *KernelContext
	optionalTextParams map[string]string
	optionalBizParams  map[string]string
}

func NewKernelClient(context *KernelContext) *KernelClient {
	return &KernelClient{context: context}
}

/**
 * 注入额外文本参数
 *
 * @param key   参数名称
 * @param value 参数的值
 * @return 本客户端本身，便于链路调用
 */
func (kc *KernelClient) InjectTextParam(k, v string) {
	kc.optionalTextParams[k] = v
}

/**
 * 注入额外业务参数
 *
 * @param key   业务参数名称
 * @param value 业务参数的值
 * @return 本客户端本身，便于链式调用
 */
func (kc *KernelClient) InjectBizParam(k, v string) {
	kc.optionalBizParams[k] = v
}

/**
 * 获取时间戳，格式yyyy-MM-dd HH:mm:ss
 *
 * @return 当前时间戳
 */
func (kc *KernelClient) GetTimestamp() string {
	return time.Now().Format(time.DateTime)
}

/**
 * 获取Config中的配置项
 *
 * @param key 配置项的名称
 * @return 配置项的值
 */
func (kc *KernelClient) GetConfig(key string) string {
	return kc.context.GetConfig(key)
}

/**
 * 获取SDK版本信息
 *
 * @return SDK版本信息
 */
func (kc *KernelClient) GetSdkVersion() string {
	return kc.context.SdkVersion
}

/**
 * 将业务参数和其他额外文本参数按www-form-urlencoded格式转换成HTTP Body中的字节数组，注意要做URL Encode
 *
 * @param bizParams 业务参数
 * @return HTTP Body中的字节数组
 */
func (kc *KernelClient) ToUrlEncodedRequestBody(bizParams map[string]any) string {
	sortedMap := kc.getSortedMap(nil, bizParams, nil)
	return kc.buildQueryString(sortedMap)
}

func (kc *KernelClient) buildQueryString(sortedMap map[string]string) string {
	needAnd := false
	sb := &strings.Builder{}
	for k, v := range sortedMap {
		if needAnd {
			sb.WriteByte('&')
		} else {
			needAnd = true
		}
		sb.WriteString(k)
		sb.WriteByte('=')
		sb.WriteString(url.QueryEscape(v))
	}
	return sb.String()
}
func (kc *KernelClient) SortMap(input map[string]string) map[string]string {
	return input
}
func (kc *KernelClient) getSortedMap(systemParams map[string]string, bizParams map[string]any,
	textParams map[string]string) map[string]string {
	kc.addOtherParams(textParams, bizParams)
	sortedMap := map[string]string{}
	for k, v := range systemParams {
		sortedMap[k] = v
	}
	if len(bizParams) > 0 {
		sortedMap[AlipayConstants.BIZ_CONTENT_FIELD] = util.ToJsonString(bizParams)
	}
	for k, v := range textParams {
		sortedMap[k] = v
	}
	return sortedMap
}

func (kc *KernelClient) addOtherParams(textParams map[string]string, bizParams map[string]any) {
	if textParams != nil {
		for k, v := range kc.optionalTextParams {
			if _, ok := textParams[k]; !ok {
				textParams[k] = v
			}
		}
		kc.setNotifyUrl(textParams)
	}
	if bizParams != nil {
		for k, v := range kc.optionalBizParams {
			if _, ok := bizParams[k]; !ok {
				bizParams[k] = v
			}
		}
	}
}

func (kc *KernelClient) setNotifyUrl(params map[string]string) {
	if len(kc.GetConfig(AlipayConstants.NOTIFY_URL_CONFIG_KEY)) > 0 {
		if _, ok := params[AlipayConstants.NOTIFY_URL_FIELD]; !ok {
			params[AlipayConstants.NOTIFY_URL_FIELD] = kc.GetConfig(AlipayConstants.NOTIFY_URL_CONFIG_KEY)
		}
	}
}

/* 将网关响应发序列化成Map，同时将API的接口名称和响应原文插入到响应Map的method和body字段中
*
* @param response HTTP响应
* @param method   调用的OpenAPI的接口名称
* @return 响应反序列化的Map
 */
func (kc *KernelClient) ReadAsJson(response, method string) map[string]any {
	m := map[string]any{}
	json.Unmarshal([]byte(response), &m)
	m[AlipayConstants.BODY_FIELD] = response
	m[AlipayConstants.METHOD_FIELD] = method
	return m
}

/**
 * 从响应Map中提取返回值对象的Map，并将响应原文插入到body字段中
 *
 * @param respMap 响应Map
 * @return 返回值对象Map
 */
func (kc *KernelClient) ToRespModel(respMap map[string]any) (map[string]any, error) {
	methodName := respMap[AlipayConstants.METHOD_FIELD].(string)
	responseNodeName := strings.ReplaceAll(methodName, ".", "_") + AlipayConstants.RESPONSE_SUFFIX
	errorNodeName := AlipayConstants.ERROR_RESPONSE
	//先找正常响应节点
	for k, v := range respMap {
		if responseNodeName == k {
			model := v.(map[string]any)
			model[AlipayConstants.BODY_FIELD] = respMap[AlipayConstants.BODY_FIELD]
			return model, nil
		}
	}
	//再找异常响应节点
	for k, v := range respMap {
		if errorNodeName == k {
			model := v.(map[string]any)
			model[AlipayConstants.BODY_FIELD] = respMap[AlipayConstants.BODY_FIELD]
			return model, nil
		}
	}
	return nil, fmt.Errorf("响应格式不符合预期，找不到" + responseNodeName + "或" + errorNodeName + "节点")
}

/**
 * 生成随机分界符，用于multipart格式的HTTP请求Body的多个字段间的分隔
 *
 * @return 随机分界符
 */
func (kc *KernelClient) GetRandomBoundary() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

/**
 * 将其他额外文本参数和文件参数按multipart/form-data格式转换成HTTP Body中的字节数组流
 *
 * @param textParams 其他额外文本参数
 * @param fileParams 业务文件参数
 * @param boundary   HTTP Body中multipart格式的分隔符
 * @return Multipart格式的字节流
 */

func (kc *KernelClient) ToMultipartRequestBody(textParams, fileParams map[string]string, boundary string) (*bytes.Buffer, error) {
	stream := &bytes.Buffer{}
	//补充其他额外参数
	kc.addOtherParams(textParams, nil)
	for k, v := range textParams {
		if len(k) > 0 && len(v) > 0 {
			stream.Write(util.GetEntryBoundary(boundary))
			stream.Write(util.GetTextEntry(k, v))
		}
	}
	//组装文件参数
	for k, v := range fileParams {
		if len(k) > 0 && len(v) > 0 {
			content, err := os.ReadFile(v)
			if err != nil {
				return nil, err
			}
			stream.Write(util.GetEntryBoundary(boundary))
			stream.Write(util.GetFileEntry(k, v))
			stream.Write(content)
		}
	}
	stream.Write(util.GetEndBoundary(boundary))
	return stream, nil
}

/**
 * 生成页面类请求所需URL或Form表单
 *
 * @param method       GET或POST，决定是生成URL还是Form表单
 * @param systemParams 系统参数集合
 * @param bizParams    业务参数集合
 * @param textParams   其他额外文本参数集合
 * @param sign         所有参数的签名值
 * @return 生成的URL字符串或表单
 */
func (kc *KernelClient) GeneratePage(method string, systemParams map[string]string,
	bizParams map[string]any, textParams map[string]string, sign string) (string, error) {
	if AlipayConstants.GET == strings.ToUpper(method) {
		sortedMap := kc.getSortedMap(systemParams, bizParams, textParams)
		sortedMap[AlipayConstants.SIGN_FIELD] = sign
		return kc.getGatewayServerUrl() + "?" + kc.buildQueryString(sortedMap), nil
	} else if AlipayConstants.POST == strings.ToUpper(method) {
		urlParams := kc.getSortedMap(systemParams, nil, textParams)
		urlParams[AlipayConstants.SIGN_FIELD] = sign
		actionUrl := kc.getGatewayServerUrl() + "?" + kc.buildQueryString(urlParams)
		kc.addOtherParams(nil, bizParams)
		formParams := map[string]string{
			AlipayConstants.BIZ_CONTENT_FIELD: util.ToJsonString(bizParams),
		}
		return util.BuildForm(actionUrl, formParams), nil
	} else {
		return "", errors.New("_generatePage中method只支持传入GET或POST")
	}
}

func (kc *KernelClient) getGatewayServerUrl() string {
	return fmt.Sprintf("%s://%s/gateway.do",
		kc.GetConfig(AlipayConstants.PROTOCOL_CONFIG_KEY),
		kc.GetConfig(AlipayConstants.HOST_CONFIG_KEY))
}

/**
 * 获取商户应用公钥证书序列号，从证书模式运行时环境对象中直接读取
 *
 * @return 商户应用公钥证书序列号
 */

func (kc *KernelClient) GetMerchantCertSN() string {
	if kc.context.certEnvironment == nil {
		return ""
	}
	return kc.context.certEnvironment.GetMerchantCertSN()
}

/**
 * 从响应Map中提取支付宝公钥证书序列号
 *
 * @param respMap 响应Map
 * @return 支付宝公钥证书序列号
 */
func (kc *KernelClient) GetAlipayCertSN(respMap map[string]any) string {
	return respMap[AlipayConstants.ALIPAY_CERT_SN_FIELD].(string)
}

/**
 * 获取支付宝根证书序列号，从证书模式运行时环境对象中直接读取
 *
 * @return 支付宝根证书序列号
 */
func (kc *KernelClient) GetAlipayRootCertSN() string {
	if kc.context.certEnvironment == nil {
		return ""
	}
	return kc.context.certEnvironment.GetRootCertSN()
}

/**
 * 是否是证书模式
 *
 * @return true：是；false：不是
 */
func (kc *KernelClient) IsCertMode() bool {
	return kc.context.certEnvironment == nil
}

/**
 * 获取支付宝公钥，从证书运行时环境对象中直接读取
 * 如果缓存的用户指定的支付宝公钥证书的序列号与网关响应中携带的支付宝公钥证书序列号不一致，需要报错给出提示或自动更新支付宝公钥证书
 *
 * @param alipayCertSN 网关响应中携带的支付宝公钥证书序列号
 * @return 支付宝公钥
 */
func (kc *KernelClient) ExtractAlipayPublicKey(alipayCertSN string) string {
	if kc.context.certEnvironment == nil {
		return ""
	}
	return kc.context.certEnvironment.GetAlipayPublicKey(alipayCertSN)
}

/**
 * 验证签名
 *
 * @param respMap         响应Map，可以从中提取出sign和body
 * @param alipayPublicKey 支付宝公钥
 * @return true：验签通过；false：验签不通过
 */
func (kc *KernelClient) Verify(respMap map[string]any) bool {
	sign := respMap[AlipayConstants.SIGN_FIELD].(string)
	content := util.GetSignSourceData(respMap[AlipayConstants.BODY_FIELD].(string),
		respMap[AlipayConstants.METHOD_FIELD].(string))
	pub := kc.context.pub
	if kc.IsCertMode() {
		pub = kc.context.certEnvironment.pub
	}
	return util.Verify(content, sign, pub)
}

/**
 * 计算签名，注意要去除key或value为null的键值对
 *
 * @param systemParams       系统参数集合
 * @param bizParams          业务参数集合
 * @param textParams         其他额外文本参数集合
 * @param merchantPrivateKey 私钥
 * @return 签名值的Base64串
 */
func (kc *KernelClient) Sign(systemParams map[string]string, bizParams map[string]any, textParams map[string]string) string {
	sortedMap := kc.getSortedMap(systemParams, bizParams, textParams)

	return util.Sign(util.GetSignCheckContent(sortedMap), kc.context.priv)
}

/**
 * AES加密
 *
 * @param plainText 明文
 * @param key       密钥
 * @return 密文
 */
func (kc *KernelClient) AesEncrypt(plainText, key string) (string, error) {
	content, err := util.Encrypt([]byte(plainText), []byte(key))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

/**
 * AES解密
 *
 * @param cipherText 密文
 * @param key        密钥
 * @return 明文
 */
func (kc *KernelClient) AesDecrypt(cipherText, key string) (string, error) {
	content, err := util.Decrypt([]byte(cipherText), []byte(key))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

/**
 * 生成订单串
 *
 * @param systemParams 系统参数集合
 * @param bizParams    业务参数集合
 * @param textParams   额外文本参数集合
 * @param sign         所有参数的签名值
 * @return 订单串
 */
func (kc *KernelClient) GenerateOrderString(method string, systemParams map[string]string,
	bizParams map[string]any, textParams map[string]string, sign string) string {
	sortedMap := kc.getSortedMap(systemParams, bizParams, textParams)
	sortedMap[AlipayConstants.SIGN_FIELD] = sign
	//将所有参数置于URL中
	return kc.buildQueryString(sortedMap)
}

/**
 * 对支付类请求的异步通知的参数集合进行验签
 *
 * @param parameters 参数集合
 * @param publicKey  支付宝公钥
 * @return true：验证成功；false：验证失败
 */
func (kc *KernelClient) VerifyParams(parameters map[string]string) bool {
	return util.VerifyParams(parameters, kc.context.pub)
}
