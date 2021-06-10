package bmak_tools

import "strconv"

func convertBytetoInt(byteSlice []byte) []int {
	intSlice := make([]int, len(byteSlice))
	for i, b := range byteSlice {
		intSlice[i] = int(b)
	}
	return intSlice
}

func convertIntToString(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))
	for i, b := range intSlice {
		stringSlice[i] = strconv.Itoa(b)
	}
	return stringSlice
}