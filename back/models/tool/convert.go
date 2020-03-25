package tool

func Convert(src map[string][]string) map[string]interface{} {
	dest := map[string]interface{}{}
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

func Convert2(src map[string]string) map[string]interface{} {
	dest := map[string]interface{}{}
	for k, v := range src {
		dest[k] = v
	}
	return dest
}
