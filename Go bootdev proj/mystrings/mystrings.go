package mystrings

func Reverse(s string) string { // capital R so we can use it outside this directory
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
