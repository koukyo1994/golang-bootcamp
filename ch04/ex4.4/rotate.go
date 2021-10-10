package rotate

func Rotate(slice []int, k int) []int {
	n := len(slice)
	k %= n
	if k == 0 {
		return slice
	}
	return append(slice[n-k:], slice[:n-k]...)
}
