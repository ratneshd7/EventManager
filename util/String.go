package util

// IsStringNil used to check string empty or not
func IsStringNil(str *string) bool {
	if str == nil {
		return true
	}
	return false
}
