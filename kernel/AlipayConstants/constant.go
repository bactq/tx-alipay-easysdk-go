package AlipayConstants

const (
	/**
	 * Config配置参数Key值
	 */
	PROTOCOL_CONFIG_KEY              = "protocol"
	HOST_CONFIG_KEY                  = "gatewayHost"
	ALIPAY_CERT_PATH_CONFIG_KEY      = "alipayCertPath"
	MERCHANT_CERT_PATH_CONFIG_KEY    = "merchantCertPath"
	ALIPAY_ROOT_CERT_PATH_CONFIG_KEY = "alipayRootCertPath"
	SIGN_TYPE_CONFIG_KEY             = "signType"
	NOTIFY_URL_CONFIG_KEY            = "notifyUrl"
	SIGN_PROVIDER_CONFIG_KEY         = "signProvider"

	/**
	 * 与网关HTTP交互中涉及到的字段值
	 */
	BIZ_CONTENT_FIELD    = "biz_content"
	ALIPAY_CERT_SN_FIELD = "alipay_cert_sn"
	SIGN_FIELD           = "sign"
	SIGN_TYPE_FIELD      = "sign_type"
	BODY_FIELD           = "http_body"
	NOTIFY_URL_FIELD     = "notify_url"
	METHOD_FIELD         = "method"
	RESPONSE_SUFFIX      = "_response"
	ERROR_RESPONSE       = "error_response"

	/**
	 * 默认字符集编码，EasySDK统一固定使用UTF-8编码，无需用户感知编码，用户面对的总是String而不是bytes
	 */
	DEFAULT_CHARSET = "UTF_8"

	/**
	 * 默认的签名算法，EasySDK统一固定使用RSA2签名算法（即SHA_256_WITH_RSA），但此参数依然需要用户指定以便用户感知，因为在开放平台接口签名配置界面中需要选择同样的算法
	 */
	RSA2 = "RSA2"

	/**
	 * RSA2对应的真实签名算法名称
	 */
	SHA_256_WITH_RSA = "SHA256WithRSA"

	/**
	 * RSA2对应的真实非对称加密算法名称
	 */
	RSA = "RSA"

	/**
	 * 申请生成的重定向网页的请求类型，GET表示生成URL
	 */
	GET = "GET"

	/**
	 * 申请生成的重定向网页的请求类型，POST表示生成form表单
	 */
	POST = "POST"

	/**
	 * 使用Aliyun KMS签名服务时签名提供方的名称
	 */
	AliyunKMS = "AliyunKMS"
)
