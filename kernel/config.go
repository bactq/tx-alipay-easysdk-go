package kernel

type Config struct {
	/**
	 * 通信协议，通常填写https
	 */
	Protocol string `json:"protocol"`
	/**
	 * 网关域名
	 * 线上为：openapi.alipay.com
	 * 沙箱为：openapi.alipaydev.com
	 */
	GatewayHost string `json:"gatewayHost"`
	/**
	 * AppId
	 */
	AppId string `json:"appId"`
	/**
	 * 签名类型，Alipay Easy SDK只推荐使用RSA2，估此处固定填写RSA2
	 */
	SignType string `json:"signType"`
	/**
	 * 支付宝公钥
	 */
	AlipayPublicKey string `json:"alipayPublicKey"`
	/**
	 * 应用私钥
	 */
	MerchantPrivateKey string `json:"merchantPrivateKey"`
	/**
	 * 应用公钥证书文件路径
	 */
	MerchantCertPath string `json:"merchantCertPath"`
	/**
	 * 支付宝公钥证书文件路径 `json:"protocol"`
	 */
	AlipayCertPath string `json:"alipayCertPath"`
	/**
	 * 支付宝根证书文件路径
	 */
	AlipayRootCertPath string `json:"alipayRootCertPath"`
	/**
	 * 异步通知回调地址（可选）
	 */
	NotifyUrl string `json:"notifyUrl"`
	/**
	 * AES密钥（可选）
	 */
	EncryptKey string `json:"encryptKey"`
	/**
	 * 签名提供方的名称(可选)，例：Aliyun KMS签名，signProvider = "AliyunKMS"
	 */
	SignProvider string `json:"signProvider"`
	/**
	 * 代理地址（可选）
	 * 例如：http://127.0.0.1:8080
	 */
	HttpProxy string `json:"httpProxy"`
	/**
	 * 忽略证书校验（可选）
	 */
	IgnoreSSL bool `json:"ignoreSSL"`
}
