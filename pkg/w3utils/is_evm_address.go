package w3utils

func IsEVMAddress(address string) bool {
	return IsAddress(address)
}

func IsAddress(address string) bool {
	return len(address) == 42 && IsHex(address)
}

func IsHex(str string) bool {
	for _, c := range str {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
