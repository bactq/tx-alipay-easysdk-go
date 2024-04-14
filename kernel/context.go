package kernel

import (
	"crypto/rsa"
	"fmt"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/AlipayConstants"
	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/util"
)

type KernelContext struct {
	SdkVersion      string
	config          map[string]any
	certEnvironment *CertEnvironment
	priv            *rsa.PrivateKey
	pub             *rsa.PublicKey
}

func NewKernelContext(options Config, sdkVersion string) (*KernelContext, error) {
	this := &KernelContext{
		config: util.BuildStructMap(options),
	}
	if AlipayConstants.RSA2 != this.GetConfig(AlipayConstants.SIGN_TYPE_CONFIG_KEY) {
		return nil, fmt.Errorf("err Alipay Easy SDK只允许使用RSA2签名方式，RSA签名方式由于安全性相比RSA2弱已不再推荐。")
	}
	this.SdkVersion = sdkVersion
	if len(this.GetConfig(AlipayConstants.ALIPAY_CERT_PATH_CONFIG_KEY)) > 0 {
		var err error
		this.certEnvironment, err = NewCertEnvironment(
			this.GetConfig(AlipayConstants.MERCHANT_CERT_PATH_CONFIG_KEY),
			this.GetConfig(AlipayConstants.ALIPAY_CERT_PATH_CONFIG_KEY),
			this.GetConfig(AlipayConstants.ALIPAY_ROOT_CERT_PATH_CONFIG_KEY))
		if err != nil {
			return nil, err
		}
	}
	var err error
	this.priv, err = util.ParsePKCS8PrivateKey([]byte(options.MerchantPrivateKey))
	if err != nil {
		return nil, err
	}
	this.pub, err = util.ParsePKIXPublicKey([]byte(options.AlipayPublicKey))
	if err != nil {
		return nil, err
	}
	return this, nil
}

func (kc *KernelContext) GetConfig(key string) string {
	if val, ok := kc.config[key]; ok {
		return fmt.Sprint(val)
	} else {
		return ""
	}
}
