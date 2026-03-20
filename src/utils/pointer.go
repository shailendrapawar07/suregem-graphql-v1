package utils

func GetStringPtr(m map[string]interface{}, key string) *string {
	if val, ok := m[key].(string); ok {
		return &val
	}
	return nil
}
