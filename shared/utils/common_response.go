package utils

// CommonResponse responses common json data.
type CommonResponse struct {
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}
