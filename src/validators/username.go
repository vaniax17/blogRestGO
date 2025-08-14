package validators

func IsMinOrMaxLengthOfUsername(username string) bool {
	if len(username) < 3 && len(username) > 30 {
		return false
	}
	return true
}
