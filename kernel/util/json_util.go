package util

import "encoding/json"

// toJsonString
func ToJsonString(input map[string]any) string {
	jsonBytes, _ := json.Marshal(input)
	return string(jsonBytes)
}
