package kernel

import (
	"crypto/rsa"
	"errors"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/util"
)

type CertEnvironment struct {
	/**
	 * 支付宝根证书内容
	 */
	rootCertContent []byte

	/**
	 * 支付宝根证书序列号
	 */
	rootCertSN []byte

	/**
	 * 商户应用公钥证书序列号
	 */
	merchantCertSN []byte

	/**
	 * 缓存的不同支付宝公钥证书序列号对应的支付宝公钥
	 */
	cachedAlipayPublicKey map[string]string
	pub                   *rsa.PublicKey
}

func NewCertEnvironment(merchantCertPath, alipayCertPath, alipayRootCertPath string) (*CertEnvironment, error) {
	if len(merchantCertPath) == 0 || len(alipayCertPath) == 0 || len(alipayRootCertPath) == 0 {
		return nil, errors.New("证书参数merchantCertPath、alipayCertPath或alipayRootCertPath设置不完整。")
	}
	this := &CertEnvironment{}
	var err error
	this.rootCertContent, err = util.ReadCertContent(alipayRootCertPath)
	if err != nil {
		return nil, err
	}
	rootCert, err := util.ReadPemCert(this.rootCertContent)
	if err != nil {
		return nil, err
	}
	this.rootCertSN = util.GetCertSN(rootCert)
	// merchantCertPath
	certContent, err := util.ReadCertContent(merchantCertPath)
	if err != nil {
		return nil, err
	}
	cert, err := util.ReadPemCert(certContent)
	if err != nil {
		return nil, err
	}
	this.merchantCertSN = util.GetCertSN(cert)

	// alipayCertPath
	alipayCertContent, err := util.ReadCertContent(alipayCertPath)
	if err != nil {
		return nil, err
	}
	alipayCert, err := util.ReadPemCert(alipayCertContent)
	if err != nil {
		return nil, err
	}
	this.cachedAlipayPublicKey[string(util.GetCertSN(alipayCert))] = util.GetCertPublicKey(alipayCert)
	this.pub = alipayCert.PublicKey.(*rsa.PublicKey)
	return this, nil
}

func (c *CertEnvironment) GetRootCertSN() string {
	return string(c.rootCertSN)
}

func (c *CertEnvironment) GetMerchantCertSN() string {
	return string(c.merchantCertSN)
}
func (c *CertEnvironment) GetAlipayPublicKey(sn string) string {
	return c.cachedAlipayPublicKey[sn]
}
