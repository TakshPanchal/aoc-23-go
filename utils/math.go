package utils

func SumAll(arr []int) (sum int) {
	for i := range arr {
		sum += arr[i]
	}
	return
}
