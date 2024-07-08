package factory

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	AppId      string
	BaseUrl    string
	priv       *rsa.PrivateKey // 应用私钥
	pub        *rsa.PublicKey  // 支付宝公钥
	TextParams []string
	BizParams  []string
}

func NewAliPay() (*Client, error) {
	return &Client{}, nil
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

// sign
func Sign(content []byte, pk *rsa.PrivateKey) (string, error) {
	sm := sha256.Sum256(content)
	sb, err := pk.Sign(rand.Reader, sm[:], crypto.SHA256.HashFunc())
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sb), nil
}

// verify
func Verify(content, sign string, pk *rsa.PublicKey) bool {
	sm := sha256.Sum256([]byte(content))
	signContent, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}

	err = rsa.VerifyPKCS1v15(pk, crypto.SHA256.HashFunc(), sm[:], signContent)
	return err == nil
}

func (c *Client) SignParams(datas ...[]string) ([]string, error) {
	var kvs [][2]string
	for _, data := range datas {
		for i := 0; i+1 < len(data); i += 2 {
			kvs = append(kvs, [2]string{data[i], data[i+1]})
		}
	}
	// 签名之前需要针对名字排序
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i][0] < kvs[j][0]
	})
	bb := &bytes.Buffer{}
	for i, v := range kvs {
		if i > 0 {
			bb.WriteByte('&')
		}
		bb.WriteString(v[0])
		bb.WriteByte('=')
		bb.WriteString(v[1])
	}
	sign, err := Sign(bb.Bytes(), c.priv)
	if err != nil {
		return nil, err
	}
	return []string{
		"sign", sign,
	}, nil
}

func getResponseContent(body, methodResp string) string {
	const (
		LEFT_BRACE    = '{'
		RIGHT_BRACE   = '}'
		DOUBLE_QUOTES = '"'
	)
	index := strings.Index(body, methodResp)
	if index == -1 {
		return ""
	}
	index += len(methodResp) + 2
	mod := 0
	start := 0
	end := 0
	braceDeep := 0
	for i := index; i < len(body); i++ {
		if mod == 0 {
			if body[i] == LEFT_BRACE {
				mod = LEFT_BRACE
				start = i
			} else if body[i] == DOUBLE_QUOTES {
				mod = DOUBLE_QUOTES
				start = i
			}
		} else {
			if mod == LEFT_BRACE {
				if body[i] == LEFT_BRACE {
					braceDeep++
				} else if body[i] == RIGHT_BRACE {
					if braceDeep == 0 {
						end = i + 1
						break
					}
					braceDeep--
				}
			} else if mod == DOUBLE_QUOTES {
				if body[i] == DOUBLE_QUOTES {
					end = i + 1
					break
				}
			}
		}
	}
	if mod == 0 || start >= end {
		return ""
	}
	return body[start:end]
}

func (c *Client) VerifyResp(body, method string) error {
	content := getResponseContent(body, method)
	sign := getResponseContent(body, "sign")
	if !Verify(content, sign, c.pub) {
		return fmt.Errorf("%s 验签失败", method)
	}
	return nil
}

func Encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func (c *Client) VerifyNotify(notify string) bool {
	urlMap, err := url.ParseQuery(notify)
	if err != nil {
		return false
	}
	var sign string
	if signUrl, ok := urlMap["sign"]; ok {
		sign = signUrl[0]
	}
	if len(sign) == 0 {
		return false
	}
	delete(urlMap, "sign")
	delete(urlMap, "sign_type")
	return Verify(Encode(urlMap), sign, c.pub)
}

func (c *Client) ToUrlEncoded(datas ...[]string) string {
	urlValues := url.Values{}
	for _, data := range datas {
		for i := 0; i+1 < len(data); i += 2 {
			urlValues.Add(data[i], data[i+1])
		}
	}
	return urlValues.Encode()
}

func (c *Client) GetDefaultSystemParams(method string) []string {
	return []string{
		"method", method,
		"app_id", c.AppId,
		"timestamp", time.Now().Format(time.DateTime),
		"format", "json",
		"version", "1.0",
		"charset", "UTF-8",
		"sign_type", "RSA2",
	}
}

func (c *Client) ToJson(kv []string) string {
	bb := &bytes.Buffer{}
	bb.WriteByte('{')
	for i := 1; i < len(kv); i += 2 {
		if i > 0 {
			bb.WriteByte(',')
		}
		// name
		bb.WriteByte('"')
		bb.WriteString(kv[i-1])
		bb.WriteByte('"')
		// :
		bb.WriteByte(':')
		// value
		bb.WriteByte('"')
		bb.WriteString(kv[i])
		bb.WriteByte('"')
	}
	bb.WriteByte('}')
	return bb.String()
}

func (c *Client) BizContent(bizParams []string) []string {
	return []string{
		"biz_content", c.ToJson(bizParams),
	}
}
func (c *Client) InjectTextParam(key, value string) *Client {
	c.TextParams = append(c.TextParams, key, value)
	return c
}

func (c *Client) InjectBizParam(key, value string) *Client {
	c.BizParams = append(c.BizParams, key, value)
	return c
}
