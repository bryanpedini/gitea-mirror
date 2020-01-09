package utils

// StringSliceToMap Converts a slice of strings into a map formed by
// keys equals to the slice values, and empty strings as map values.
func StringSliceToMap(slice []string) map[string]string {
	ret := make(map[string]string)

	for i := 0; i < len(slice); i++ {
		ret[slice[i]] = ""
	}

	return ret
}
