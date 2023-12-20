package shouldbebuildin

func ReverseString(s string) string {

	var s_out_rune []rune
	s_rune := []rune(s)

	for i := len(s_rune) - 1; i >= 0; i-- {
		s_out_rune = append(s_out_rune, s_rune[i])
	}

	return string(s_out_rune)
}
