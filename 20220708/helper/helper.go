package 20220708

type Helper struct{}

func (h *Helper) min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
