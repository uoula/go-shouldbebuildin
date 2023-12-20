package shouldbebuildin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	number_string := "1234567890"
	number_reversed := ReverseString(number_string)
	assert.Equal(t, "0987654321", number_reversed)

	thai_string := "กลับด้าน"
	thai_reversed := ReverseString(thai_string)
	assert.Equal(t, "นา้ดบัลก", thai_reversed)
}
