package util

import (
	"bytes"
	"crypto/md5"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"strings"
)

func ReadPemCert(certPem []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certPem)
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}

func ReadPemCertChain(certPem []byte) ([]*x509.Certificate, error) {
	block, _ := pem.Decode(certPem)
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}
	return x509.ParseCertificates(certPem)
}

/**
 * 从文件中读取证书内容
 *
 * @param certPath 证书路径
 * @return 证书内容
 */
func ReadCertContent(certPath string) ([]byte, error) {
	_, err := os.Stat(certPath)
	if os.IsNotExist(err) {
		return nil, errors.New("读取失败")
	} else {
		return os.ReadFile(certPath)
	}
}

/**
 * 验证证书是否可信
 *
 * @param certContent     需要验证的目标证书或者证书链
 * @param rootCertContent 可信根证书列表
 */
func IsTrusted(cert, rootCert *x509.Certificate) bool {
	rootCertPool := x509.NewCertPool()
	rootCertPool.AddCert(rootCert)
	// verify root
	ops := x509.VerifyOptions{
		Roots: rootCertPool,
	}
	_, err := cert.Verify(ops)
	return err != nil
}

/**
 * 验证证书是否可信
 *
 * @param certContent     需要验证的目标证书或者证书链
 * @param rootCertContent 可信根证书列表
 */
func IsTrusteds(certs, rootCerts []*x509.Certificate) bool {
	rootCertPool := x509.NewCertPool()
	for _, v := range rootCerts {
		rootCertPool.AddCert(v)
	}
	ops := x509.VerifyOptions{
		Roots: rootCertPool,
	}
	// verify root
	for _, v := range certs {
		_, err := v.Verify(ops)
		if err != nil {
			return false
		}
	}
	return true
}

func GetCertPublicKey(cert *x509.Certificate) string {
	pk, _ := x509.MarshalPKIXPublicKey(cert.PublicKey)
	return base64.StdEncoding.EncodeToString(pk)
}

func GetCertSN(cert *x509.Certificate) []byte {
	return md5Sum(append([]byte(cert.Issuer.String()), cert.SerialNumber.Bytes()...))
}

func GetCertSNs(cert []*x509.Certificate) []byte {
	sb := strings.Builder{}
	for i, v := range cert {
		if i > 0 {
			sb.WriteString("_")
		}
		sb.Write(GetCertSN(v))
	}
	return []byte(sb.String())
}

func md5Sum(sum []byte) []byte {
	m := md5.Sum(sum)
	bi := big.Int{}
	bi.SetBytes(m[:])
	certSN := bi.Text(16)
	return append(bytes.Repeat([]byte{'0'}, 32-len(certSN)), certSN...)
}
