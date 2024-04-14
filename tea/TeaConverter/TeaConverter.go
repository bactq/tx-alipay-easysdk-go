package TeaConverter

import "github.com/tianxinzizhen/tx-alipay-easysdk-go/tea"

func BuildMap[T any](pairs ...*tea.TeaPair) map[string]T {
	m := make(map[string]T, len(pairs))
	for _, pair := range pairs {
		m[pair.Key] = pair.Value.(T)
	}
	return m
}

func MergeString(params ...map[string]string) map[string]string {
	ret := map[string]string{}
	for _, m := range params {
		for k, v := range m {
			ret[k] = v
		}
	}
	return ret
}
