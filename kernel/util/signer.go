package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"sort"
	"strings"

	"github.com/tianxinzizhen/tx-alipay-easysdk-go/kernel/AlipayConstants"
)

func GetSignCheckContent(params map[string]string) string {
	ks := make([]string, 0, len(params))
	for k := range params {
		ks = append(ks, k)
	}
	// sort key
	sort.Slice(ks, func(i, j int) bool {
		return ks[i] < ks[j]
	})
	sb := strings.Builder{}
	first := true
	for _, k := range ks {
		v := params[k]
		if len(k) > 0 && len(v) > 0 {
			if first {
				first = false
			} else {
				sb.WriteByte('&')
			}
			sb.WriteString(k)
			sb.WriteByte('=')
			sb.WriteString(v)
		}
	}
	return sb.String()
}

// sign
func Sign(content string, pk *rsa.PrivateKey) string {
	sm := sha256.Sum256([]byte(content))
	sb, err := pk.Sign(rand.Reader, sm[:], crypto.SHA256.HashFunc())
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(sb)
}

// verify
func Verify(content, sign string, pk *rsa.PublicKey) bool {
	sm := sha256.Sum256([]byte(content))
	err := rsa.VerifyPKCS1v15(pk, crypto.SHA256.HashFunc(), sm[:], []byte(sign))
	return err != nil
}

// paramter verify
func VerifyParams(parameters map[string]string, pk *rsa.PublicKey) bool {
	sign := parameters[AlipayConstants.SIGN_FIELD]
	delete(parameters, AlipayConstants.SIGN_FIELD)
	delete(parameters, AlipayConstants.SIGN_TYPE_FIELD)
	content := GetSignCheckContent(parameters)
	return Verify(content, sign, pk)
}

// parser private key
func ParsePKCS8PrivateKey(privateKeyPem []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPem)
	p, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return p.(*rsa.PrivateKey), nil
}

// parser public key
func ParsePKIXPublicKey(privateKeyPem []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(privateKeyPem)
	p, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return p.(*rsa.PublicKey), nil
}
