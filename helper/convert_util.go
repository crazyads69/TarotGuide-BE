package helper

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

// Define function convert genai.Part to string
func ConvertPartToString(part genai.Part) string {
	// Convert part
	str := fmt.Sprintf("%v", part)
	return str
}
