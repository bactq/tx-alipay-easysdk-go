package tea

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type TeaPair struct {
	Key   string
	Value any
}

func NewTeaPair(key string, value any) *TeaPair {
	return &TeaPair{Key: key, Value: value}
}

type TeaRequest struct {
	Protocol string

	Port string

	Method string

	Pathname string

	Query map[string]string

	Headers map[string]string

	Body io.Reader
}

func NewTeaRequest() *TeaRequest {
	return &TeaRequest{
		Protocol: "http",
		Method:   "GET",
		Query:    map[string]string{},
		Headers:  map[string]string{},
	}
}

func Create() *TeaRequest {
	return NewTeaRequest()
}

func Sleep(d int) {
	time.Sleep(time.Duration(d))
}
func GetBackoffTime(o any, retryTimes int) int {
	backOffTime := 0
	if m, ok := o.(map[string]any); ok {
		policy := fmt.Sprint(m["policy"])
		if len(policy) == 0 || policy == "no" {
			return backOffTime
		}
		period := fmt.Sprint(m["period"])
		if len(period) > 0 {
			if periodNum, err := strconv.ParseInt(period, 10, 32); err == nil {
				if periodNum <= 0 {
					return retryTimes
				}
			}
		}
	}
	return backOffTime
}

func ToReadable(body string) io.Reader {
	return io.NopCloser(strings.NewReader(body))
}

type TeaResponse struct{}

type StringBuilder struct {
	sb strings.Builder
}

func (s *StringBuilder) append(str string) *StringBuilder {
	s.sb.WriteString(str)
	return s
}
func (s *StringBuilder) toString() string {
	return s.sb.String()
}

func composeUrl(request *TeaRequest) string {
	queries := request.Query
	host := request.Headers["host"]
	protocol := request.Protocol
	if request.Protocol == "" {
		protocol = "http"
	}
	urlBuilder := StringBuilder{}
	urlBuilder.append(protocol)
	urlBuilder.append("://").append(host)
	if len(request.Port) > 0 {
		urlBuilder.append(":").append(request.Port)
	}
	if len(request.Pathname) > 0 {
		urlBuilder.append(request.Pathname)
	}
	if len(queries) > 0 {
		if strings.Index(urlBuilder.sb.String(), "?") > 1 {
			urlBuilder.append("&")
		} else {
			urlBuilder.append("?")
		}

		for k, v := range queries {
			urlBuilder.append(url.QueryEscape(k))
			urlBuilder.append("=")
			urlBuilder.append(url.QueryEscape(v))
			urlBuilder.append("&")
		}
		urlContent := urlBuilder.toString()
		return urlContent[:len(urlBuilder.toString())-1]
	}
	return urlBuilder.toString()
}
func DoAction(request *TeaRequest, runtime map[string]any) (string, error) {
	urlString := composeUrl(request)
	req, err := http.NewRequest(request.Method, urlString, request.Body)
	if err != nil {
		return "", err
	}
	for k, v := range request.Headers {
		req.Header[k] = []string{v}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
