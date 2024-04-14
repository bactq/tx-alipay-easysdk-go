package tea

func AllowRetry(m map[string]any, retryTimes int, now int64) bool {
	if retryTimes == 0 {
		return true
	}
	if m == nil {
		return false
	}
	if shouldRetry, ok := m["retryable"].(bool); ok {
		if shouldRetry {
			if retry, ok := m["maxAttempts"].(int); ok {
				return retryTimes < retry
			}
		}
	}
	return false
}
